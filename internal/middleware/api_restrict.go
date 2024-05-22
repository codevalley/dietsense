package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RestrictToIPs restricts access to endpoints based on allowed IP addresses.
func RestrictToIPs(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		for _, allowedIP := range allowedIPs {
			if clientIP == allowedIP {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Access restricted"})
		c.Abort()
	}
}
