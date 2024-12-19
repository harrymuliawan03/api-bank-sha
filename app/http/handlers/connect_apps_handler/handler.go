package connectappshandler

import (
	"context"
	"net/http"
	"time"

	"github.com/harrymuliawan03/api-bank-sha.git/app/configs"
	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
	connectAppsService "github.com/harrymuliawan03/api-bank-sha.git/app/services/connect_apps_service"
	"github.com/harrymuliawan03/api-bank-sha.git/app/utils"

	"github.com/gofiber/fiber/v2"
)

type ConnectAppsHandler struct {
	connectAppsService connectAppsService.ConnectAppsService
	cnf         *configs.Config
}

func NewConnectAppsHandler(connectAppsService connectAppsService.ConnectAppsService, config *configs.Config) *ConnectAppsHandler {
	return &ConnectAppsHandler{
		connectAppsService: connectAppsService,
		cnf:         config,
	}
}

func (u *ConnectAppsHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.ConnectAppsCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.connectAppsService.Create(c, req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "connect apps created successfully", nil)
}

func (u *ConnectAppsHandler) FindAllByIDUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	res, err := u.connectAppsService.FindAllByUserID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get connect apps successfully", res)
}

func (u *ConnectAppsHandler) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	
	if id == 0 {
		return dto.ResponseApiError(ctx, "ConnectApps id is required", http.StatusBadRequest, nil)
	}

	res, err := u.connectAppsService.FindByID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get connect apps successfully", res)
}

func(u *ConnectAppsHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "ConnectApps id is required", http.StatusBadRequest, nil)
	}

	var req requests.ConnectAppsUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.connectAppsService.Update(c, uint(id), req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "connect apps updated successfully", nil)
}

func (u *ConnectAppsHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		return dto.ResponseApiError(ctx, "ConnectApps id is required", http.StatusBadRequest, nil)
	}

	err := u.connectAppsService.Delete(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "connect apps deleted successfully", nil)
}