package user_repo

import (
	"auction-be/internal/errors/repoerrors"
	"auction-be/models"
	"context"
	"database/sql"
	"errors"
)

type userSQLImpl struct {
	db *sql.DB
}

func NewUserSqlImpl(db *sql.DB) User {
	return &userSQLImpl{db: db}
}

func (u userSQLImpl) Save(ctx context.Context, user models.User) error {
	tx := ctx.Value("tx").(*sql.Tx)

	query := "INSERT INTO users (username, password, email, first_name, last_name) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, created_at, updated_at"
	row := tx.QueryRow(query, user.UserName, user.Password,
		user.Email, user.FirstName, user.LastName)

	err := row.Scan(&user.ID, &user.CreateAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"unique_email\"" {
			return repoerrors.ErrRepoUniqueEmailViolation
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"unique_username\"" {
			return repoerrors.ErrRepoUniqueUserNameViolation
		}
		return err
	}

	return nil
}

func (u userSQLImpl) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	query := "SELECT * FROM users WHERE username = $1"
	row := u.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.UserName, &user.Password,
		&user.Email, &user.FirstName, &user.LastName, &user.CreateAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoerrors.ErrRepoRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}
