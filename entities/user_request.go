package entities

type CreateUserRequest struct {
	Username string `json:"username"`
}

type GetUserRequest struct {
	Username string `json:"username"`
}
