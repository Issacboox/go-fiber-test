package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	// /api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", c.HelloTest)
	// BodyParser
	v1.Post("/", c.TestBodyParser)
	// Params
	v1.Get("/user/:name", c.TestParams)
	// Query
	v1.Post("/inet", c.QueryTest)
	// Validate Test
	v1.Post("/valid", c.ValidTest)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTest)
}
