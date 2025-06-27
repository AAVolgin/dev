package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || !validateCredentials(username, password) {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func validateCredentials(user, pass string) bool {
	return user == "admin" && pass == "secret"
}
