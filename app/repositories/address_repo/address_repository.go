package cityrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
)

type AddressRepository interface {
	Save(ctx context.Context, data *models.Address) error
	FindByIDUser(ctx context.Context, id uint) (models.Address, error)
	Update(ctx context.Context, data *models.Address) error
	Delete(ctx context.Context, id uint) error
}
	