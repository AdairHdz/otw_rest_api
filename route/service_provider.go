package route

import (
	"github.com/AdairHdz/OTW-Rest-API/controller"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/gin-gonic/gin"
)

var (
	serviceProviderController controller.ServiceProviderController
	reviewController controller.ReviewController
	priceRateController controller.PriceRateController
) 

func init() {
	serviceProviderController = controller.ServiceProviderController{}
	reviewController = controller.ReviewController{}
	priceRateController = controller.PriceRateController{}
	requestController = controller.RequestController{}
}

func AppendToServiceProviderRoutes(r *gin.Engine) {
	sp := r.Group("/providers")
	sp.PUT("/:serviceProviderId/image", serviceProviderController.StoreImage())
	
	sp.Use(middleware.Authentication())
	sp.GET("/:serviceProviderId", serviceProviderController.GetWithId())	
	sp.GET("/:serviceProviderId/reviews", reviewController.GetWithId())	
	sp.GET("/:serviceProviderId/priceRates", priceRateController.FindAll())
	sp.POST("/:serviceProviderId/reviews", reviewController.Store())
	sp.POST("/:serviceProviderId/reviews/:reviewId/evidence", reviewController.UploadEvidence())
	sp.GET("/:serviceProviderId/priceRates/:cityId", priceRateController.FindActivePriceRate())
	sp.GET("", serviceProviderController.Index())

	sp.Use(middleware.ServiceProviderAuthorization())
	sp.GET("/:serviceProviderId/statistics", serviceProviderController.GetStatistics())
	sp.POST("/:serviceProviderId/priceRates", priceRateController.Store())
	sp.DELETE("/:serviceProviderId/priceRates/:priceRateId", priceRateController.Delete())			
	sp.GET("/:serviceProviderId/requests/:serviceRequestId", requestController.GetRequestProvider())
	sp.GET("/:serviceProviderId/requests", requestController.IndexRequestProvider())
}