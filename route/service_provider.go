package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var(
	serviceProviderController controller.ServiceProviderController
	reviewController controller.ReviewController
	priceRateController controller.PriceRateController
) 

func init() {
	serviceProviderController = controller.ServiceProviderController{}
	reviewController = controller.ReviewController{}
	priceRateController = controller.PriceRateController{}
}

func AppendToServiceProviderRoutes(r *gin.Engine) {
	sp := r.Group("/providers")
	sp.PUT("/:serviceProviderId/image", serviceProviderController.StoreImage())
	
	sp.GET("/:serviceProviderId", serviceProviderController.GetWithId())
	sp.GET("", serviceProviderController.Index())

	sp.GET("/:serviceProviderId/reviews", reviewController.GetWithId())

	sp.GET("/:serviceProviderId/priceRates", priceRateController.FindAll())
	sp.POST("/:serviceProviderId/priceRates", priceRateController.Store())
	sp.GET("/:serviceProviderId/priceRates/:cityId", priceRateController.FindActivePriceRate())
}