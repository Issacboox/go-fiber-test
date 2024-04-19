package controllers

import (
	"log"

	m "go-fiber-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func TestBodyParser(c *fiber.Ctx) error {
	p := new(m.Person)
	if err := c.BodyParser(p); err != nil {
		l ///
		return c.Status(fiber.StatusBadRequest).SendString("Bad Req")
	}
	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func TestParams(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}
