package cityrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/facades"
	"github.com/harrymuliawan03/api-bank-sha.git/app/models"

	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	orm *gorm.DB
}

// Delete implements AddressRepository.
func (c *AddressRepositoryImpl) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindByIDUser implements AddressRepository.
func (c *AddressRepositoryImpl) FindByIDUser(ctx context.Context, id uint) (result models.Address, err error) {
	err = c.orm.Where("user_id = ?", id).First(&result).Error
	return
}

// Save implements AddressRepository.
func (c *AddressRepositoryImpl) Save(ctx context.Context, data *models.Address) error {
	panic("unimplemented")
}

// Update implements AddressRepository.
func (c *AddressRepositoryImpl) Update(ctx context.Context, data *models.Address) error {
	panic("unimplemented")
}

func NewAddressRepository() AddressRepository {
	return &AddressRepositoryImpl{orm: facades.Orm()}
}
