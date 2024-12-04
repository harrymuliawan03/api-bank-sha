package addresshandler

import (
	"github.com/harrymuliawan03/api-bank-sha.git/app/configs"
	addressrepo "github.com/harrymuliawan03/api-bank-sha.git/app/repositories/address_repo"
	addressservice "github.com/harrymuliawan03/api-bank-sha.git/app/services/address_service"

	"github.com/gofiber/fiber/v2"
)

func AddressRoute(route fiber.Router, cnf *configs.Config) {
	repo := addressrepo.NewAddressRepository()

	service := addressservice.NewAddressService(repo)

	handler := NewAddressHandler(service, cnf)
	address := route.Group("/address")

	address.Get("/:id", handler.Show)
	address.Put("/:id", handler.Update)
	address.Post("/", handler.Create)
	address.Delete("/:id", handler.Delete)
}
