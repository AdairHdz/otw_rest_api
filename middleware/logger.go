package middleware

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("IP Address: %s | TimeStamp: [%s] | Requested Method: %s |" +
			" Path: %s | StatusCode: %d\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,					
		)
	})
}