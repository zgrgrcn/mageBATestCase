package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mageBATestCase/model/dto"
	"mageBATestCase/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

// /ap1/v1/user/register test
func TestRegister_userAlreadyExistsERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/register", new(UserController).Register)
	newUser := dto.UserRequest{
		Username: "test",
		Password: "test",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	userAlreadyExistsMessage := "write exception: write errors: [E11000 duplicate key error collection: mage.user index: username_1 dup key: { username: \"test\" }]"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 400")
	assert.Equal(realResponse.Message, userAlreadyExistsMessage, "User already exists, need to give error message")
}

func TestRegister_badRequestERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/register", new(UserController).Register)
	newUser := dto.UserRequest{
		Username: "",
		Password: "test",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	badRequestMessage := "Please input all fields"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 400")
	assert.Equal(realResponse.Message, badRequestMessage, "Bad request, need to give error message")
}

func TestRegister_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/register", new(UserController).Register)
	newUser := dto.UserRequest{
		Username: util.NameGenerator(10),
		Password: "123456",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonValue))
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

// /ap1/v1/user/login test
func TestLogin_wrongPasswordERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/login", new(UserController).Login)
	newUser := dto.UserRequest{
		Username: "ozgur2",
		Password: "1251rsr3tw34",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	wrongPasswordMessage := "password is not correct"
	assert.Equal(http.StatusUnauthorized, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, wrongPasswordMessage, "Wrong password, need to give error message")
}

func TestLogin_userNotExistERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/login", new(UserController).Login)
	newUser := dto.UserRequest{
		Username: "yok boyle user",
		Password: "1251rsr3tw34",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	userNotExistMessage := "there is no users with the same username"
	assert.Equal(http.StatusUnauthorized, w.Code, "Status Code should be 401")
	assert.Equal(realResponse.Message, userNotExistMessage, "User not exist, need to give error message")
}

func TestLogin_badRequestERROR(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/login", new(UserController).Login)

	req, _ := http.NewRequest("POST", "/v1/user/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var realResponse dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &realResponse)
	assert.Nil(err, "Error should be nil")

	badRequesttMessage := "invalid request"
	assert.Equal(http.StatusBadRequest, w.Code, "Status Code should be 400")
	assert.Equal(realResponse.Message, badRequesttMessage, "Bad request, need to give error message")
}

func TestLogin_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	r := gin.Default()
	r.POST("/v1/user/login", new(UserController).Login)
	newUser := dto.UserRequest{
		Username: "test",
		Password: "test",
	}
	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonValue))
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
