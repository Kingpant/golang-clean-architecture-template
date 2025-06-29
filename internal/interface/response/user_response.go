package response

type GetUsersResponse struct {
	Users []string `json:"users"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}
