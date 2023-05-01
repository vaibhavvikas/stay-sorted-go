package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/app/middleware"
	"github.com/vaibhavvikas/stay-sorted/controllers/house_controller"
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
	baseRouter.Use(middleware.DbMiddleware(db))
	baseRouter.Use(middleware.JSONAppErrorReporter())

	usersRouter := baseRouter.Group("/users")
	{
		usersRouter.POST("/signup", user_controller.UserControllerImpl{}.Create)
		usersRouter.POST("/login", user_controller.UserControllerImpl{}.Login)
		usersRouter.GET("/:user_pid", user_controller.UserControllerImpl{}.GetUserByPid)
	}

	houseRouter := baseRouter.Group("/houses")
	{
		houseRouter.GET("", house_controller.HouseControllerImpl{}.GetAllHouses)
		houseRouter.GET("/filter", house_controller.HouseControllerImpl{}.GetFilteredHouses)
		houseRouter.GET("/:house_pid", house_controller.HouseControllerImpl{}.GetHouse)
	}
	houseRouter.Use(middleware.Authentication())
	{
		houseRouter.POST("/create", house_controller.HouseControllerImpl{}.CreateHouse)
		// houseRouter.POST("/update/:house_pid")
		// houseRouter.DELETE("/:house_pid")
	}

	// userLoginRouter := usersRouter.Group("")
	// userLoginRouter.Use(middleware.DbMiddleware(db))
	// usersRouter.Use(middleware.Authentication())
	// userLoginRouter.Use(middleware.JSONAppErrorReporter())
	// {
	// 	userLoginRouter.POST("/signup", user_controller.UserControllerImpl{}.Create)
	// 	userLoginRouter.POST("/login", user_controller.UserControllerImpl{}.Login)
	// }

	return router
}
