package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoRouteHandler(t *testing.T) {
	assert := assert.New(t)
	r := gin.Default()
	r.GET("/", noRouteHandler())
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.ApiResponse
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	assert.Equal(http.StatusNotImplemented, w.Code, "Status code should be 501")
	assert.Equal("501", realResponse.Status, "In response object status code should be 501")
	assert.NotEmpty(realResponse.TimeStamp, "TimeStamp should not be empty")
	assert.NotEmpty(realResponse.Result, "Result should not be empty")
}

func TestInitRoute(t *testing.T) {
	assert := assert.New(t)
	r := InitRoute()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.ApiResponse
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	assert.Equal(http.StatusNotImplemented, w.Code, "Status code should be 501")
	assert.Equal("501", realResponse.Status, "In response object status code should be 501")
	assert.NotEmpty(realResponse.TimeStamp, "TimeStamp should not be empty")
	assert.NotEmpty(realResponse.Result, "Result should not be empty")
}
