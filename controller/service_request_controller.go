package controller

import (
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type RequestController struct{}

func (RequestController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var serviceRequest request.Request

		err := context.BindJSON(&serviceRequest)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}		

		request, err := serviceRequest.ToEntity()
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

func (RequestController) GetWithId() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.Param("serviceRequestId")
		_, err := uuid.FromString(requestId)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		request := entity.ServiceRequest{}

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db.Preload("DeliveryAddress").Preload("ServiceProvider.User").
		Preload("ServiceRequester.User").Where("id = ?", requestId).Find(&request)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a request with the ID you provided.",
			})
			return
		}		
		result := response.ServiceRequest{}

		result = mapper.CreateRequestAsResponse(request)
		context.JSON(http.StatusOK, result)
	}

}

func (RequestController) IndexRequester() gin.HandlerFunc {
	return func(context *gin.Context) {
		requesterId := context.Param("serviceRequesterId")
		_, err := uuid.FromString(requesterId)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		date := context.Query("date")
		_, err = time.Parse("2006-01-02", date)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid date",
				Message: "The date format you provided is not valid.",
			})
			return
		}

		request := []entity.ServiceRequest{}

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").
		Preload("ServiceRequester.User").Where("service_requester_id = ?", requesterId).
		Where("date = ?", date).Find(&request)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a request with the serviceRequesterId or date you provided.",
			})
			return
		}		
		result := []response.ServiceRequestWithCity{}

		for _, request := range request {
			result = append(result, mapper.CreateRequestsAsResponse(request))
		}
		context.JSON(http.StatusOK, result)
	}

}

func (RequestController) IndexProvider() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerId := context.Param("serviceProviderId")
		_, err := uuid.FromString(providerId)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		date := context.Query("date")
		_, err = time.Parse("2006-01-02", date)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid date",
				Message: "The date format you provided is not valid.",
			})
			return
		}

		request := []entity.ServiceRequest{}

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").
		Preload("ServiceRequester.User").Where("service_provider_id = ?", providerId).
		Where("date = ?", date).Find(&request)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a request with the serviceProviderId or date you provided.",
			})
			return
		}		
		result := []response.ServiceRequestWithCity{}

		for _, request := range request {
			result = append(result, mapper.CreateRequestsAsResponse(request))
		}
		context.JSON(http.StatusOK, result)
	}

}