package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var serviceProviderController controller.ServiceProviderController

func init() {
	serviceProviderController = controller.ServiceProviderController{}
}

func AppendToServiceProviderRoutes(r *gin.Engine) {
	sp := r.Group("/service-providers")
	sp.PUT("/:providerID/image", serviceProviderController.StoreImage())
	sp.GET("/:providerID", serviceProviderController.GetWithId())
	sp.GET("", serviceProviderController.Index())
}