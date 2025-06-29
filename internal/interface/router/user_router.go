package router

import (
	"github.com/Kingpant/golang-template/internal/interface/handler"
	"github.com/Kingpant/golang-template/internal/interface/middleware"
	"github.com/Kingpant/golang-template/internal/interface/request"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRouter(app *fiber.App, userHandler handler.UserHandler) {
	userRouter := app.Group("/user")

	userRouter.Get("/", userHandler.GetUsers)
	userRouter.Post("/", middleware.ValidateBody[request.CreateUserRequest](), userHandler.Create)
}
