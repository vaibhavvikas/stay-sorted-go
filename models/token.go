package models

type TokenResponse struct {
	Token   string       `json:"token"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}
