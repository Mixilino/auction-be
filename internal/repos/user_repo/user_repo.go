package user_repo

import (
	"auction-be/models"
	"context"
)

type User interface {
	Save(ctx context.Context, user models.User) error
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}
