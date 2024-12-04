package addressservice

import (
	"context"
	"errors"

	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
	addressrepo "github.com/harrymuliawan03/api-bank-sha.git/app/repositories/address_repo"
	"github.com/harrymuliawan03/api-bank-sha.git/app/schemas"
	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	addressRepository addressrepo.AddressRepository
}

func NewAddressService(cr addressrepo.AddressRepository) AddressService {
	return &AddressServiceImpl{addressRepository: cr}
}

// Create implements AddressService.
func (c *AddressServiceImpl) Create(ctx context.Context, req requests.AddressCreateRequest) error {
	panic("unimplemented")
}

// Delete implements AddressService.
func (c *AddressServiceImpl) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindAll implements AddressService.
func (c *AddressServiceImpl) FindAll(ctx context.Context) ([]dto.AddressData, error) {
	panic("unimplemented")
}

// FindByID implements AddressService.
func (c *AddressServiceImpl) FindByID(ctx context.Context, id uint) (*dto.AddressData, error) {
	address, err := c.addressRepository.FindByIDUser(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status: schemas.ApiErrorNotFound,
			Message: "address not found",
		}
		return nil, err
	}else if err != nil {
		err = &schemas.ResponseApiError{
			Status: schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return nil, err
	}

	return &dto.AddressData{ID: address.ID, StreetAddress: address.StreetAddress, City: address.City, State: address.State, PostalCode: address.PostalCode, IsPrimary: address.IsPrimary}, nil
}

// Show implements AddressService.
func (c *AddressServiceImpl) Show(ctx context.Context, id uint) (*dto.AddressData, error) {
	panic("unimplemented")
}

// Update implements AddressService.
func (c *AddressServiceImpl) Update(ctx context.Context, req requests.AddressUpdateRequest) error {
	panic("unimplemented")
}
