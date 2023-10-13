package auth_service

import (
	user_dtos2 "auction-be/dtos/user_dtos"
	"auction-be/internal/errors/resterrors"
)

type Auth interface {
	Login(userLoginDTO *user_dtos2.UserLoginRequest) (string, *resterrors.RestError)
	SignUp(userSignUpDTO *user_dtos2.SignUpRequest) (*user_dtos2.Response, *resterrors.RestError)
}
