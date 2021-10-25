package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var serviceRequesterController controller.ServiceRequesterController
var addressController controller.AddressController

func init() {
	serviceRequesterController = controller.ServiceRequesterController{}
	addressController = controller.AddressController{}
}

func AppendToServiceRequesterRoutes(r *gin.Engine) {
	sp := r.Group("/requesters")
	sp.POST("/:serviceRequesterId/addresses", addressController.Store())
}