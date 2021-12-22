package web

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

type fileRender struct {
	data        []byte
	contentType string
}

func (r *fileRender) Render(w http.ResponseWriter) error {
	_, err := w.Write(r.data)
	return err
}

func (r *fileRender) WriteContentType(w http.ResponseWriter) {
	//w.Header().Set("content-type", r.contentType)
}

func loadFile(prePath, fileName string, fSys fs.FS) (file fs.File, err error) {
	prePath = strings.TrimRight(prePath, "/")
	file, err = fSys.Open(path.Join(prePath, fileName))
	if err == nil && isFile(file) {
		return
	}
	file, err = fSys.Open(path.Join(prePath, fileName, "index.html"))
	if err == nil && isFile(file) {
		return
	}
	file, err = fSys.Open(path.Join(prePath, fileName, "README.md"))
	if err == nil && isFile(file) {
		return
	}
	file, err = fSys.Open(path.Join(prePath, fileName, "INDEX.md"))
	if err == nil && isFile(file) {
		return
	}
	file, err = fSys.Open(path.Join(prePath, "index.html"))
	if err == nil && isFile(file) {
		return
	}
	file, err = fSys.Open(path.Join(prePath, "404.html"))
	return
}

func isFile(file fs.File) bool {
	info, _ := file.Stat()
	return !info.IsDir()
}

func FileServer(fileSystem fs.FS, prePath, prefix string, excludes ...string) gin.HandlerFunc {
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

		file, err := loadFile(prePath, filePath, fileSystem)
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
		//ctx.Writer.Write(data)
		//ctx.Writer.Flush()
		ctx.Render(http.StatusOK, &fileRender{data: data, contentType: "text/html"})
		ctx.Abort()
	}
}
