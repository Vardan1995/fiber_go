package controller

import "github.com/gofiber/fiber/v2"

func SendResponse(err error,message string , data interface{})fiber.Map {
	var status = "success"
	if err != nil {
		status = "error"
	}
	return	fiber.Map{"status": status, "message": message, "data": err}
}