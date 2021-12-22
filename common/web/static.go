package web

import (
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func StaticServer(root, prefix string, excludes ...string) gin.HandlerFunc {
	if !strings.HasPrefix(root, "/") {
		root = path.Join("./", root)
	}
	fileSystem := os.DirFS(root)
	return FileServer(fileSystem, "/", prefix, excludes...)
}
