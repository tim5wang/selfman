package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/util"
	"github.com/tim5wang/selfman/common/web"
)

type ImageModule struct {
	conf *configservice.ConfigService
}

func NewImageModule(conf *configservice.ConfigService) web.Module {
	return &ImageModule{
		conf: conf,
	}
}

func (m *ImageModule) Init(r web.Router) {
	g := r.Group("api/file")
	{
		g.POST("/upload", m.Upload)
		g.POST("/clone", m.Clone)
	}
}

type UploadImageReq struct {
	FileList []*multipart.FileHeader `form:"file[]"`
}
type UploadImageRsp struct {
	ErrorFiles []string          `json:"errFiles"`
	SuccessMap map[string]string `json:"succMap"`
}

func (m *ImageModule) Upload(ctx *gin.Context, req *UploadImageReq) {
	util.Print(req)
	rsp := &UploadImageRsp{
		ErrorFiles: make([]string, 0),
		SuccessMap: make(map[string]string),
	}
	for _, file := range req.FileList {
		dst, err := os.Create(path.Join(m.conf.GetString("upload.image.dir"), file.Filename))
		if err != nil {
			rsp.ErrorFiles = append(rsp.ErrorFiles, file.Filename)
			continue
		}
		f, err := file.Open()
		if err != nil {
			rsp.ErrorFiles = append(rsp.ErrorFiles, file.Filename)
			continue
		}
		_, err = io.Copy(dst, f)
		if err != nil {
			rsp.ErrorFiles = append(rsp.ErrorFiles, file.Filename)
			continue
		}
		rsp.SuccessMap[file.Filename] = path.Join(m.conf.GetString("upload.image.path"), file.Filename)
	}
	web.Success(ctx, rsp)
}

type CloneImageReq struct {
	URL string `json:"url"`
}

type CloneImageRsp struct {
	OriginalURL string `json:"originalURL"`
	NewURL      string `json:"url"`
}

func (m *ImageModule) Clone(ctx *gin.Context, req *CloneImageReq) {
	rsp, err := http.Get(req.URL)
	if err != nil {
		web.Success(ctx, CloneImageRsp{OriginalURL: req.URL, NewURL: req.URL})
		return
	}
	defer rsp.Body.Close()
	ps := strings.Split(req.URL, "/")
	var name string
	if len(ps) > 1 {
		name = ps[len(ps)-1]
	}

	out, err := os.Create(path.Join(m.conf.GetString("upload.image.dir"), name))
	defer out.Close()
	pix, err := ioutil.ReadAll(rsp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		web.Success(ctx, CloneImageRsp{OriginalURL: req.URL, NewURL: req.URL})
		return
	}
	web.Success(ctx, CloneImageRsp{OriginalURL: req.URL, NewURL: path.Join(m.conf.GetString("upload.image.path"), name)})
	return
}
