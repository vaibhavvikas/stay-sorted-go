package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/app/middleware"
	"github.com/vaibhavvikas/stay-sorted/controllers/user_controller"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	usersRouter := baseRouter.Group("/users")

	userLoginRouter := usersRouter.Group("")
	userLoginRouter.Use(middleware.DbMiddleware(db))
	userLoginRouter.Use(middleware.JSONAppErrorReporter())
	{
		userLoginRouter.POST("/create", user_controller.UserControllerImpl{}.Create)
		userLoginRouter.POST("/login", user_controller.UserControllerImpl{}.Login)
	}

	usersRouter.Use(middleware.DbMiddleware(db))
	usersRouter.Use(middleware.Authentication())
	usersRouter.Use(middleware.JSONAppErrorReporter())
	{
		usersRouter.GET("/:user_pid", user_controller.UserControllerImpl{}.GetUserByPid)
	}

	return router
}
