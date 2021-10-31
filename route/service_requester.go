package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var(
	serviceRequesterController controller.ServiceRequesterController
	addressController controller.AddressController
) 

func init() {
	serviceRequesterController = controller.ServiceRequesterController{}
	addressController = controller.AddressController{}
	requestController = controller.RequestController{}
}

func AppendToServiceRequesterRoutes(r *gin.Engine) {
	sp := r.Group("/requesters")
	sp.POST("/:serviceRequesterId/addresses", addressController.Store())
	sp.GET(":serviceRequesterId/addresses", addressController.Index())

	sp.GET("/:serviceRequesterId/requests/:serviceRequestId", requestController.GetRequestRequester())
	sp.GET("/:serviceRequesterId/requests", requestController.IndexRequestRequester())
}