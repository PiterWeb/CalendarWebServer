package api

import (
	"time"

	"github.com/PiterWeb/CalendarWebServer/models"
	"github.com/PiterWeb/CalendarWebServer/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginRoute(c *fiber.Ctx) error {

	var user *models.User

	err := c.BodyParser(&user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid request",
			"status":  "error",
		})
	}

	err = LoginUser(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	cookieValue, err := utils.StructToJson(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	cookie := utils.CreateCookie("session", cookieValue, 8760)

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User logged in",
		"status":  "ok",
	})

}

func RegisterRoute(c *fiber.Ctx) error {

	var user *models.User

	err := c.BodyParser(&user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid request",
			"status":  "error",
		})
	}

	err = registerUser(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "session"

	cookie.Value, err = utils.StructToJson(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	cookie.Expires = time.Now().Add(24 * time.Hour * 365)

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User Registered",
		"status":  "ok",
	})

}

func LogoutRoute(c *fiber.Ctx) error {

	if c.Cookies("session") != "" {

		c.ClearCookie("session")

		return c.JSON(fiber.Map{
			"message": "User logged out",
			"status":  "ok",
		})

	}

	return c.JSON(fiber.Map{
		"message": "No session found on client",
		"status":  "error",
	})

}

func DeleteRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	err = deleteUser(&user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	c.ClearCookie("session")

	return c.JSON(fiber.Map{
		"message": "User deleted",
		"status":  "ok",
	})

}
