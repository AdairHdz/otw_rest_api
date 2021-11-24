package route

import (
	"github.com/AdairHdz/OTW-Rest-API/controller"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/gin-gonic/gin"
)

var(	
	addressController controller.AddressController
) 

func init() {	
	addressController = controller.AddressController{}
	requestController = controller.RequestController{}
}

func AppendToServiceRequesterRoutes(r *gin.Engine) {
	sp := r.Group("/requesters")
	sp.Use(middleware.Authentication())
	sp.Use(middleware.ServiceRequesterAuthorization())
	sp.POST("/:serviceRequesterId/addresses", addressController.Store())
	sp.GET(":serviceRequesterId/addresses", addressController.Index())
	sp.GET("/:serviceRequesterId/requests/:serviceRequestId", requestController.GetRequestRequester())
	sp.GET("/:serviceRequesterId/requests", requestController.IndexRequestRequester())
}