package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "mageBATestCase/controller/v1"
	"mageBATestCase/docs"
	"mageBATestCase/middleware"
	"mageBATestCase/model/dto"
	"mageBATestCase/util"
	"net/http"
	"time"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger()) //I can use middleware.Logger()
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler)
	router.NoRoute(noRouteHandler())
	router.SetTrustedProxies(nil)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	setAuthRoute(router)
	return router
}

func setAuthRoute(router *gin.Engine) {
	v1 := router.Group(docs.SwaggerInfo.BasePath)

	v1.POST(util.PATH_USER+util.PATH_REGISTER, new(controller.UserController).Register)
	v1.POST(util.PATH_USER+util.PATH_LOGIN, new(controller.UserController).Login)

	authGroup := router.Group(util.GetEnvVariable("BASE_PATH"))
	authGroup.Use(middleware.Authentication())
	//authGroup.GET("/profile", userController.Endgame)
	authGroup.POST(util.PATH_ENDGAME, new(controller.EndGameController).Endgame)
	authGroup.GET(util.PATH_LEADERBOARD, new(controller.LeaderBoardController).Leaderboard)

}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseModel := dto.ApiResponse{
			Status:    "404",
			TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
			Result: gin.H{
				"endpoint": c.Request.URL.String(),
				"method":   c.Request.Method,
				"hint":     "You entered an invalid Page/Endpoint, visit http://localhost:9090/swagger/index.html to see the available routes",
			}}
		c.JSON(http.StatusNotFound, responseModel)
	}
}
