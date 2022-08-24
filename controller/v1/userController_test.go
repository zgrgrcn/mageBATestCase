package v1

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/model/dto"
	"mageBATestCase/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	r := routes.InitRoute()
	req, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.ApiResponse
	json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert := assert.New(t)
	assert.Equal(http.StatusNotImplemented, w.Code, "Status code should be 501")
	assert.Equal("501", realResponse.Status, "In response object status code should be 501")
	assert.NotEmpty(realResponse.TimeStamp, "TimeStamp should not be empty")
	assert.NotEmpty(realResponse.Result, "Result should not be empty")
}

func TestLogin(t *testing.T) {

}
