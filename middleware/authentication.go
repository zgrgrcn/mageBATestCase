package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"mageBATestCase/model/dto"
	"mageBATestCase/service"
	"mageBATestCase/util"
	"net/http"
	"strings"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			//c.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			//	Message: "Authentication header is missing",
			//})
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		if len(temp) < 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{Message: "Invalid token"})
			return
		}
		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// }
			secretKey := util.GetEnvVariable("TOKEN_KEY")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Message: err.Error(),
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["Username"].(string)
			userservice := service.Userservice{}
			user, err := userservice.FindByUserName(username) //TODO: need userservice.FindByUserName
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, dto.Response{
					Message: "User not found",
				})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Message: "Token is not valid",
			})
			return
		}
	}
}
