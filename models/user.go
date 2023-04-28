package models

import "github.com/vaibhavvikas/stay-sorted/models/entities"

type CreateUserReq struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Gender   string  `json:"gender"`
	Dob      string  `json:"dob"`
	Phone    string  `json:"phone"`
	Rating   float64 `json:"rating"`
	Password string  `json:"password"`
}

type UserResponse struct {
	Pid    string  `json:"user_pid"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Gender string  `json:"gender"`
	Dob    string  `json:"dob"`
	Phone  string  `json:"phone"`
	Rating float64 `json:"rating"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	User     entities.User
}
