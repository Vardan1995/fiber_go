package controller

import (
	"github.com/Vardan1995/fiber-crud/database"
	"github.com/Vardan1995/fiber-crud/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := *database.DB

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	//todo hash pass
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}
	user.Password = string(bytes)
	

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "unable create user", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User has been created successfully"})
}

func GetUser(c *fiber.Ctx)error{
	id:=c.Params("id")
	db:= database.DB
	var user *model.User

	if err:=db.First(&user, id).Error;err != nil {
		return c.Status(404).JSON(fiber.Map{"message":"user specifide id not found"})
	}
 isAdmin := c.Locals("user")
	return c.Status(200).JSON(fiber.Map{"data":user,"user":isAdmin})


}
func GetUsers(c *fiber.Ctx)error{
	db:= database.DB
	var users []*model.User

	if err:=db.Find(&users).Error;err != nil {
		return c.Status(404).JSON(fiber.Map{"message":"users not found"})
	}
 
	return c.Status(200).JSON(fiber.Map{"data":users})


}
 func Pong(c *fiber.Ctx)error{

	
	 return c.SendString("pong")


}

func Login(c *fiber.Ctx) error {
	db := database.DB
	var login struct {
		Id       uint
		Password string
	}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var user *model.User

	if err := db.Where("id=?", login.Id).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "user not found", "data": err})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if  err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "wrong password", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "User found", "data": user})
}
