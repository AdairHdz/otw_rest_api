package controller

import (
	"net/http"

	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/gin-gonic/gin"
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