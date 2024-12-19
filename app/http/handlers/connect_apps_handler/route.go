package connectappshandler

import (
	"github.com/harrymuliawan03/api-bank-sha.git/app/configs"
	connectappsrepo "github.com/harrymuliawan03/api-bank-sha.git/app/repositories/connect_apps_repo"
	connectappsservice "github.com/harrymuliawan03/api-bank-sha.git/app/services/connect_apps_service"

	"github.com/gofiber/fiber/v2"
)

func AddressRoute(route fiber.Router, cnf *configs.Config) {
	repo := connectappsrepo.NewConnectAppsRepository()

	service := connectappsservice.NewConnectAppsService(repo)

	handler := NewConnectAppsHandler(service, cnf)
	address := route.Group("/connect_apps")

	address.Get("/:id", handler.Show)
	address.Get("/user/:id", handler.FindAllByIDUser)
	address.Put("/:id", handler.Update)
	address.Post("/", handler.Create)
	address.Delete("/:id", handler.Delete)
}
