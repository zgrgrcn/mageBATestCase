package v1

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
	"mageBATestCase/service"
	"net/http"
	"time"
)

type LeaderBoardController struct{}

// Leaderboard godoc
// @Summary Get Leaderboard
// @Description Get Leaderboard with username and score
// @Tags leaderboard
// @ID Leaderboard
// @Security bearerAuth
// @Produce json
// @Success 200 {object} dto.ApiResponse
// @Failure 401 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /leaderboard [get]
func (controller *LeaderBoardController) Leaderboard(c *gin.Context) {
	_, isTokenExist := c.Get("user")
	if !isTokenExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: "Please login first"})
		return
	}

	var leaderBoardList []entity.Player

	leaderboardService := service.LeaderboardService{}
	leaderBoardList, _ = leaderboardService.FindAll()
	if len(leaderBoardList) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, dto.Response{Message: "There is no leaderboard yet, please wait for the end of the game"})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Status:    "200",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
		Result:    leaderBoardList,
	})
}
