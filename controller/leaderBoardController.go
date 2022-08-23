package controller

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/model/entity"
)

type LeaderBoardController struct{}

func (auth *LeaderBoardController) Leaderboard(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	c.JSON(200, gin.H{
		"Username": user.Username,
		"Password": user.Password,
	})
}
