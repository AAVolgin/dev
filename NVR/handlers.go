package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCameras(c *gin.Context) {
	var cameras []Camera
	DB.Find(&cameras)
	c.JSON(http.StatusOK, cameras)
}

func AddCamera(c *gin.Context) {
	var camera Camera
	if err := c.ShouldBindJSON(&camera); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&camera)
	c.JSON(http.StatusOK, camera)
}

func DeleteCamera(c *gin.Context) {
	id := c.Param("id")
	DB.Delete(&Camera{}, id)
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
