package connectappsrepo

import (
	"context"

	"github.com/harrymuliawan03/api-bank-sha.git/app/facades"
	"github.com/harrymuliawan03/api-bank-sha.git/app/models"

	"gorm.io/gorm"
)

type ConnectAppsRepositoryImpl struct {
	orm *gorm.DB
}

// FindAllByIDUser implements ConnectAppsRepository.
func (c *ConnectAppsRepositoryImpl) FindAllByIDUser(ctx context.Context, id uint) (result []models.ConnectApps, err error) {
	err = c.orm.Where("user_id = ?", id).Find(&result).Error
	return
}

// Delete implements ConnectAppsRepository.
func (c *ConnectAppsRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return c.orm.WithContext(ctx).Where("id = ?", id).Delete(&models.ConnectApps{}).Error
}

// FindByIDUser implements ConnectAppsRepository.
func (c *ConnectAppsRepositoryImpl) FindByID(ctx context.Context, id uint) (result models.ConnectApps, err error) {
	err = c.orm.Where("id = ?", id).First(&result).Error
	return
}

// Save implements ConnectAppsRepository.
func (c *ConnectAppsRepositoryImpl) Save(ctx context.Context, data *models.ConnectApps) error {
	return c.orm.WithContext(ctx).Create(data).Error
}

// Update implements ConnectAppsRepository.
func (c *ConnectAppsRepositoryImpl) Update(ctx context.Context, data *models.ConnectApps) error {
	return c.orm.WithContext(ctx).Where("id = ?", data.ID).Updates(data).Error
}

func NewConnectAppsRepository() ConnectAppsRepository {
	return &ConnectAppsRepositoryImpl{orm: facades.Orm()}
}
