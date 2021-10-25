package controller

import (
	"net/http"
	"strconv"
	"time"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		kindOfService := 0
		if context.Query("kindOfService") != "" {
			kindOfService, err = strconv.Atoi(context.Query("kindOfService"))
			if err != nil {
				context.JSON(http.StatusBadRequest, response.ErrorResponse {
					Error: "Bad Request",
					Message: "Invalid kind of service parameter",
				})
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
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a service provider with the ID you provided or he does not have rates with the filters entered.",
			})
			return
		}	

		result := []mapper.PriceRateWorkingDays{}
		
		for _, priceRate := range priceRates {
			result = append(result, mapper.CreatePriceRateWorkingDaysAsResponse(priceRate))
		}

		context.JSON(http.StatusOK, result)

	}
}

func (PriceRateController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var priceRate request.PriceRate

		providerID := context.Param("providerID")

		_, err := uuid.FromString(providerID)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		err = context.BindJSON(&priceRate)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you send valid data",
			})
			return
		}

		e, err := priceRate.ToEntity(providerID)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		startingHour, err := time.Parse("15:04", e.StartingHour)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		endingHour, err := time.Parse("15:04", e.EndingHour)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		if startingHour.Equal(endingHour) || startingHour.After(endingHour) {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "The starting hour cannot occur neither before nor at the same time than the ending hour",
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

		var matchingPriceRates []entity.PriceRate
		r := db.Where(&entity.PriceRate{WorkingDays: e.WorkingDays, KindOfService: e.KindOfService, CityID: e.CityID, ServiceProviderID: providerID}).Find(&matchingPriceRates)
		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		if len(matchingPriceRates) > 0 {
			collides, err := priceRatesCollide(priceRate, matchingPriceRates)
			if err != nil {
				context.JSON(http.StatusConflict, response.ErrorResponse {
					Error: "Internal Error",
					Message: "There was an unexpected error while processing your data. Please try again later",
				})
				return
			}

			if collides {
				context.JSON(http.StatusConflict, response.ErrorResponse {
					Error: "Colliding Price Rate",
					Message: "You already have a price rate that applies for the criteria you established",
				})
				return
			}
		}

		r = db.Create(&e)
		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}
		context.Status(http.StatusOK)
	}
}

func priceRatesCollide(newPriceRate request.PriceRate, priceRates []entity.PriceRate) (collides bool, err error) {
	var startingNew time.Time
	var startingHour time.Time
	var endingHour time.Time

	startingNew, err = time.Parse("15:04", newPriceRate.StartingHour)
	for _, p := range priceRates {
		startingHour, err = time.Parse("15:04", p.StartingHour)
		if err != nil {
			return
		}

		endingHour, err = time.Parse("15:04", p.EndingHour)
		if err != nil {
			return
		}

		if (startingNew.After(startingHour) || startingNew.Equal(startingHour)) && (startingNew.Before(endingHour) || startingNew.Equal(endingHour)) {
			collides = true
			return
		}

	}
	return
}