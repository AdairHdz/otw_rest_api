package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/AdairHdz/OTW-Rest-API/route"
	"github.com/gin-gonic/gin"
)

func main() {
	
	go func() {
        time.Sleep(30 * time.Second)
		_, err := database.New()

		if err != nil {
			log.Fatal(err)
		}
    }()
	
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(middleware.CORSMiddleware())
	route.AppendUserRoutes(r)
	route.AppendStateRoutes(r)
	route.AppendToServiceProviderRoutes(r)
	route.AppendToReviewRoutes(r)
	route.AppendToPriceRateRoutes(r)
	r.StaticFS("/images", http.Dir("./images"))
	r.Run("0.0.0.0:8000")
}