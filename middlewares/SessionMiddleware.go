package middlewares

import (
	"piterdev.com/app/models"
	"github.com/gofiber/fiber/v2"
	"piterdev.com/app/utils"
	"piterdev.com/app/routes/api"
)

func Session (c *fiber.Ctx) error {

	var user *models.User

	if c.Cookies("session") != "" {

		session, err := utils.JsonToStruct(c.Cookies("session"))

		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid session",
			})
		}

		user = &models.User{
			Username: session["username"],
			Password: session["password"],
		}

		err = api.LoginUser(user)

		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
				"status":  "error",
			})
		}

		cookie := utils.CreateCookie("session", c.Cookies("session"), 8760)

		c.Cookie(cookie)

		return c.Next()

	}

	return c.JSON(fiber.Map{
		"message": "Invalid session",
		"status":  "error",
	})

}