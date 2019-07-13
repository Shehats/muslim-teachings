package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// AuthenicationMiddleware authenticates users
func AuthenicationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if reqMethod := context.Request.Method; reqMethod == "GET" {
			fmt.Println("In middleware")
		}
	}
}
