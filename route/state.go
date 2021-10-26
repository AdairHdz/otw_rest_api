package route

import (
	"github.com/gin-gonic/gin"	
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var stateController controller.StateController

func init() {
	stateController = controller.StateController{}
}

func AppendStateRoutes(r *gin.Engine) {	
	u := r.Group("/states")
	u.GET("", stateController.Index())
	u.GET("/:stateId/cities", stateController.IndexCities())
}