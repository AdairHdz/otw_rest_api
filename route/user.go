package route

import (
	"github.com/AdairHdz/OTW-Rest-API/controller"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/gin-gonic/gin"
)

var userController controller.UserController

func init() {
	userController = controller.UserController{}
}

func AppendUserRoutes(r *gin.Engine) {
	u := r.Group("/users")	
	u.POST("", userController.Store())
	u.POST("/login", userController.Login())
	u.POST("/:userId/token/refresh", userController.RefreshToken())
	u.Use(middleware.Authentication())
	u.PUT("/:userId/verify", userController.SendEmail())
	u.PATCH("/:userId/verify", userController.Verify())
}