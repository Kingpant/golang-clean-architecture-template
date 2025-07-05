package handler

import (
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/middleware"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/request"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/response"
	"github.com/Kingpant/golang-clean-architecture-template/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	UpdateUserEmail(c *fiber.Ctx) error
}

type userHandler struct {
	userUseCase usecase.UserUsecase
}

func NewUserHandler(userUseCase usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

// GetUsers godoc
//
//	@Tags		User
//	@Summary	Get all users
//	@Produce	json
//	@Success	200	{object}	response.BaseResponse{data=response.GetUsersResponse}
//	@Failure	500	{object}	response.ErrorResponse
//	@Router		/user [get]
func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, userIDs, getUsersErr := h.userUseCase.GetUsers(c.Context())
	if getUsersErr != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, getUsersErr.Error())
	}

	return response.JSONOK(c, response.GetUsersResponse{Users: users, UserIDs: userIDs})
}

// CreateUser godoc
//
//	@Tags		User
//	@Summary	Create a new user
//	@Accept		json
//	@Produce	json
//	@Param		user	body		request.CreateUserRequest	true	"User Info"
//	@Success	200		{object}	response.BaseResponse{data=response.CreateUserResponse}
//	@Failure	500		{object}	response.ErrorResponse
//	@Router		/user [post]
func (h *userHandler) Create(c *fiber.Ctx) error {
	body := c.Locals(middleware.ValidatedBodyKey).(request.CreateUserRequest)

	userID, createUserErr := h.userUseCase.CreateUser(c.Context(), body.Name, body.Email)
	if createUserErr != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, createUserErr.Error())
	}

	return response.JSONOK(c, response.CreateUserResponse{
		UserID: userID,
	})
}

// UpdateUserEmail godoc
//
//	@Tags		User
//	@Summary	Update user email
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string							true	"User ID"
//	@Param		user	body		request.UpdateUserEmailRequest	true	"User Email Info"
//	@Success	200		{object}	response.BaseResponse
//	@Failure	500		{object}	response.ErrorResponse
//	@Router		/user/email/{id} [patch]
func (h *userHandler) UpdateUserEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	body := c.Locals(middleware.ValidatedBodyKey).(request.UpdateUserEmailRequest)

	updateErr := h.userUseCase.UpdateUserEmail(c.Context(), id, body.Email)
	if updateErr != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, updateErr.Error())
	}

	return response.JSONOK(c, nil)
}
