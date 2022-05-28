package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teris-io/shortid"
	"piterdev.com/app/db"
	"piterdev.com/app/models"
	"piterdev.com/app/utils"
)

func CreateEventRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	var event models.Event

	err = c.BodyParser(&event)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Event",
		})
	}

	event.Id = shortid.MustGenerate()

	err = db.CreateUserCalendarEvent(&user, event)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Event created",
	})

}

func GetUserEventsRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	events, err := db.GetUserCalendarEvents(&user)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"message": events,
	})

}

func GetUserImportantEventsRoute(c *fiber.Ctx) error {
	
	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	events, err := db.GetUserCalendarImportantEvents(&user)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"message": events,
	})
}

func GetUserEventsByDateRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	var date models.Date

	err = c.BodyParser(&date)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Date",
		})
	}

	events, err := db.GetUserCalendarEventsByDate(&user, date)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": events,
	})

}

func GetUserEventByIdRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	id := c.Params("id")

	if id == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
	}

	event, err := db.GetUserCalendarEventById(&user, id)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": event,
	})

}

func DeleteUserEventByIdRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	id := c.Query("id")

	if id == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
	}

	err = db.DeleteUserCalendarEventById(&user, id)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Event deleted or not found",
	})

}

func UpdateUserEventByIdRoute(c *fiber.Ctx) error {

	user, err := utils.GetUserFromSession(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	id := c.Query("id")

	if id == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
	}

	var event models.Event

	err = c.BodyParser(&event)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Event",
		})
	}

	err = db.UpdateUserCalendarEventById(&user, id, event)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Event updated",
	})

}