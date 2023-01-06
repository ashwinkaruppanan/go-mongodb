package router

import (
	"github.com/ashwin/go-mongodb/controller"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.GET("/api/movie", controller.GetAllMovies)
	r.POST("/api/movie", controller.CreateMovie)
	r.PUT("/api/movie/:id", controller.MarkAsWatched)
	r.DELETE("/api/movie/:id", controller.DeleteOneMovie)
	r.DELETE("/api/movie", controller.DeleteAllMovies)

	r.Run()
}
