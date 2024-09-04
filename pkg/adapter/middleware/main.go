package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, ginErr := range c.Errors {
		fmt.Println(ginErr)
	}

	if appEnv := os.Getenv("APP_ENV"); appEnv == "development" {
		// status -1 doesn't overwrite existing status code
		c.JSON(-1, c.Errors.JSON())
	}

	c.JSON(-1, "Error occured!!!")
}
