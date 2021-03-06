package routes

import (
	"fmt"

	"github.com/PiterWeb/CalendarWebServer/middlewares"
	"github.com/PiterWeb/CalendarWebServer/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routes() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	port := configRoutes(app)

	indexRoutes := app.Group("/")

	indexRoutes.Get("/", Index)
	indexRoutes.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	// User API protected Routes

	userRoutesProtected := app.Group("/api/user")

	userRoutesProtected.Use(middlewares.Session)

	userRoutesProtected.Post("/createEvent", api.CreateEventRoute)
	userRoutesProtected.Get("/getEvents", api.GetUserEventsRoute)
	userRoutesProtected.Get("/getImportantEvents", api.GetUserImportantEventsRoute)
	// userRoutesProtected.Post("/getEventsByDate", api.GetUserEventsByDateRoute)
	// userRoutesProtected.Get("/getEventsById", api.GetUserEventByIdRoute)
	userRoutesProtected.Delete("/deleteEventById", api.DeleteUserEventByIdRoute)
	userRoutesProtected.Post("/updateEventById", api.UpdateUserEventByIdRoute)

	// AUTH API protected routes

	apiProtectedRoutes := app.Group("/api/protected")

	apiProtectedRoutes.Use(middlewares.Session)

	apiProtectedRoutes.Delete("/delete", api.DeleteRoute)

	// AUTH API public Routes

	apiRoutes := app.Group("/api")

	apiRoutes.Post("/login", api.LoginRoute)

	apiRoutes.Post("/register", api.RegisterRoute)

	apiRoutes.Delete("/logout", api.LogoutRoute)

	fmt.Println(port)

	app.Listen(port)

}
