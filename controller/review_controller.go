package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ReviewController struct{}

func (ReviewController) GetWithId() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerId := context.Param("serviceProviderId")
		_, err := uuid.FromString(providerId)

		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		page, err := strconv.Atoi(context.Query("page"))
		if err != nil || page > 50 || page <= 0 {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid page parameter.",
			})
			return
		}

		pageElements, err := strconv.Atoi(context.Query("pageSize"))
		if err != nil || pageElements > 50 || pageElements <= 0 {
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

		r := db.Scopes(utility.Paginate(page, pageElements)).Preload("ServiceRequester.User").Preload("Evidences").
		Where("service_provider_id = ?", providerId).Find(&reviews)
		
		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There is not a service provider with the ID you provided or he has no reviews.",
			})
			return
		}	
		
		var rowCount int64
		db.Model(&entity.Review{}).Where("service_provider_id = ?", providerId).Count(&rowCount)

		result := []response.ReviewWithEvidence{}
		
		for _, review := range reviews {
			result = append(result, mapper.CreateReviewWithEvidenceAsResponse(review))
		}

		round := float64(rowCount) / float64(pageElements)
		
		lastPage := int(utility.Roundf(round))
		
		var previousPage int = 1
		var nextPage int

		if page-1 > 1 {
			previousPage = page - 1
		}

		if page+1 <= lastPage {
			nextPage = page + 1
		} else {
			nextPage = lastPage
		}

		dataResponse := struct {
			Links struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			} `json:"links"`
			Page    int                                    `json:"page"`
			Pages   int                                    `json:"pages"`
			PerPage int                                    `json:"perPage"`
			Total   int64                                  `json:"total"`
			Data    []response.ReviewWithEvidence		   `json:"data"`
		}{
			Links: struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			}{
				First: fmt.Sprintf("providers/%s/reviews?page=%d&pageSize=%d", providerId, 1, pageElements),
				Last:  fmt.Sprintf("providers/%s/reviews?page=%d&pageSize=%d", providerId, lastPage, pageElements),
				Prev:  fmt.Sprintf("providers/%s/reviews?page=%d&pageSize=%d", providerId, previousPage, pageElements),
				Next:  fmt.Sprintf("providers/%s/reviews?page=%d&pageSize=%d", providerId, nextPage, pageElements),
			},
			Page:    page,
			Pages:   lastPage,
			PerPage: pageElements,
			Total:   rowCount,
			Data:    result,
		}
		context.JSON(http.StatusOK, dataResponse)

	}

}

func (ReviewController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("serviceProviderId")
		_, err := uuid.FromString(providerID)

		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		var reviewBody request.Review
		err = context.BindJSON(&reviewBody)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Please make sure you have entered the required fields in a valid format",
			})
			return
		}

		validator := utility.NewValidator()
		err = validator.Struct(reviewBody)
		
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Please make sure you have entered the required fields ina valid format",
			})
			return
		}

		db, err := database.New()

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		hasBeenReviwed := db.Preload("ServiceRequest", "has_been_reviewed = ? AND id = ?", true, reviewBody.ServiceRequestID).First(&entity.Review{})
		if hasBeenReviwed.RowsAffected != 0 {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Service already reviewed",
				Message: "The service request has already been reviewed by the service requester",
			})
			return
		}

		reviewEntity := reviewBody.ToEntity(providerID)
		r := db.Create(&reviewEntity)
		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		reviewRequest := db.Model(&entity.ServiceRequest{}).Where("id = ?", reviewEntity.ServiceRequestID).Update( "has_been_reviewed" , true)
		if reviewRequest.RowsAffected == 0 {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}	

		response := mapper.CreateReviewWithRequesterIDAsResponse(reviewEntity)
		context.JSON(http.StatusOK, response)

	}
}

func (ReviewController) UploadEvidence() gin.HandlerFunc {
	return func(context *gin.Context) {
		var reviewId string = context.Param("reviewId")
		const maxFileSize = 10855731
		form, _ := context.MultipartForm()
		files := form.File["evidence[]"]
		path := fmt.Sprintf("./public/reviews/%s", reviewId)
		directoryCreationError := utility.CreateDirectory(path)

		if directoryCreationError != nil {
			println(directoryCreationError.Error())
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			return
		}

		dirIsEmpty, directoryEmptinessVerificationError := utility.DirIsEmpty(path)

		if directoryEmptinessVerificationError != nil {
			println(directoryEmptinessVerificationError.Error())
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			return
		}

		if !dirIsEmpty {
			println("Attempted to add files to a review that already has files registered")
			context.AbortWithStatusJSON(http.StatusConflict, "Attempted to add files to a review that already has files registered")
			return
		}

		if len(files) == 0 {
			println("Request should contain at least one file")
			context.AbortWithStatusJSON(http.StatusBadRequest, "Request should contain at least one file")
			return
		} else if len(files) > 3 {
			println("Can't upload more than 3 files per request")
			context.AbortWithStatusJSON(http.StatusBadRequest, "Can't upload more than 3 files per request")
			return
		}

		for _, file := range files {
			var fileSizeTotal int64 = file.Size
			if fileSizeTotal > maxFileSize {
				println("One or more files have a size greater than 10 MB")
				context.AbortWithStatusJSON(http.StatusConflict, "One or more files have a size greater than 10 MB")
				return
			}
			fileExtension := filepath.Ext(file.Filename)
			if !utility.EvidenceHasValidFormat(fileExtension) {
				println("One or more files have invalid format")
				context.AbortWithStatusJSON(http.StatusBadRequest, "One or more files have invalid format")
				return
			}
		}

		for _, file := range files {
			fileSavingError := context.SaveUploadedFile(file, path+"/"+file.Filename)
			if fileSavingError != nil {
				println("There was an error while trying to save the evidence", fileSavingError.Error())
				context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			}
		}

		context.Status(http.StatusCreated)
	}
}
