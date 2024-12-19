package connectappsrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
)

type ConnectAppsRepository interface {
	Save(ctx context.Context, data *models.ConnectApps) error
	FindByID(ctx context.Context, id uint) (models.ConnectApps, error)
	FindAllByIDUser(ctx context.Context, id uint) ([]models.ConnectApps, error)
	Update(ctx context.Context, data *models.ConnectApps) error
	Delete(ctx context.Context, id uint) error
}
	