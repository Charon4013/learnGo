package main

import (
	"github.com/gin-gonic/gin"
	"go-example/controllers"
	"go-example/middlewares"
	"log"
)

func main() {
	server := gin.Default()

	server.Use(middlewares.MyAuth())
	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong")
	})

	server.Static("/resources", "./resources")
	server.StaticFile("/qianfeng", "./resources/qianfeng.jpg")

	pictureController := controllers.NewPictureController()

	pictureGroup := server.Group("/pictures")

	pictureGroup.Use(middlewares.MyLogger())

	// GET /pictures
	pictureGroup.GET("", pictureController.GetAll)
	// PUT /pictures/123
	pictureGroup.PUT("/:id", pictureController.Update)
	// POST /pictures
	pictureGroup.POST("", pictureController.Create)
	// DELETE /pictures/123
	pictureGroup.DELETE("/:id", pictureController.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
