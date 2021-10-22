package route

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OTW-Rest-API/controller"
)

var reviewController controller.ReviewController

func init() {
	reviewController = controller.ReviewController{}
}

func AppendToReviewRoutes(r *gin.Engine) {
	sp := r.Group("/reviews")
	sp.GET("/:providerID", reviewController.GetWithId())
}