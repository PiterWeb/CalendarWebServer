package utils

import (
	"github.com/PiterWeb/CalendarWebServer/models"

	"github.com/gofiber/fiber/v2"
)

var user *models.User

func GetUserFromSession(c *fiber.Ctx) (models.User, error) {

	session, err := JsonToStruct(c.Cookies("session"))

	if err != nil {
		return *user, err
	}

	user = &models.User{
		Username: session["username"],
		Password: session["password"],
	}

	return *user, nil

}
