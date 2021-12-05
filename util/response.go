package util

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/serror"
)

// Response represents an abstraction of http response.
// A Response must implements the `Render` method, which accepts a *gin.Context.
type Response interface {
	Render(*gin.Context)
}

// StaticFileResponse returns a static file resposne that writes the specified file into the body stream in a efficient way.
func StaticFileResponse(filepath string) Response {
	return &staticFileResponse{filepath}
}

type staticFileResponse struct {
	Path string
}

func (fs *staticFileResponse) Render(ctx *gin.Context) {
	ctx.File(fs.Path)
}

// FileResponse returns a file resposne that writes the specified reader into the body stream and updates the HTTP code.
func FileResponse(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) Response {
	return &fileResponse{
		code:          code,
		contentLength: contentLength,
		contentType:   contentType,
		reader:        reader,
		extraHeaders:  extraHeaders,
	}
}

type fileResponse struct {
	code          int
	contentLength int64
	contentType   string
	reader        io.Reader
	extraHeaders  map[string]string
}

func (fs *fileResponse) Render(ctx *gin.Context) {
	ctx.DataFromReader(fs.code, fs.contentLength, fs.contentType, fs.reader, fs.extraHeaders)
}

// JSONResponse returns a json response that serializes the given struct as JSON into the response body.
func JSONResponse(httpStatusCode int, err serror.Error, data interface{}) Response {
	if err == nil {
		err = serror.Success
	}
	return &jsonResponse{
		jsonResposneData: jsonResposneData{
			ErrCode: err.Code(),
			ErrMsg:  err.Msg(),
			Data:    data,
		},
		httpStatusCode: httpStatusCode,
	}
}

type jsonResponse struct {
	jsonResposneData
	httpStatusCode int
}

type jsonResposneData struct {
	ErrCode uint32      `json:"code"`
	ErrMsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (resp *jsonResponse) Render(ctx *gin.Context) {
	if resp == nil {
		return
	}

	ctx.JSON(resp.httpStatusCode, resp.jsonResposneData)
}

// RedirectResponse returns a redirect response that returns a HTTP redirect to the specific location.
func RedirectResponse(code int, location string) Response {
	return &redirectResponse{
		code:     code,
		location: location,
	}
}

type redirectResponse struct {
	code     int
	location string
}

func (resp *redirectResponse) Render(ctx *gin.Context) {
	ctx.Redirect(resp.code, resp.location)
}

// HTMLResponse returns a html response that renders the HTTP template specified by its file name.
func HTMLResponse(code int, name string, obj interface{}) Response {
	return &htmlResponse{
		code: code,
		name: name,
		obj:  obj,
	}
}

type htmlResponse struct {
	code int
	name string
	obj  interface{}
}

func (resp *htmlResponse) Render(ctx *gin.Context) {
	ctx.HTML(resp.code, resp.name, resp.obj)
}
