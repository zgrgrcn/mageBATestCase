package controller

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/model/entity"
)

type EndGameController struct{}

// Endgame is to provide current user info
func (auth *EndGameController) Endgame(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	c.JSON(200, gin.H{
		"Username": user.Username,
		"Password": user.Password,
	})
}
