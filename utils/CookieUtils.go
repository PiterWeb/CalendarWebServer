package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateCookie(name string, value string, durationHours int) *fiber.Cookie {

	cookie := new(fiber.Cookie)

	cookie.Name = name

	cookie.Value = value

	cookie.Expires = time.Now().Add(time.Duration(durationHours) * time.Hour)

	return cookie

}
