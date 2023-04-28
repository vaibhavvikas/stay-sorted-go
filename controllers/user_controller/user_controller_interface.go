package user_controller

import (
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Create(ctx *gin.Context)
	GetUserByPid(ctx *gin.Context)
}
