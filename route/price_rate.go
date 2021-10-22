package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var priceRateController controller.PriceRateController

func init() {
	priceRateController = controller.PriceRateController{}
}

func AppendToPriceRateRoutes(r *gin.Engine) {
	sp := r.Group("/price-rates")
	sp.GET("/:providerID", priceRateController.FindAll())
	sp.POST("/:providerID", priceRateController.Store())
}