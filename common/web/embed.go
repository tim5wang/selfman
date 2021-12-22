package web

import (
	"io/fs"

	"github.com/gin-gonic/gin"
)

func EmbedServer(root fs.FS, prePath, prefix string, excludes ...string) gin.HandlerFunc {
	return FileServer(root, prePath, prefix, excludes...)
}
