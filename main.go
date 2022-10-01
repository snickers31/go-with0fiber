// ./main.go

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snickers31/go-with-fiber/pkg/configs"
	"github.com/snickers31/go-with-fiber/pkg/middleware"
	"github.com/snickers31/go-with-fiber/pkg/routes"
	"github.com/snickers31/go-with-fiber/pkg/utils"

	"github.com/gofiber/swagger"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/snickers31/go-with-fiber/docs"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)
	app.Get("/swagger/*", swagger.HandlerDefault)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)

}
