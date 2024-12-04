package addressservice

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
)

type AddressService interface {
	Create(ctx context.Context, req requests.AddressCreateRequest) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]dto.AddressData, error)
	FindByID(ctx context.Context, id uint) (*dto.AddressData, error)
	Show(ctx context.Context, id uint) (*dto.AddressData, error)
	Update(ctx context.Context, req requests.AddressUpdateRequest) error
}
