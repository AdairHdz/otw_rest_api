package controller

import (	
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"fmt"
	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/mapper"
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/utility"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/kennygrant/sanitize"
)

type ServiceProviderController struct{}

func (ServiceProviderController) StoreImage() gin.HandlerFunc {
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
		
		path := "./images/" + providerID
		err = utility.CreateDirectory(path)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an error while trying to save your image.",
			})
			return
		}

		serviceProvider := entity.ServiceProvider{}
		
		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an error while trying to save your image.",
			})
			return
		}

		r := db.Where("id = ?", providerID).Find(&serviceProvider)
		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a service provider with the ID you provided.",
			})
			return
		}		

		dirIsEmpty, err := utility.DirIsEmpty(path)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an error while trying to save your image.",
			})
			return
		}

		file, noFileSentError := context.FormFile("image")
		if noFileSentError != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "You didn't provide any file.",
			})
			return
		}

		fileExtension := filepath.Ext(file.Filename)

		if !utility.IsImage(fileExtension) {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid image format. Please make sure your file has jpg, jpeg, or png extension",
			})
			return
		}

		if !dirIsEmpty {
			pathOfImageToBeDeleted := path + "/" + serviceProvider.BusinessPicture
			os.Remove(pathOfImageToBeDeleted)
		}

		sanitizedFileName := sanitize.Name(file.Filename)
		err = context.SaveUploadedFile(file, path+"/"+ sanitizedFileName)

		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an error while trying to save your image.",
			})
			return
		}

		serviceProvider.BusinessPicture = sanitizedFileName
		r = db.Model(&entity.ServiceProvider{}).Where("id = ?", serviceProvider.ID).Update("business_picture", serviceProvider.BusinessPicture)
		
		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an error while trying to save your image.",
			})
			return
		}		

		context.Status(http.StatusOK)
	}
}

func (ServiceProviderController) GetWithId() gin.HandlerFunc {
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
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "There is not a service provider with the ID you provided.",
			})
			return
		}		

		score := entity.Score{}
		averageScore := 0.00
		s := db.Where("user_id = ?", serviceProvider.User.ID).Find(&score)
		if s.RowsAffected != 0 {
			averageScore = score.AverageScore
		}

		var response struct{
			ID string `json:"id"`
			Names string `json:"names"`
			Lastname string `json:"lastname"`
			BusinessPicture string `json:"businessPicture"`
			BusinessName string `json:"businessName"`
			AverageScore float64 `json:"averageScore"`
		}

		pathPicture := serviceProvider.ID + "/" + serviceProvider.BusinessPicture

		response = struct{ID string "json:\"id\""; Names string "json:\"names\""; Lastname string "json:\"lastname\""; BusinessPicture string "json:\"businessPicture\""; BusinessName string "json:\"businessName\""; AverageScore float64 "json:\"averageScore\""}{			
			ID: serviceProvider.ID,
			Names: serviceProvider.User.Names,
			Lastname: serviceProvider.User.Lastname,
			BusinessPicture: pathPicture,
			BusinessName: serviceProvider.BusinessName,
			AverageScore: averageScore,			
		}
		

		context.JSON(http.StatusOK, response)
	}

}

func (ServiceProviderController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
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
				Message: "Invalid page size parameter.",
			})
			return
		}

		kindOfService, err := strconv.Atoi(context.Query("kindOfService"))
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid kind of service parameter.",
			})
			return
		}
		

		price, err := strconv.ParseFloat(context.Query("maxPriceRate"), 64)
		if err != nil || price <= 0.0000 {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad Request",
				Message: "Invalid price rate parameter.",
			})
			return
		}
		

		cityId := context.Query("cityId")
		_, err = uuid.FromString(cityId)
			if err != nil {
				context.JSON(http.StatusBadRequest, response.ErrorResponse {
					Error: "Invalid ID",
					Message: "The city ID you provided has an invalid format",
				})
				return
			}

		filters := &FiltersKindOfServiceAndCity{
			KindOfService: kindOfService,
			CityID: cityId,
		}
		

		var price_rates []entity.PriceRate

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		r := db
		
		currentHour := time.Now()		
		hour := currentHour.Format("15:04")
		
		r = db.Scopes(utility.Paginate(page, pageElements)).
			Preload("ServiceProvider.User").Preload("ServiceProvider.User.Score").
			Preload("ServiceProvider").Where("price_rates.price <= ?", price).
			Where(filters).Where("? >= price_rates.starting_hour AND ? < price_rates.ending_hour", hour, hour).
			Where("price_rates.id IN (?)", db.Table("pricerate_workingdays").Select("price_rate_id").Where("pricerate_workingdays.price_rate_id = price_rates.id")).
			Find(&price_rates)

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not found",
				Message: "There are no service providers that match the filters.",
			})
			return
		}	
		
		result := []response.ServiceProvider{}
		
		for _, price_rate := range price_rates {
			result = append(result, mapper.CreateServiceProvidersAsResponse(price_rate))
		}

		var rowCount int64
		db.Model(&entity.PriceRate{}).Where("price_rates.price <= ?", price).
		Where(filters).Where("? >= price_rates.starting_hour AND ? < price_rates.ending_hour", hour, hour).
		Where("price_rates.id IN (?)", db.Table("pricerate_workingdays").Select("price_rate_id").
		Where("pricerate_workingdays.price_rate_id = price_rates.id")).Count(&rowCount)

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
			Data    []response.ServiceProvider		  	   `json:"data"`
		}{
			Links: struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			}{
				First: fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&cityId=%s&page=%d&pageSize=%d", price, kindOfService, cityId, 1, pageElements),
				Last:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&cityId=%s&page=%d&pageSize=%d", price, kindOfService, cityId, lastPage, pageElements),
				Prev:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&cityId=%s&page=%d&pageSize=%d", price, kindOfService, cityId, previousPage, pageElements),
				Next:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&cityId=%s&page=%d&pageSize=%d", price, kindOfService, cityId, nextPage, pageElements),
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

func (ServiceProviderController) GetStatistics() gin.HandlerFunc {
	return func(context *gin.Context) {
		startingDate := context.Query("startingDate")
		endingDate := context.Query("endingDate")
		providerID := context.Param("serviceProviderId")		

		_, err := uuid.FromString(providerID)

		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Invalid ID",
				Message: "The ID you provided has an invalid format",
			})
			return
		}

		_, err = time.Parse("2006-01-02", startingDate)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad request",
				Message: "The starting date you provided has a non-valid format.",
			})
			return
		}

		_, err = time.Parse("2006-01-02", endingDate)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse {
				Error: "Bad request",
				Message: "The ending date you provided has a non-valid format.",
			})
			return
		}		

		type RequestedServicesPerWeekdayqueryResult struct {
			RequestedServices int `json:"requestedServices"`
			Weekday           int `json:"weekday"`
		}

		type KindOfServicesQueryResult struct {
			RequestedServices int `json:"requestedServices"`
			KindOfService     int `json:"kindOfService"`
		}
		
		type Statistics struct {
			RequestedServicesPerWeekdayqueryResult []RequestedServicesPerWeekdayqueryResult `json:"requestedServicesPerWeekday"`
			KindOfServicesQueryResult []KindOfServicesQueryResult `json:"requestedServicesPerKindOfService"`
		}

		var kindOfServicesQueryResult []KindOfServicesQueryResult
		var requestedServicesPerWeekdayqueryResult []RequestedServicesPerWeekdayqueryResult

		db, err := database.New()
		if err != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}		

		r := db.Raw("SELECT COUNT(`id`) AS 'requested_services', "+
		"(WEEKDAY(DATE(`date`)) + 1) AS 'weekday' "+
		"FROM service_requests "+
		"WHERE service_provider_id = ? AND DATEDIFF(?, `date`) <= 0 "+
		"AND DATEDIFF(?, `date`) >= 0 "+
		"GROUP BY (WEEKDAY(DATE(`date`)) + 1)", providerID, startingDate, endingDate).Scan(&requestedServicesPerWeekdayqueryResult)

		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "We were unable to retrieve info about the dates you specified",
			})
			return
		}

		r = db.Raw("SELECT COUNT(`id`) AS 'requested_services', "+
		"`kind_of_service` AS 'kind_of_service' "+
		"FROM service_requests "+
		"WHERE service_provider_id = ? "+
		"AND DATEDIFF(?, `date`) <= 0 "+
		"AND DATEDIFF(?, `date`) >= 0 "+
		"GROUP BY(`kind_of_service`)", providerID, startingDate, endingDate).Scan(&kindOfServicesQueryResult)

		if r.Error != nil {
			context.JSON(http.StatusConflict, response.ErrorResponse {
				Error: "Internal Error",
				Message: "There was an unexpected error while processing your data. Please try again later",
			})
			return
		}

		if r.RowsAffected == 0 {
			context.JSON(http.StatusNotFound, response.ErrorResponse {
				Error: "Not Found",
				Message: "We were unable to retrieve info about the dates you specified",
			})
			return
		}

		statisticsReport := Statistics {
			RequestedServicesPerWeekdayqueryResult: requestedServicesPerWeekdayqueryResult,
			KindOfServicesQueryResult: kindOfServicesQueryResult,
		}

		context.JSON(http.StatusOK, statisticsReport)
	}
}
