package response

type GetUsersResponse struct {
	UserIDs []string `json:"user_ids"`
	Users   []string `json:"users"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}
