package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/app/middlewares"
	"github.com/vaibhavvikas/stay-sorted/controllers/user_controller"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	usersRouter := baseRouter.Group("/users")

	userLoginRouter := usersRouter.Group("/login")
	userLoginRouter.Use(middlewares.DbMiddleware(db))
	userLoginRouter.Use(middlewares.JSONAppErrorReporter())
	{
		userLoginRouter.POST("", user_controller.UserControllerImpl{}.Login)
	}

	usersRouter.Use(middlewares.DbMiddleware(db))
	usersRouter.Use(middlewares.Authentication())
	usersRouter.Use(middlewares.JSONAppErrorReporter())
	{
		usersRouter.POST("/create", user_controller.UserControllerImpl{}.Create)
		usersRouter.GET("/:user_pid", user_controller.UserControllerImpl{}.GetUserByPid)
	}

	return router
}
