package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "host=localhost port=5432 user=postgres password=omonov2006 dbname=forest sslmode=disable"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Not connected")
	}
	DB.AutoMigrate(&User{})

}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)

}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := new(User)
	DB.First(&user, id)
	if user.Email == "" {
		return ctx.Status(500).SendString("User not available")
	}

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return ctx.JSON(&user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return ctx.Status(500).SendString("user not available")
	}

	DB.Delete(&user)
	return ctx.SendString("User is deleted")
}

func SaveUser(ctx *fiber.Ctx) error {
	user := new(User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return ctx.JSON(&user)
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user User
	DB.Find(&user, id)
	return ctx.JSON(&user)
}
