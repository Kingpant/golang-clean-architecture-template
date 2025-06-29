package handler

import (
	"github.com/Kingpant/golang-template/internal/interface/request"
	"github.com/Kingpant/golang-template/internal/interface/response"
	"github.com/Kingpant/golang-template/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type userHandler struct {
	userUseCase usecase.UserUsecase
}

func NewUserHandler(userUseCase usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, getUsersErr := h.userUseCase.GetUsers(c.Context())
	if getUsersErr != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, getUsersErr.Error())
	}

	return response.JSONOK(c, response.GetUsersResponse{Users: users})
}

func (h *userHandler) Create(c *fiber.Ctx) error {
	body := c.Locals("validatedBody").(request.CreateUserRequest)

	user, createUserErr := h.userUseCase.CreateUser(c.Context(), body.Name, body.Email)
	if createUserErr != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, createUserErr.Error())
	}

	return response.JSONOK(c, user)
}
