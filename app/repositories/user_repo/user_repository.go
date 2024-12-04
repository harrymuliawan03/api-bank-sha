package userrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (models.User, error)
	FindByID(ctx context.Context, id uint) (models.User, error)
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id uint) error
}
