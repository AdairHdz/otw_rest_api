package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
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
	
	sp.GET("/:serviceProviderId", serviceProviderController.GetWithId())
	sp.GET("", serviceProviderController.Index())

	sp.GET("/:serviceProviderId/reviews", reviewController.GetWithId())

	sp.GET("/:serviceProviderId/priceRates", priceRateController.FindAll())
	sp.POST("/:serviceProviderId/priceRates", priceRateController.Store())
	sp.POST("/:serviceProviderId/reviews", reviewController.Store())
	sp.POST("/:serviceProviderId/reviews/:reviewId/evidence", reviewController.UploadEvidence())
	sp.GET("/:serviceProviderId/priceRates/:cityId", priceRateController.FindActivePriceRate())

	sp.GET("/:serviceProviderId/requests/:serviceRequestId", requestController.GetRequestProvider())
	sp.GET("/:serviceProviderId/requests", requestController.IndexRequestProvider())
}