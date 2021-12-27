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

type fileRender struct {
	file     fs.File
	data     []byte
	fileName string
}

func (r *fileRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	data := r.getData()
	_, err := w.Write(data)
	return err
}

func (r *fileRender) getData() []byte {
	if r.data != nil {
		return r.data
	}
	data, err := ioutil.ReadAll(r.file)
	if err != nil {
		return []byte{}
	}
	r.data = data
	return r.data
}

func (r *fileRender) getHeader() string {
	info, _ := r.file.Stat()
	suffix := path.Ext(info.Name())
	header := util.GetFileHeader(suffix, r.getData())
	if header == "" {
		header = "*/*; charset=utf-8"
	}
	return header
}

func (r *fileRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.getHeader())
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
			ctx.Abort()
			return
		}
		ctx.Render(http.StatusOK, &fileRender{file: file})
		ctx.Writer.Flush()
		ctx.Abort()
	}
}
