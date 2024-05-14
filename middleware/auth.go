package middleware

import (
	"github.com/AshishKothariii/gotalkapi/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for the presence of a JWT and validates it
func AuthMiddleware() gin.HandlerFunc {
        return func(c *gin.Context) {
                tokenString := c.GetHeader("Authorization")

                if tokenString == "" {
                        c.JSON(401, gin.H{"error": "Unauthorized"})
                        c.Abort()
                        return
                }

                userID, err := utils.ParseJWT(tokenString)
                if err != nil {
                        c.JSON(401, gin.H{"error": "Unauthorized: Invalid token"})
                        c.Abort()
                        return
                }

                // Set userID to Gin Context
                c.Set("userID", userID)

                c.Next()
        }
}