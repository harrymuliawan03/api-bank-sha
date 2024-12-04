package routes

import (
	"github.com/harrymuliawan03/api-bank-sha.git/app/configs"
	addresshandler "github.com/harrymuliawan03/api-bank-sha.git/app/http/handlers/address_handler"
	authhandler "github.com/harrymuliawan03/api-bank-sha.git/app/http/handlers/auth_handler"
	"github.com/harrymuliawan03/api-bank-sha.git/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, jwt fiber.Handler, cnf *configs.Config) {
	api := app.Group("/api")
	authhandler.UserRoute(api, jwt, cnf)
	addresshandler.AddressRoute(api, cnf)

	// Listing routes
	pkg.ListRoutes(app)
}
