package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vahemuradyan2744/stream/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, stream")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())

	if err := router.Run("127.0.0.1:8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
