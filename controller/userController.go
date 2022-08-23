package controller

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/model/entity"
	"mageBATestCase/service"
	"time"
)

type UserController struct{}

func (auth *UserController) Register(c *gin.Context) {
	var userPayload entity.User
	if err := c.ShouldBindJSON(&userPayload); err != nil || userPayload.Password == "" || userPayload.Username == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Please input all fields"})
		return
	}
	dbUser, err := service.Userservice{}.Create(&userPayload)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, entity.ApiResponse{
			Status:    "200",
			TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
			Result:    dbUser,
		})
	}
	return
}

// Login is to process login request
func (auth *UserController) Login(c *gin.Context) {

	var loginInfo entity.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	//TODO
	userservice := service.Userservice{}
	dbUser, err := userservice.Find(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := dbUser.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	dbUser.Token = token
	c.JSON(200, entity.ApiResponse{
		Status:    "200",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
		Result:    dbUser,
	})
}
