package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AdairHdz/OTW-Rest-API/database"
	"github.com/AdairHdz/OTW-Rest-API/middleware"
	"github.com/AdairHdz/OTW-Rest-API/route"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func init() {	
	f, _ := os.Create("./logs/requests.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)	
}

func main() {
	
	go func() {
        time.Sleep(30 * time.Second)
		_, err := database.New()

		if err != nil {
			log.Fatal(err)
		}
    }()
	
	os.Setenv("TZ", "America/Mexico_City")	
	limiter := tollbooth.NewLimiter(50, &limiter.ExpirableOptions{DefaultExpirationTTL: (1 * time.Minute)})
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	r := gin.Default()
	r.Use(middleware.Log())
	r.Use(tollbooth_gin.LimitHandler(limiter))
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