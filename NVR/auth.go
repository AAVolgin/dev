package main

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Простейшая заглушка: всегда разрешаем
		c.Next()
	}
}
