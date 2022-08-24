package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/middleware"
	"mageBATestCase/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

// /ap1/api/v1/leaderboard test
func TestLeaderBoard_userUnauthorizedERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.GET("/api/v1/leaderboard", new(LeaderBoardController).Leaderboard)

	req, _ := http.NewRequest("GET", "/api/v1/leaderboard", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	userUnauthorizedMessage := "Please login first"
	assert.Equal(http.StatusUnauthorized, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, userUnauthorizedMessage, "User unauthorized, need to give error message")
}

func TestLeaderBoard_wrongTokenERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.GET("/api/v1/leaderboard", new(LeaderBoardController).Leaderboard)

	req, _ := http.NewRequest("GET", "/api/v1/leaderboard", nil)
	req.Header.Set("Authorization", "Bearer zgrgrcniIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QifQ.3Z-Plg8xL4cBHfD5BemsAuKIg-qcJ14wbe7hPE5IVXQ")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	wrongJWTMessage := "invalid character 'Î' looking for beginning of value"
	assert.Equal(http.StatusUnauthorized, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, wrongJWTMessage, "Wrong JWT, need to give error message")
}

func TestLeaderBoard_invalidTokenERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.GET("/api/v1/leaderboard", new(LeaderBoardController).Leaderboard)

	req, _ := http.NewRequest("GET", "/api/v1/leaderboard", nil)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QifQ.3Z-Plg8xL4cBHfD5BemsAuKIg-qcJ14wbe7hPE5IVXQ")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	invalidTokenMessage := "Invalid token"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, invalidTokenMessage, "Invalid token, need to give error message")
}

func TestLeaderBoard_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.GET("/api/v1/leaderboard", new(LeaderBoardController).Leaderboard)

	req, _ := http.NewRequest("GET", "/api/v1/leaderboard", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QifQ.3Z-Plg8xL4cBHfD5BemsAuKIg-qcJ14wbe7hPE5IVXQ")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.ApiResponse
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	assert.Equal(http.StatusOK, w.Code, "Status Code should be 200")
	assert.Equal("200", realResponse.Status, "In response object status code should be 200")
	assert.NotEmpty(realResponse.TimeStamp, "TimeStamp should not be empty")
	assert.NotEmpty(realResponse.Result, "Result should not be empty")
}
