package connectappsservice

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
)

type ConnectAppsService interface {
	Create(ctx context.Context, req requests.ConnectAppsCreateRequest) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]dto.ConnectAppsData, error)
	FindByID(ctx context.Context, id uint) (*dto.ConnectAppsData, error)
	FindAllByUserID(ctx context.Context, id uint) ([]dto.ConnectAppsData, error)
	Show(ctx context.Context, id uint) (*dto.ConnectAppsData, error)
	Update(ctx context.Context, id uint,req requests.ConnectAppsUpdateRequest) error
}
