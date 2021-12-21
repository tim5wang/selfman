package web

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/util"
)

func EmbedServer(root fs.FS, prePath, prefix string, excludes ...string) gin.HandlerFunc {
	fileSystem := http.FS(root)
	return FileServer(fileSystem, prePath, prefix, excludes...)
}

func FileServer(fileSystem http.FileSystem, prePath, prefix string, excludes ...string) gin.HandlerFunc {
	notFoundPage, _ := fileSystem.Open(path.Join(prePath, "404.html"))
	return func(ctx *gin.Context) {
		// 排除
		if ctx.Request.Method != http.MethodGet {
			ctx.Next()
			return
		}
		urlPath := ctx.Request.URL.Path
		// 排除url无此前缀的
		if !strings.HasPrefix(urlPath, prefix) {
			ctx.Next()
			return
		}
		// 白名单路径
		for _, exclude := range excludes {
			if exclude == "" {
				continue
			}
			exclude = strings.TrimRight(exclude, "*")
			if strings.HasPrefix(urlPath, exclude) {
				ctx.Next()
				return
			}
		}
		filePath := strings.TrimLeft(urlPath, prefix)
		filePath = path.Join(prePath, filePath)
		if strings.TrimRight(filePath, "/") == strings.TrimRight(prePath, "/") {
			filePath = path.Join(filePath, "index.html")
		}
		util.Print(filePath)
		//ctx.FileFromFS(filePath, fileSystem) // 自带的有重定向泥潭
		file, err := fileSystem.Open(filePath)
		if err != nil {
			file = notFoundPage
			if file == nil {
				return
			}
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return
		}
		ctx.Writer.Write(data)
		ctx.Writer.Flush()
	}
}
