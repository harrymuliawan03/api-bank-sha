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

// FindAllByIDUser implements AddressRepository.
func (c *AddressRepositoryImpl) FindAllByIDUser(ctx context.Context, id uint) (result []models.Address, err error) {
	err = c.orm.Where("user_id = ?", id).Find(&result).Error
	return
}

// Delete implements AddressRepository.
func (c *AddressRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return c.orm.WithContext(ctx).Where("id = ?", id).Delete(&models.Address{}).Error
}

// FindByIDUser implements AddressRepository.
func (c *AddressRepositoryImpl) FindByID(ctx context.Context, id uint) (result models.Address, err error) {
	err = c.orm.Where("id = ?", id).First(&result).Error
	return
}

// Save implements AddressRepository.
func (c *AddressRepositoryImpl) Save(ctx context.Context, data *models.Address) error {
	return c.orm.WithContext(ctx).Create(data).Error
}

// Update implements AddressRepository.
func (c *AddressRepositoryImpl) Update(ctx context.Context, data *models.Address) error {
	return c.orm.WithContext(ctx).Where("id = ?", data.ID).Updates(data).Error
}

func NewAddressRepository() AddressRepository {
	return &AddressRepositoryImpl{orm: facades.Orm()}
}
