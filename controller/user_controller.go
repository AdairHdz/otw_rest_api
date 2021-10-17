package controller

import (	
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (UserController) Store() gin.HandlerFunc {
	return func(context *gin.Context) {				
		
		requestData := request.User{}
		err := context.BindJSON(&requestData)

		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}

		v := utility.NewValidator()				

		err = v.Struct(requestData)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Input",
				Message: "Please make sure you've entered the required fields in the specified format. For more details, check the API documentation",
			})
			return
		}		

		serviceRequester, serviceProvider, err := requestData.ToEntity()

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

		r := db.Where("email_address = ?", requestData.EmailAddress).First(&entity.Account{})
		if r.RowsAffected != 0 {
			context.JSON(452, response.ErrorResponse {
				Error: "Duplicated email",
				Message: "The email address you entered is already registered in our records",
			})
			return
		}

		var result *gorm.DB
		if serviceRequester != nil {
			result = db.Create(&serviceRequester)
		} else if serviceProvider != nil {
			result = db.Create(&serviceProvider)
		}						

		if result.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		context.Status(http.StatusOK)
	}
}