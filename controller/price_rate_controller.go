package controller

import (
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type FiltersKindOfServiceAndCity struct{
	KindOfService int
	CityID string
	ServiceProviderID string
} 

type PriceRateController struct{}

func (PriceRateController) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerID")
		_, err := uuid.FromString(providerID)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Invalid UUID")
			return
		}

		kindOfService := 0
		if context.Query("kindOfService") != "" {
			kindOfService, err = strconv.Atoi(context.Query("kindOfService"))
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid kind of service parameter")
				return
			}
		} 

		cityID := context.Query("cityID")

		filters := &FiltersKindOfServiceAndCity{
			KindOfService: kindOfService,
			CityID: cityID,
			ServiceProviderID: providerID,
		}

		var priceRates []entity.PriceRate

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		
		r := db.Preload("City").Preload("WorkingDays").Where(filters).Find(&priceRates)
		if r.RowsAffected == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There is not a service provider with the ID you provided or he does not have rates with the filters entered.")
			return
		}	

		result := []mapper.PriceRateWorkingDays{}
		
		for _, priceRate := range priceRates {
			result = append(result, mapper.CreatePriceRateWorkingDaysAsResponse(priceRate))
		}

		context.JSON(http.StatusOK, result)

	}

}