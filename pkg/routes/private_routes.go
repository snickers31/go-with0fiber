// ./pkg/routes/private_routes.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snickers31/go-with-fiber/app/controllers"
	"github.com/snickers31/go-with-fiber/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/books", middleware.JWTProtected(), controllers.CreateBook)
	route.Put("/books/:id", middleware.JWTProtected(), controllers.UpdateBook)
	route.Delete("/books/:id", middleware.JWTProtected(), controllers.DeleteBook)
	route.Get("/users", controllers.GetUsers)
	route.Post("/users", controllers.CreateUser)
}
