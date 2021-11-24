package middleware

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := context.Cookie("jwt-token")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}
		context.Next()
	}
}