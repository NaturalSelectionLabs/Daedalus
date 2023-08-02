package command

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"net/http"
)

var (
	port int
	//go:embed dist/*
	dist     embed.FS
	tryFiles = []string{"index.html", "fallback.html"}
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "serve web ui",
	Run: func(cmd *cobra.Command, args []string) {
		serveUI(port)
	},
}

func init() {
	uiCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number for the web server")
}

func serveUI(port int) {
	r := gin.Default()
	r.StaticFS("/", getFileSystem())
	r.NoRoute(func(c *gin.Context) {
		data, err := dist.ReadFile("dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting web server at http://localhost%s\n", addr)

	log.Fatal(r.Run(addr))
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func accessLogsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		log.Printf("Method: %s, URL: %s, RemoteAddr: %s\n", r.Method, r.URL.String(), r.RemoteAddr)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
