// ./pkg/routes/public_routes.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snickers31/go-with-fiber/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Get("/books/:id", controllers.GetBook)
	route.Post("/login", controllers.GenerateNewAccessToken)

}
