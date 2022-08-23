package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
