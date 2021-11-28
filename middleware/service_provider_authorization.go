package middleware

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
)

func ServiceProviderAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		signedStringToken, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return utility.PublicKey, nil
		}, request.WithClaims(&utility.CustomClaims{}))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}

		claimsFromToken, err := utility.ExtractCustomClaims(signedStringToken.Raw)
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

		if claimsFromToken.SpecificID != serviceProviderID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "You have no permission to access or modify the requested resource",
			})
			return
		}

		context.Next()
	}
}