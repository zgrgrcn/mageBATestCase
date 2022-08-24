package v1

import (
	"github.com/gin-gonic/gin"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
	"mageBATestCase/service"
	"net/http"
	"time"
)

type EndGameController struct{}

// Endgame godoc
// @Summary record the end of the game as score
// @Description record the end of the game as core and return the leaderboard
// @Tags endgame
// @ID Endgame
// @Security bearerAuth
// @Consume application/json
// @Produce json
// @Param PlayerRequest body []entity.Player true "user_id and score List of the users"
// @Success 200 {object} dto.ApiResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /endgame [post]
func (controller *EndGameController) Endgame(c *gin.Context) {
	_, isTokenExist := c.Get("user")
	if !isTokenExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Please login first"})
		return
	}

	var playerList []entity.Player
	if err := c.ShouldBindJSON(&playerList); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	leaderboardService := service.LeaderboardService{}
	err := leaderboardService.ValidateUserList(playerList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	leaderboardService.PutResults(playerList)
	var leaderBoardList, _ = leaderboardService.FindAll()
	c.JSON(http.StatusOK, dto.ApiResponse{
		Status:    "200",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
		Result:    leaderBoardList,
	})
}
