package controller

import (
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type AddressController struct{}

func (AddressController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var address request.Address

		requesterID := context.Param("serviceRequesterId")
		_, err := uuid.FromString(requesterID)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		err = context.BindJSON(&address)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}
		
		v := utility.NewValidator()				

		err = v.Struct(address)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}	

		addressEntity, err := address.ToEntity(requesterID)

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

		result := db.Create(&addressEntity)

		if result.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := mapper.CreateAddressAsResponse(addressEntity)
		context.JSON(http.StatusOK, r)
	}
}

func (AddressController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {

		requester_id := context.Param("serviceRequesterId")
		_, err := uuid.FromString(requester_id)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The RequesterId you provided has an invalid format",
			})
			return
		}

		city_id := context.Query("cityId")
		_, err = uuid.FromString(city_id)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The CityId you provided has an invalid format",
			})
			return
		}

		var addresses []entity.Address

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r:= db.Where("city_id=? AND service_requester_id=?", city_id, requester_id).Find(&addresses)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There are no addresses in that city that belong to that service requester.",
			})
			return
		}	
		
		result := []response.Address{}
		
		for _, address := range addresses {
			result = append(result, mapper.CreateAddressesAsResponse(address))
		}
		context.JSON(http.StatusOK, result)
	}
}