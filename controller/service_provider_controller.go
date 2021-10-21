package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceProviderController struct{}

func (ServiceProviderController) StoreImage() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerID")
		_, err := uuid.FromString(providerID)		
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Invalid UUID")
			return
		}
		
		path := "./images/" + providerID
		err = utility.CreateDirectory(path)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		serviceProvider := entity.ServiceProvider{}
		
		db, err := database.New()
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		r := db.Where("id = ?", providerID).Find(&serviceProvider)
		if r.RowsAffected == 0 {
				context.AbortWithStatusJSON(http.StatusNotFound, "There is not a service provider with the ID you provided.")
				return
		}		

		dirIsEmpty, err := utility.DirIsEmpty(path)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		file, noFileSentError := context.FormFile("image")
		if noFileSentError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "You didn't provide any file.")
			return
		}

		fileExtension := filepath.Ext(file.Filename)

		if !utility.IsImage(fileExtension) {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid image format. Please make sure your file has jpg, jpeg, or png extension")
			return
		}

		if !dirIsEmpty {
			pathOfImageToBeDeleted := path + "/" + serviceProvider.BusinessPicture
			os.Remove(pathOfImageToBeDeleted)
		}

		err = context.SaveUploadedFile(file, path+"/"+file.Filename)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		serviceProvider.BusinessPicture = file.Filename
		r = db.Model(&entity.ServiceProvider{}).Where("id = ?", serviceProvider.ID).Update("business_picture", serviceProvider.BusinessPicture)
		
		if r.Error != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}		

		context.Status(http.StatusOK)
	}
}

func (ServiceProviderController) GetWithId() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerID")
		_, err := uuid.FromString(providerID)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Invalid UUID")
			return
		}

		serviceProvider := entity.ServiceProvider{}

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db.Preload("User").Where("id = ?", providerID).Find(&serviceProvider)
		if r.RowsAffected == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There is not a service provider with the ID you provided.")
			return
		}		

		score := entity.Score{}
		averageScore := 0.00
		s := db.Where("id = ?", serviceProvider.User.ID).Find(&score)
		if s.RowsAffected != 0 {
			averageScore = score.AverageScore
		}

		var response []struct{
			ID string `json:"id"`
			Names string `json:"names"`
			Lastname string `json:"lastname"`
			BusinessPicture string `json:"businessPicture"`
			BusinessName string `json:"businessName"`
			AverageScore float64 `json:"averageScore"`
		}

		pathPicture := serviceProvider.ID + "/" + serviceProvider.BusinessPicture

		response = append(response, struct{ID string "json:\"id\""; Names string "json:\"names\""; Lastname string "json:\"lastname\""; BusinessPicture string "json:\"businessPicture\""; BusinessName string "json:\"businessName\""; AverageScore float64 "json:\"averageScore\"" }{ID: serviceProvider.ID, Names: serviceProvider.User.Names, Lastname: serviceProvider.User.Lastname, BusinessPicture: pathPicture, BusinessName: serviceProvider.BusinessName, AverageScore: averageScore ,},
		)

		context.JSON(http.StatusOK, response)
	}

}