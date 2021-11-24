package middleware

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		signedStringToken, err := context.Cookie("jwt-token")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}

		if isValid := utility.ValidateSignedString(signedStringToken); !isValid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}
		
		context.Next()
	}
}