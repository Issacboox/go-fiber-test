package controllers

import (
	db "go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	// "regexp"
	"strconv"
	// "unicode"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	return regexp.MustCompile(`^[a-z0-9-]{2,28}\.[a-z]{2,5}$`).MatchString(website)
}

// CRUD ⌨️
// สร้างfunc ที่หาค่ามากกว่า 100 ไปใช้กับการ Query ฐานข้อมูล
func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

// 💭 Read only
func GetDogs(c *fiber.Ctx) error {
	var dogs []m.Dogs
	db := db.DBConn

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

// 📌 Read and search
func GetDog(c *fiber.Ctx) error {
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs
	db := db.DBConn

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

// 🥛 Create
func AddDog(c *fiber.Ctx) error {
	var dog m.Dogs
	db := db.DBConn

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

// 🔱 Update
func UpdateDog(c *fiber.Ctx) error {
	var dog m.Dogs
	id := c.Params("id")
	db := db.DBConn

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

// 🗑️ Delete
func RemoveDog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// 💡 Get Json
func GetDogsJson(c *fiber.Ctx) error {
	db := db.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

// Company
// Get Company
func GetCompany(c *fiber.Ctx) error {
	var comp []m.Company
	db := db.DBConn

	db.Find(&comp)
	return c.Status(200).JSON(comp)
}

// Get Filter Company
func GetCompanyFilter(c *fiber.Ctx) error {
	search := strings.TrimSpace(c.Query("search"))
	var comp []m.Company
	db := db.DBConn

	result := db.Find(&comp, "id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&comp)
}

// Post Company
func AddCompany(c *fiber.Ctx) error {
	var comp m.Company
	db := db.DBConn

	if err := c.BodyParser(&comp); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	validate := validator.New()

	// Register custom validation functions
	validate.RegisterValidation("website", isValidWebsiteLink)

	if errors := validate.Struct(comp); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.Error(),
		})
	}

	// return c.JSON(account)

	db.Create(&comp)
	return c.Status(201).JSON(comp)
}

// Put Company
func UpdateCompany(c *fiber.Ctx) error {
	var comp m.Company
	id := c.Params("id")
	db := db.DBConn

	if err := c.BodyParser(&comp); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&comp)
	return c.Status(200).JSON(comp)
}

// Delete Company
func RemoveCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var comp m.Company

	result := db.Delete(&comp, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// Get Deleted Dogs
func GetDeletedDogs(c *fiber.Ctx) error {
	var dogs []m.Dogs
	db := db.DBConn

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func DogIDGreaterThan50LessThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ? AND dog_id < ?", 50, 100)
}

// 💭 Read only
func GetDogs7(c *fiber.Ctx) error {
	var dogs []m.Dogs
	db := db.DBConn

	db.Scopes(DogIDGreaterThan50LessThan100).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

// Json Dog
// Get Json
func GetDogsColorJson(c *fiber.Ctx) error {
	db := db.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) // 10 ตัว

	var dataResults []m.DogsRes
	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNoColor := 0

	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNoColor++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs),
		Sum_Red:     sumRed,
		Sum_Green:   sumGreen,
		Sum_Pink:    sumPink,
		Sum_NoColor: sumNoColor,
	}

	return c.Status(200).JSON(r)

}

// CRUD Profile
func GetProfiles(c *fiber.Ctx) error {
	var prof []m.Profile
	db := db.DBConn

	db.Find(&prof)
	return c.Status(200).JSON(prof)
}
func GetProfile(c *fiber.Ctx) error {
	search := strings.TrimSpace(c.Query("search"))
	var prof []m.Profile
	db := db.DBConn

	result := db.Find(&prof, "employee_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&prof)
}

// 🥛 Create
func AddProfile(c *fiber.Ctx) error {
	var prof m.Profile
	db := db.DBConn

	if err := c.BodyParser(&prof); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// Generate a random 5-digit INET number
	inet := generateRandomInet()
	// Calculate age from birthday
	age := calculateAge(prof.Birthday)
	prof.Age = age
	// Set the emp_id to the generated INET number
	prof.EmployeeID = inet

	db.Create(&prof)
	return c.Status(fiber.StatusCreated).JSON(prof)
}

func calculateAge(birthday string) int {
	// Parse the birthday string into a time.Time object
	// formate YYYY-MM-DD
	birthdate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0
	}
	// Calculate the age by subtracting the birthdate from the current time
	age := int(time.Since(birthdate).Hours() / (24 * 365))

	return age
}

func generateRandomInet() string {
	// Generate a random 5-digit number
	randNum := rand.Intn(99999) + 10000

	// Convert the number to a string
	inet := strconv.Itoa(randNum)

	// Return the INET number
	return inet
}

// 🔱 Update
func UpdateProfile(c *fiber.Ctx) error {
	var prof m.Profile
	id := c.Params("id")
	db := db.DBConn

	if err := c.BodyParser(&prof); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&prof)
	return c.Status(200).JSON(prof)
}

// 🗑️ Delete
func RemoveProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.DBConn
	var prof m.Profile

	result := db.Delete(&prof, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetProfileJson(c *fiber.Ctx) error {
	db := db.DBConn
	var prof []m.Profile

	db.Find(&prof) // 10 ตัว

	var dataProResults []m.ProfileRes
	GenZ := 0
	GenY := 0
	GenX := 0
	BabyBoomer := 0
	GIGeneration := 0
	for _, v := range prof {
		typeStr := ""
		if v.Age < 24 {
			typeStr = "GenZ"
			GenZ++
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
			GenY++
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
			GenX++
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
			BabyBoomer++
		} else {
			typeStr = "G.I. Generation"
			GIGeneration++
		}

		d := m.ProfileRes{
			Name: v.Name,
			Age:  v.Age,
			Type: typeStr,
		}
		dataProResults = append(dataProResults, d)
	}

	r := m.ResultProfileData{
		Data:         dataProResults,
		Name:         "golang-test",
		Count:        len(dataProResults),
		GenZ:         GenZ,
		GenX:         GenX,
		GenY:         GenY,
		BabyBoomer:   BabyBoomer,
		GIGeneration: GIGeneration,
	}

	return c.Status(200).JSON(r)
}

func SearchProfile(c *fiber.Ctx) error {
	search := strings.TrimSpace(c.Query("search"))
	var prof []m.Profile
	db := db.DBConn

	result := db.Where("employee_id LIKE ? OR name LIKE ? OR last_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&prof)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&prof)
}
