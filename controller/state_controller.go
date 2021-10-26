package controller

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type StateController struct{}

func (StateController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.New()
		if err != nil {
			c.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		var states []entity.State
		db.Find(&states)		

		var response []struct{
			ID string `json:"id"`
			Name string `json:"name"`
		}
		for _, state := range states {
			response = append(response, struct{ID string "json:\"id\""; Name string "json:\"name\""}{ID: state.ID,Name: state.Name,},
			)
		}
		
		c.JSON(http.StatusOK, response)
	}

}

func (StateController) IndexCities() gin.HandlerFunc {
	return func(context *gin.Context) {
		stateID := context.Param("stateId")
		_, err := uuid.FromString(stateID)
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
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

		var cities []entity.City
		r := db.Where("state_id = ?", stateID).Find(&cities)	
		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There is not a state with the ID you provided or he has no associates cities.",
			})
			return
		}		


		result := []response.City{}
		
		for _, city := range cities {
			result = append(result, struct{ID string "json:\"id\""; Name string "json:\"name\""}{ID: city.ID,Name: city.Name,},)
		}

		context.JSON(http.StatusOK, result)
	}

}