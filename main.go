package main

import (	
	"log"
	"net/http"
	"os"
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
	
	os.Setenv("TZ", "America/Mexico_City")	
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(middleware.CORSMiddleware())
	route.AppendUserRoutes(r)
	route.AppendStateRoutes(r)
	route.AppendToServiceProviderRoutes(r)
	route.AppendToServiceRequesterRoutes(r)
	route.AppendToRequestRoutes(r)
	r.StaticFS("/images", http.Dir("./images"))
	r.StaticFS("/reviews", http.Dir("./public/reviews"))
	r.Run("0.0.0.0:8000")
}