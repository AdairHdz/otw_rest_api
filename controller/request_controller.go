package controller

import (
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

type RequestController struct{}

func (RequestController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var service_request request.Request

		err := context.BindJSON(&service_request)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}		

		request, err := service_request.ToEntity()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: err.Error(),
			})
			return
		}

		db, err := database.New()

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		
		t := time.Now()
		request.Date = fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
		request.HasBeenReviewed = false
		request.Status = 1

		result := db.Create(&request)

		if result.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		context.Status(http.StatusNoContent)
	}
}