package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/middleware"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

// /ap1/v1/endgame test
func TestEndGame_userUnauthorizedERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	req, _ := http.NewRequest("POST", "/api/v1/endgame", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	userUnauthorizedMessage := "Please login first"
	assert.Equal(http.StatusUnauthorized, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, userUnauthorizedMessage, "User unauthorized, need to give error message")
}

func TestEndGame_wrongTokenERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	req, _ := http.NewRequest("POST", "/api/v1/endgame", nil)
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

func TestEndGame_invalidTokenERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	req, _ := http.NewRequest("POST", "/api/v1/endgame", nil)
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

func TestEndGame_badRequestFormatERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	jsonValue, _ := json.Marshal("[\n\t{\n\t\t\"UserID\": \"63050093f0fb611425eb8741\",\n\t\t\"score\": 1\n\t},\n\t{\n\t\t\"UserID\": \"63050093f0fb611425eb8741\",\n\t\t\"score\": 5\n\t}\n]")
	req, _ := http.NewRequest("POST", "/api/v1/endgame", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QifQ.3Z-Plg8xL4cBHfD5BemsAuKIg-qcJ14wbe7hPE5IVXQ")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	badRequestFormatMessage := "json: cannot unmarshal string into Go value of type []entity.Player"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 400")
	assert.Equal(realResponse.Message, badRequestFormatMessage, "Bad request format, need to give error message")
}

func TestEndGame_notExistUserERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	playerList := [3]entity.Player{{
		UserID: "test",
		Score:  1,
	}, {
		UserID: "test2",
		Score:  2,
	}, {
		UserID: "test3",
		Score:  3,
	}}
	jsonValue, _ := json.Marshal(playerList)
	req, _ := http.NewRequest("POST", "/api/v1/endgame", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QifQ.3Z-Plg8xL4cBHfD5BemsAuKIg-qcJ14wbe7hPE5IVXQ")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	invalidUserIdMessage := "`user_id` is not valid or missing, error: the provided hex string is not a valid ObjectID"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 400")
	assert.Equal(realResponse.Message, invalidUserIdMessage, "Invalid user_id, need to give error message")
}

func TestEndGame_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.Use(middleware.Authentication())
	r.POST("/api/v1/endgame", new(EndGameController).Endgame)

	playerList := [2]entity.Player{{
		UserID: "63050093f0fb611425eb8741",
		Score:  1,
	}, {
		UserID: "63050093f0fb611425eb8741",
		Score:  2,
	}}
	jsonValue, _ := json.Marshal(playerList)
	req, _ := http.NewRequest("POST", "/api/v1/endgame", bytes.NewBuffer(jsonValue))
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
