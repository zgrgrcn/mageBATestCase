package v1

import (
	_ "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
	"mageBATestCase/routes"
	"mageBATestCase/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	assert := assert.New(t)
	r := routes.InitRoute()

	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []Company
	json.Unmarshal(w.Body.Bytes(), &companies)
}

func TestEndgame(t *testing.T) {
	_, isTokenExist := c.Get("user")
	if !isTokenExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: "Please login first"})
		return
	}

	var playerList []entity.Player
	if err := c.ShouldBindJSON(&playerList); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
		return
	}
	leaderboardService := service.LeaderboardService{}
	err := leaderboardService.ValidateUserList(playerList)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
		return
	}
	err = leaderboardService.PutResults(playerList)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
		return
	}
	var leaderBoardList, _ = leaderboardService.FindAll()
	c.JSON(http.StatusOK, dto.ApiResponse{
		Status:    "200",
		TimeStamp: time.Now().Format("2006/01/02 15:04:05"),
		Result:    leaderBoardList,
	})
}
