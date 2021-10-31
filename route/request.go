package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var requestController controller.RequestController

func init() {
	requestController = controller.RequestController{}
}

func AppendToRequestRoutes(r *gin.Engine) {
	sp := r.Group("/requests")
	sp.POST("", requestController.Store())

	sp.PATCH("/:serviceRequestId", requestController.StoreStatus())
}