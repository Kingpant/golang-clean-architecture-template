package request

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
