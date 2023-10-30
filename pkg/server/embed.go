package server

//import (
//	"embed"
//	"fmt"
//	"github.com/gin-contrib/static"
//	"io/fs"
//	"net/http"
//	"strings"
//)
//
//var (
//	//go:embed dist/*
//	dist embed.FS
//)
//
//type staticFileSystem struct {
//	http.FileSystem
//}
//
//var _ static.ServeFileSystem = (*staticFileSystem)(nil)
//
//func newStaticFileSystem() *staticFileSystem {
//	sub, err := fs.Sub(dist, "dist")
//	if err != nil {
//		panic(err)
//	}
//	return &staticFileSystem{
//		FileSystem: http.FS(sub),
//	}
//}
//
//func (s *staticFileSystem) Exists(prefix string, path string) bool {
//	buildpath := fmt.Sprintf("dist%s", path)
//
//	// support for folders
//	if strings.HasSuffix(path, "/") {
//		_, err := dist.ReadDir(strings.TrimSuffix(buildpath, "/"))
//		return err == nil
//	}
//
//	// support for files
//	f, err := dist.Open(buildpath)
//	if f != nil {
//		_ = f.Close()
//	}
//	return err == nil
//}
//
//type fallbackFileSystem struct {
//	staticFileSystem *staticFileSystem
//}
//
//var _ static.ServeFileSystem = (*fallbackFileSystem)(nil)
//var _ http.FileSystem = (*fallbackFileSystem)(nil)
//
//func newFallbackFileSystem(staticFileSystem *staticFileSystem) *fallbackFileSystem {
//	return &fallbackFileSystem{
//		staticFileSystem: staticFileSystem,
//	}
//}
//
//func (f *fallbackFileSystem) Open(path string) (http.File, error) {
//	return f.staticFileSystem.Open("/index.html")
//}
//
//func (f *fallbackFileSystem) Exists(prefix string, path string) bool {
//	return strings.Index(path, "/api") == -1
//}
