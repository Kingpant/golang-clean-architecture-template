package router

import (
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/handler"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/middleware"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/request"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRouter(app *fiber.App, userHandler handler.UserHandler) {
	userRouter := app.Group("/user")

	userRouter.Get("/", userHandler.GetUsers)
	userRouter.Post("/", middleware.ValidateBody[request.CreateUserRequest](), userHandler.Create)
	userRouter.Patch("/email/:id", middleware.ValidateBody[request.UpdateUserEmailRequest](), userHandler.UpdateUserEmail)
}
