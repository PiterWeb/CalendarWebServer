package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"os"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func configRoutes(app *fiber.App) string {

	godotenv.Load()

	app.Use(etag.New())

	app.Use(func(c * fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})

	app.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
		AllowHeaders: "Content-Type,Access-Control-Allow-Origin",
	}))

	cookieKey := os.Getenv("SECRET_COOKIE")

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: cookieKey,
	}))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
		return port
	}

	port = ":" + port

	return port
}
