package auth_service

import (
	"auction-be/conversions"
	"auction-be/dtos/user_dtos"
	"auction-be/internal/errors/repoerrors"
	"auction-be/internal/errors/resterrors"
	"auction-be/internal/repos/user_repo"
	"auction-be/internal/utils"
	"auction-be/models"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

type AuthImpl struct {
	db       *sql.DB
	userRepo user_repo.User
}

func NewAuthService(db *sql.DB, user user_repo.User) Auth {
	return &AuthImpl{db: db, userRepo: user}
}

func (a AuthImpl) Login(userLoginDTO *user_dtos.UserLoginRequest) (string, *resterrors.RestError) {
	user, err := a.userRepo.FindByUsername(context.Background(), userLoginDTO.UserName)
	if err != nil {
		if errors.Is(err, repoerrors.ErrRepoRecordNotFound) {
			return "", resterrors.CreateNew(http.StatusNotFound, "User_Not_Found", "User not found")
		}

		return "", resterrors.CreateNew(http.StatusInternalServerError, "DB__Transaction_Error", "Error when finding user")
	}

	if !utils.CheckPasswordHash(userLoginDTO.Password, user.Password) {
		return "", resterrors.CreateNew(http.StatusUnauthorized, "Invalid_Password", "Password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", resterrors.CreateNew(http.StatusInternalServerError, "Token_Generation_Error", "Error when generating token")
	}

	return token, nil
}

func (a AuthImpl) SignUp(userDTO *user_dtos.SignUpRequest) (*user_dtos.Response, *resterrors.RestError) {
	if err := utils.ValidatePassword(userDTO.Password); err != nil {
		return nil, resterrors.CreateNew(http.StatusBadRequest, "Invalid_Password", "Password must contain at least 8 characters, 1 uppercase, 1 lowercase and 1 number")
	}

	hashedPass, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		return nil, resterrors.CreateNew(http.StatusInternalServerError, "Hashing_Password_Error", "Error when hashing password")
	}

	var user *models.User
	user = conversions.UserSignUpRequestToUserModel(*userDTO, hashedPass)

	tx, err := a.db.Begin()
	if err != nil {
		return nil, resterrors.CreateNew(http.StatusInternalServerError, "DB__Transaction_Error", "Error when starting transaction")
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "tx", tx)

	if err := a.userRepo.Save(ctx, *user); err != nil {
		tx.Rollback()
		if errors.Is(err, repoerrors.ErrRepoUniqueEmailViolation) {
			return nil, resterrors.CreateNew(http.StatusConflict, "DB__Unique_Email_Violation", "Email already exists")
		}

		if errors.Is(err, repoerrors.ErrRepoUniqueUserNameViolation) {
			return nil, resterrors.CreateNew(http.StatusConflict, "DB__Unique_Username_Violation", "Username already exists")
		}

		return nil, resterrors.CreateNew(http.StatusInternalServerError, "DB__Transaction_Error", "Error when saving user")
	}

	if err := tx.Commit(); err != nil {
		return nil, resterrors.CreateNew(http.StatusInternalServerError, "DB__Transaction_Error", "Error when committing transaction")
	}
	return nil, nil
}
