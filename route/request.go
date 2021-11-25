package route

import (
	"github.com/AdairHdz/OTW-Rest-API/controller"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/gin-gonic/gin"
)

var requestController controller.RequestController

func init() {
	requestController = controller.RequestController{}
}

func AppendToRequestRoutes(r *gin.Engine) {
	sp := r.Group("/requests")
	sp.Use(middleware.Authentication())
	sp.POST("", requestController.Store())
	sp.PATCH("/:serviceRequestId", requestController.StoreStatus())
}