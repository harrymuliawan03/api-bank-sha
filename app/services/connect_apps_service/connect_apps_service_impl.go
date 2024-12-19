package connectappsservice

import (
	"context"
	"errors"

	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
	"github.com/harrymuliawan03/api-bank-sha.git/app/models"
	connectappsrepo "github.com/harrymuliawan03/api-bank-sha.git/app/repositories/connect_apps_repo"
	"github.com/harrymuliawan03/api-bank-sha.git/app/schemas"
	"gorm.io/gorm"
)

type ConnectAppsServiceImpl struct {
	connectAppsRepository connectappsrepo.ConnectAppsRepository
}

// FindAllByUserID implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) FindAllByUserID(ctx context.Context, id uint) ([]dto.ConnectAppsData, error) {
	var result []dto.ConnectAppsData

	connectAppses, err := c.connectAppsRepository.FindAllByIDUser(ctx, id)
	for _, connectApps := range connectAppses {
		result = append(result, dto.ConnectAppsData{
			ID:          connectApps.ID,
			UserID:      connectApps.UserID,
			AppName:     connectApps.AppName,
		})
	}

	return result, err
}

func NewConnectAppsService(cr connectappsrepo.ConnectAppsRepository) ConnectAppsService {
	return &ConnectAppsServiceImpl{connectAppsRepository: cr}
}

// Create implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) Create(ctx context.Context, req requests.ConnectAppsCreateRequest) error {
	connectApps := models.ConnectApps{UserID: req.UserID, AppName: req.AppName}
	return c.connectAppsRepository.Save(ctx, &connectApps)
}

// Delete implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) Delete(ctx context.Context, id uint) error {
	return c.connectAppsRepository.Delete(ctx, id)
}

// FindAll implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) FindAll(ctx context.Context) ([]dto.ConnectAppsData, error) {
	panic("unimplemented")
}

// FindByID implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) FindByID(ctx context.Context, id uint) (*dto.ConnectAppsData, error) {
	connectApps, err := c.connectAppsRepository.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "connect apps not found",
		}
		return nil, err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return nil, err
	}

	return &dto.ConnectAppsData{ID: connectApps.ID, UserID: connectApps.UserID, AppName: connectApps.AppName}, nil
}

// Show implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) Show(ctx context.Context, id uint) (*dto.ConnectAppsData, error) {
	panic("unimplemented")
}

// Update implements ConnectAppsService.
func (c *ConnectAppsServiceImpl) Update(ctx context.Context, id uint, req requests.ConnectAppsUpdateRequest) error {
	_, err := c.connectAppsRepository.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "connect apps not found",
		}
		return err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return err
	}

	connectApps := models.ConnectApps{ID: id, UserID: req.UserID, AppName: req.AppName}
	return c.connectAppsRepository.Update(ctx, &connectApps)
}
