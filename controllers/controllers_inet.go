package controllers

import (
	m "go-fiber-test/models"
	"log"
	"regexp"

	// "regexp"
	"strconv"
	// "unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func TestBodyParser(c *fiber.Ctx) error {
	p := new(m.Person)
	if err := c.BodyParser(p); err != nil {
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

//Test 5.1

func FindFacts(c *fiber.Ctx) error {
	numberParam := c.Params("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid number provided",
		})
	}
	fact := factorial(number)
	return c.JSON(fiber.Map{
		"number":    number,
		"factorial": fact,
	})
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Test 5.2
func ConvertAscii(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")
	ascii := ConvertToAscii(taxID)
	return c.JSON(ascii)
}

func ConvertToAscii(taxID string) string {
	ascii := ""
	for _, char := range taxID {
		ascii += strconv.Itoa(int(char)) + " "
	}
	return ascii
}

// Test 6
func RegisterForm(c *fiber.Ctx) error {
	// Connect to database
	account := new(m.Register)
	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()

	// Register custom validation functions
	validate.RegisterValidation("viladate-username", isValidUsername)
	validate.RegisterValidation("viladate-website", isValidWebsiteLink)

	// Validate the account struct
	if errors := validate.Struct(account); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.Error(),
		})
	}

	return c.JSON(account)
}

func isValidUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(username)
}

func isValidWebsiteLink(fl validator.FieldLevel) bool {
	website := fl.Field().String()
	return regexp.MustCompile(`^[a-z0-9-]{2,28}\.[a-z]{2,20}$`).MatchString(website)
}
