package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"john":    "doe",
	// 		"admin":   "123456",
	// 		"gofiber": "21022566",
	// 	},
	// }))

	//Test 5.0
	authV1 := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
			"testgo":  "23012023",
		},
	})

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

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	//Dog get Del Data
	dog.Get("/deldata", c.GetDeletedDogs)
	//Get dog data more than 50 less than 100
	dog.Get("/moreless", c.GetDogs7)
	//Get dog Json color
	dog.Get("/colorjson", c.GetDogsColorJson)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTest)

	//Test 5.1
	v1.Get("/fact/:number", authV1, c.FindFacts)
	//Test 5.2
	v3 := api.Group("/v3")
	v3.Post("/bam", authV1, c.ConvertAscii)
	// Test 6
	v1.Post("/register", c.RegisterForm)

	//Company 7.0.1
	company := v1.Group("/company")
	company.Get("", c.GetCompany)
	company.Get("/filter", c.GetCompanyFilter)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

	//CRUD profile
	profile := v1.Group("/profile")
	profile.Get("", c.GetProfiles)
	profile.Get("/filter", authV1, c.GetProfile)
	profile.Post("/", authV1, c.AddProfile)
	profile.Put("/:id", authV1, c.UpdateProfile)
	profile.Delete("/:id", authV1, c.RemoveProfile)

	//Search by employee_id, name ,lastname
	profile.Get("/find", authV1, c.SearchProfile)
	profile.Get("/json", authV1, c.GetProfileJson)

}
