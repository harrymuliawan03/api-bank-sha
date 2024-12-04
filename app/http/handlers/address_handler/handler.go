package addresshandler

import (
	"context"
	"net/http"
	"time"

	"github.com/harrymuliawan03/api-bank-sha.git/app/configs"
	"github.com/harrymuliawan03/api-bank-sha.git/app/dto"
	"github.com/harrymuliawan03/api-bank-sha.git/app/http/requests"
	addressservice "github.com/harrymuliawan03/api-bank-sha.git/app/services/address_service"
	"github.com/harrymuliawan03/api-bank-sha.git/app/utils"

	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	addressService addressservice.AddressService
	cnf         *configs.Config
}

func NewAddressHandler(addressService addressservice.AddressService, config *configs.Config) *AddressHandler {
	return &AddressHandler{
		addressService: addressService,
		cnf:         config,
	}
}

func (u *AddressHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.AddressCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.addressService.Create(c, req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "address created successfully", nil)
}

func (u *AddressHandler) FindByID(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.AddressGetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	res, err := u.addressService.FindByID(c, req.UserId)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get address successfully", res)
}

func (u *AddressHandler) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	
	if id == 0 {
		return dto.ResponseApiError(ctx, "Address id is required", http.StatusBadRequest, nil)
	}

	res, err := u.addressService.FindByID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get address successfully", res)
}

func(u *AddressHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.AddressUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.addressService.Update(c, req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "address updated successfully", nil)
}

func (u *AddressHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		return dto.ResponseApiError(ctx, "Address id is required", http.StatusBadRequest, nil)
	}

	err := u.addressService.Delete(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "address deleted successfully", nil)
}