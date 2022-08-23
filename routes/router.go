package routes

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/controller"
	"mageBATestCase/middleware"
	"mageBATestCase/util"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler)
	setAuthRoute(router)
	return router
}

func setAuthRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	router.POST(util.GetEnvVariable("BASE_PATH")+util.PATH_USER+util.PATH_REGISTER, new(controller.UserController).Register)
	router.POST(util.GetEnvVariable("BASE_PATH")+util.PATH_USER+util.PATH_LOGIN, new(controller.UserController).Login)

	authGroup := router.Group("/")
	authGroup.Use(middleware.Authentication())
	//authGroup.GET("/profile", userController.Endgame)
	authGroup.GET(util.GetEnvVariable("BASE_PATH")+util.PATH_ENDGAME, new(controller.EndGameController).Endgame)
	authGroup.GET(util.GetEnvVariable("BASE_PATH")+util.PATH_LEADERBOARD, new(controller.LeaderBoardController).Leaderboard)

}
