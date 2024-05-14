package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware configures the CORS policy for the router
func CORSMiddleware() gin.HandlerFunc {
        return func(c *gin.Context) {
                c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
                c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
                c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
                c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

                if c.Request.Method == "OPTIONS" {
                        c.AbortWithStatus(http.StatusNoContent)
                        return
                }

                c.Next()
        }
}