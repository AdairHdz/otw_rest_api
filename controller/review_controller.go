package controller

import (
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type ReviewController struct{}

func (ReviewController) GetWithId() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerID")
		_, err := uuid.FromString(providerID)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		page, err := strconv.Atoi(context.Query("page"))
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid page parameter.",
			})
			return
		}

		pageElements, err := strconv.Atoi(context.Query("pageElements"))
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid page elements parameter.",
			})
			return
		}

		var reviews []entity.Review

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		
		r := db.Scopes(utility.Paginate(page, pageElements)).Preload("ServiceRequester.User").Preload("Evidences").Where("service_provider_id = ?", providerID).Find(&reviews)
		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There is not a service provider with the ID you provided or he has no reviews.",
			})
			return
		}	

		result := []response.ReviewWithEvidence{}
		
		for _, review := range reviews {
			result = append(result, mapper.CreateReviewWithEvidenceAsResponse(review))
		}

		context.JSON(http.StatusOK, result)

	}

}