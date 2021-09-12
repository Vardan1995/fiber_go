package router

import (
	"github.com/Vardan1995/fiber-crud/controller"
	"github.com/Vardan1995/fiber-crud/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App){
// api:=app.Group("/api",logger.New())
app.Get("/ping" , controller.Pong)
api:=app.Group("/api")
user := api.Group("/user")
user.Post("/", controller.CreateUser)
user.Get("/" , controller.GetUsers)
user.Get("/:id" ,middleware.Protected(), controller.GetUser)
user.Post("/login", controller.Login)

}