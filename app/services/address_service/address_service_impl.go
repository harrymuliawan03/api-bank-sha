package addressservice

import (
	"context"
	"errors"

	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
	addressrepo "github.com/harrymuliawan03/api-bank-sha.git/app/repositories/address_repo"
	"github.com/harrymuliawan03/api-bank-sha.git/app/schemas"
	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	addressRepository addressrepo.AddressRepository
}

// FindAllByUserID implements AddressService.
func (c *AddressServiceImpl) FindAllByUserID(ctx context.Context, id uint) ([]dto.AddressData, error) {
	var result []dto.AddressData

	addresses, err := c.addressRepository.FindAllByIDUser(ctx, id)
	for _, address := range addresses {
		result = append(result, dto.AddressData{
			ID:          address.ID,
			City:        address.City,
			StreetAddress: address.StreetAddress,
			State:       address.State,
			PostalCode:  address.PostalCode,
			IsPrimary:   address.IsPrimary,
		})
	}

	return result, err
}

func NewAddressService(cr addressrepo.AddressRepository) AddressService {
	return &AddressServiceImpl{addressRepository: cr}
}

// Create implements AddressService.
func (c *AddressServiceImpl) Create(ctx context.Context, req requests.AddressCreateRequest) error {
	address := models.Address{UserID: req.UserID, City: req.City, StreetAddress: req.StreetAddress, State: req.State, PostalCode: req.PostalCode, IsPrimary: false}
	return c.addressRepository.Save(ctx, &address)
}

// Delete implements AddressService.
func (c *AddressServiceImpl) Delete(ctx context.Context, id uint) error {
	return c.addressRepository.Delete(ctx, id)
}

// FindAll implements AddressService.
func (c *AddressServiceImpl) FindAll(ctx context.Context) ([]dto.AddressData, error) {
	panic("unimplemented")
}

// FindByID implements AddressService.
func (c *AddressServiceImpl) FindByID(ctx context.Context, id uint) (*dto.AddressData, error) {
	address, err := c.addressRepository.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "address not found",
		}
		return nil, err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
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
func (c *AddressServiceImpl) Update(ctx context.Context, id uint, req requests.AddressUpdateRequest) error {
	_, err := c.addressRepository.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "address not found",
		}
		return err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return err
	}

	address := models.Address{ID: id, City: req.City, StreetAddress: req.StreetAddress, State: req.State, PostalCode: req.PostalCode, IsPrimary: false}
	return c.addressRepository.Update(ctx, &address)
}
