// ./app/controllers/token_controller.go

package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/snickers31/go-with-fiber/app/models"
	"github.com/snickers31/go-with-fiber/pkg/utils"
	"github.com/snickers31/go-with-fiber/platform/database"
)

func GenerateNewAccessToken(c *fiber.Ctx) error {
	user := &models.User{}

	fmt.Println(string(c.Request().Body()))

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	uc, err := db.UserCredential(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = utils.ComparePassword(user.Password, uc.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   "Invalid password.",
			})
	}

	token, err := utils.GenerateNewAccessToken(uc.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})

}
