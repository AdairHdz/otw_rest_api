package middleware

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
)

func ServiceProviderAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		signedStringToken, err := context.Cookie("jwt-token")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}

		serviceProviderIDFromToken, err := utility.ExtractCustomClaims(signedStringToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}

		serviceProviderID := context.Param("serviceProviderId")

		if serviceProviderID == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Please make sure you provide a service provider ID in the URL",
			})
			return
		}

		if serviceProviderIDFromToken != serviceProviderID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "You have no permission to access or modify the requested resource",
			})
			return
		}

		context.Next()
	}
}