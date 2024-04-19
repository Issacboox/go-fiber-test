package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":    "doe",
			"admin":   "123456",
			"gofiber": "21022566",
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

	//Test 5.0
	authV1 := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	})

	//Test 5.1
	v1.Get("/fact/:number", c.FindFacts, authV1)

	//Test 5.2
	v3 := api.Group("/v3")

	v3.Post("/bam", c.ConvertAscii, authV1)

	// Test 6
	v1.Post("/register", c.RegisterForm)
}
