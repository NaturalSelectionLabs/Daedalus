package command

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
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
	//http.HandleFunc("/", handleRequest)
	http.Handle("/", http.FileServer(getFileSystem()))
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting web server at http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting the web server:", err)
	}
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the requested file path
	requestedFile := r.URL.Path[1:]

	// Try to find the requested file in the dist folder
	if data, err := tryFile(requestedFile); err == nil {
		w.Header().Set("Content-Type", getContentType(requestedFile))
		w.Write(data)
		return
	}

	// If the requested file doesn't exist, return a 404 page or handle the error
	http.NotFound(w, r)
}

func tryFile(filePath string) ([]byte, error) {
	for _, file := range tryFiles {
		fullPath := "dist/" + file
		data, err := fs.ReadFile(dist, fullPath)
		if err == nil {
			return data, nil
		}
	}
	return nil, fs.ErrNotExist
}

func getContentType(filePath string) string {
	// Implement a mapping for Content-Type based on the file extension, if needed.
	// For simplicity, this example assumes the Content-Type is always "text/html".
	return "text/html"
}
