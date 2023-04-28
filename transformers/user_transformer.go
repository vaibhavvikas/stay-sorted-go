package transformers

import (
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
)

func GetUser(user entities.User) models.UserResponse {
	return models.UserResponse{
		Pid:    user.Pid,
		Name:   user.Name,
		Email:  user.Email,
		Gender: user.Gender,
		Phone:  user.Phone,
		Dob:    user.Dob,
		Rating: user.Rating,
	}
}

func GetTokenResponse(token string, userReq models.UserLoginReq) models.TokenResponse {
	return models.TokenResponse{
		Token:   token,
		Data:    GetUser(userReq.User),
		Message: "Token Generated Successfully",
	}
}
