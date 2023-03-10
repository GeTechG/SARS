package sessions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetSessions().Exists(c.Request.Context(), "uid") {
			handler(c)
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error:": "You are not authorized!"})
	}
}
