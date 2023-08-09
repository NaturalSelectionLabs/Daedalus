package server

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Run(port int) {
	r := gin.Default()
	embeddedBuildFolder := newStaticFileSystem()
	fallbackFileSystem := newFallbackFileSystem(embeddedBuildFolder)
	r.Use(static.Serve("/", embeddedBuildFolder))
	r.Use(static.Serve("/", fallbackFileSystem))

	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/gh-action")
	})

	api := r.Group("/api")
	{
		api.GET("/aaa", func(context *gin.Context) {
			context.String(200, "ok")
		})
	}
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting web server at http://localhost%s\n", addr)

	log.Fatal(r.Run(addr))
}
