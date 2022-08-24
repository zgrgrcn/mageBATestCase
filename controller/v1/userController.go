package v1

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/middleware"
	"mageBATestCase/model/dto"
	"mageBATestCase/service"
	"net/http"
	"time"
)

type UserController struct{}

// Paths Information

// Register godoc
// @Summary register a new user
// @Description register a new user with username and password
// @Tags user
// @ID Register
// @Consume application/json
// @Produce json
// @Param UserRequest body dto.UserRequest true "username and password"
// @Success 200 {object} dto.ApiResponse
// @Failure 400 {object} dto.Response
// @Router /user/register [post]
func (controller *UserController) Register(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil || userRequest.Password == "" || userRequest.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Message: "Please input all fields",
		})
		return
	}
	dbUser, err := service.UserService{}.Create(&userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, dto.ApiResponse{
			Status:    "200",
			TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
			Result:    dbUser,
		})
	}
	return
}

// Login godoc
// @Summary Login a user
// @Description Login a new user with username and password and return a token
// @Tags user
// @ID Login
// @Consume application/json
// @Produce json
// @Param UserRequest body dto.UserRequest true "username and password"
// @Success 200 {object} dto.ApiResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /user/login [post]
func (controller *UserController) Login(c *gin.Context) {
	middleware.Authentication()(c)
	//(entity.User)
	var loginInfo dto.UserRequest
	user, isTokenExist := c.Get("user")
	if isTokenExist {
		loginInfo = user.(dto.UserRequest)
	} else {
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
			return
		}
	}
	userservice := service.UserService{}
	dbUser, err := userservice.Find(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: err.Error()})
		return
	}

	token, err := dbUser.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}
	dbUser.Token = token
	c.JSON(http.StatusOK, dto.ApiResponse{
		Status:    "200",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
		Result:    dbUser,
	})
}
