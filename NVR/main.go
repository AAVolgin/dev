package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	InitLogger()             // должно быть в logger.go
	InitDB()                 // должно быть в models.go
	StartRecorderManager()   // должно быть в recorder.go

	r := gin.Default()
	r.Use(AuthMiddleware())  // должно быть в auth.go

	api := r.Group("/api")
	{
		api.GET("/cameras", ListCameras)   // должно быть в handlers.go
		api.POST("/cameras", AddCamera)
		api.DELETE("/cameras/:id", DeleteCamera)
	}

	r.Static("/archive", "./storage")

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}