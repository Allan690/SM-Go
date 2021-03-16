package middleware

import (
	user "StoreManager/model/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// middleware to validate token
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		_, err := user.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			log.Print(err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,gin.H{
					"status": "error", "message": err.Error(),
				})
		}
	}
}
