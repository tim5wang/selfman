package web

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func StaticServer(root, prefix string, excludes ...string) gin.HandlerFunc {
	if !strings.HasPrefix(root, "/") {
		root = path.Join("./", root)
	}
	fileSystem := http.Dir(root)
	return FileServer(fileSystem, "/", prefix, excludes...)
}
