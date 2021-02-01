package entities

type GetUserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type CreateUserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
