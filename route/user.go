package route

import (
	"github.com/gin-gonic/gin"	
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var userController controller.UserController

func init() {
	userController = controller.UserController{}
}

func AppendUserRoutes(r *gin.Engine) {
	u := r.Group("/users")	
	u.POST("", userController.Store())
}