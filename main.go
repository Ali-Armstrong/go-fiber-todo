package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	//requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	config "github.com/mohammedalizi/go-fiber-todo/config"
	routes "github.com/mohammedalizi/go-fiber-todo/routes"
)


func setupRoutes(app *fiber.App) {

    // give response when at /
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "success":  true,
            "message": "You are at the endpoint 😉",
        })
    })

    // api group
    api := app.Group("/api")

    // give response when at /api
    api.Get("", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "success": true,
            "message": "You are at the api endpoint 😉",
        })
    })

    // connect todo routes
    routes.TodoRoute(api.Group("/todos"))
}

func main() {
    app := fiber.New()

	//app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
	  Format: "${pid} ${locals:requestid} ${status} ${method} ${latency} ${path}\n",
	}))

    // dotenv
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // config db
    config.ConnectDB()


	setupRoutes(app)

    // Listen on server 8000 and catch error if any
    err = app.Listen(":8000")

    // handle error
    if err != nil {
        panic(err)
    }
}