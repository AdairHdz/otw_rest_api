package middleware

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
)

func Authentication() gin.HandlerFunc {
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

		if !signedStringToken.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse {
				Error: "Unauthorized",
				Message: "Please authenticate to proceed",
			})
			return
		}
		
		context.Next()
	}
}