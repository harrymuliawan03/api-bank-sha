package authverifrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
)

type AuthVerifRepository interface {
	Save(ctx context.Context, data *models.AuthVerif) error
	Update(ctx context.Context, data *models.AuthVerif) error
	Delete(ctx context.Context, id uint) error
	FindByCode(ctx context.Context, code string) (models.AuthVerif, error)
}
	