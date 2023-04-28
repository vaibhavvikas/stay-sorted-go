package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/services/user_service"
	"github.com/vaibhavvikas/stay-sorted/transformers"
	"github.com/vaibhavvikas/stay-sorted/validators/user_validator"
)

type UserControllerImpl struct {
	UserService user_service.UserServiceImpl
}

func (u UserControllerImpl) GetUserByPid(ctx *gin.Context) {
	userPid, err := user_validator.ValidateGetUserByPidRequest(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := u.UserService.GetUserByPid(ctx, userPid)
	if err != nil {
		ctx.Error(err)
		return
	}

	userResp := transformers.GetUser(user)
	ctx.IndentedJSON(http.StatusOK, userResp)
}

func (u UserControllerImpl) Create(ctx *gin.Context) {
	userReq, err := user_validator.ValidateCreateUserRequest(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := u.UserService.CreateUser(ctx, userReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	userResp := transformers.GetUser(user)
	ctx.IndentedJSON(http.StatusOK, userResp)
}

func (u UserControllerImpl) Login(ctx *gin.Context) {
	userReq, err := user_validator.ValidateLoginUserRequest(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	token, err := u.UserService.AuthenticateUser(ctx, userReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	tokenResp := transformers.GetTokenResponse(token, userReq)
	ctx.IndentedJSON(http.StatusOK, tokenResp)
}
