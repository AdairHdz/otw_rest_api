package controller

import (	
	"net/http"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/request"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	uuid "github.com/satori/go.uuid"
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

		verificationCode := utility.GenerateCode()

		var result *gorm.DB
		if serviceRequester != nil {
			serviceRequester.User.Account.VerificationCode = verificationCode
			result = db.Create(&serviceRequester)
		} else if serviceProvider != nil {
			serviceProvider.User.Account.VerificationCode = verificationCode
			result = db.Create(&serviceProvider)
		}						

		if result.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		go utility.SendToEmail(requestData.EmailAddress, verificationCode)		

		res := response.User{}
		if serviceRequester != nil {
			res = mapper.CreateRequesterAddAsResponse(serviceRequester)
		} else if serviceProvider != nil {
			res = mapper.CreateProviderAddAsResponse(serviceProvider)
		}	
		context.JSON(http.StatusOK, res)
	}
}

func (UserController) SendEmail() gin.HandlerFunc {
	return func(context *gin.Context) {	
		userId := context.Param("userId")
		_, err := uuid.FromString(userId)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		requestData := struct {
			EmailAddress string `json:"emailAddress" validate:"required,email,max=254"`
		}{}

		err = context.BindJSON(&requestData)
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
		
		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		verificationCode := utility.GenerateCode()

		r := db.Model(&entity.Account{}).Where("email_address = ?", requestData.EmailAddress).
			Where("user_id = ?", userId).Where("verified = false").Update("verification_code", verificationCode)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There are no users with that id, email or he is already verified.",
			})
			return
		}

		go utility.SendToEmail(requestData.EmailAddress, verificationCode)		
	
		context.Status(http.StatusNoContent)
	}
}

func (UserController) Verify() gin.HandlerFunc {
	return func(context *gin.Context) {	
		userId := context.Param("userId")
		_, err := uuid.FromString(userId)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		requestData := struct {
			Code string `json:"verificationCode" validate:"required,max=8"`
		}{}

		err = context.BindJSON(&requestData)
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
		
		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db.Model(&entity.Account{}).Where("verification_code = ?", requestData.Code).
			Where("user_id = ?", userId).Update("verified", true)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There are no users with that id or the verification code is incorrect.",
			})
			return
		}
	
		context.Status(http.StatusNoContent)
	}
}
