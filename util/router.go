package util

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/serror"
)

type Handler interface{}

type Router interface {
	Group(relativePath string, middlewares ...gin.HandlerFunc) Router
	Handle(method, path string, handler Handler, middlewares ...gin.HandlerFunc)
	GET(path string, handler Handler, middlewares ...gin.HandlerFunc)
	POST(path string, handler Handler, middlewares ...gin.HandlerFunc)
	DELETE(path string, handler Handler, middlewares ...gin.HandlerFunc)
	PATCH(path string, handler Handler, middlewares ...gin.HandlerFunc)
	PUT(path string, handler Handler, middlewares ...gin.HandlerFunc)
	OPTIONS(path string, handler Handler, middlewares ...gin.HandlerFunc)
	HEAD(path string, handler Handler, middlewares ...gin.HandlerFunc)
}

type router struct {
	path string
	rg   *gin.RouterGroup
}

func (r *router) Group(relativePath string, middlewares ...gin.HandlerFunc) Router {
	if !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}
	return &router{
		path: r.path + relativePath,
		rg:   r.rg.Group(relativePath, middlewares...),
	}
}

func (r *router) Handle(method, path string, handler Handler, middlewares ...gin.HandlerFunc) {
	chain := make([]gin.HandlerFunc, 0)
	chain = append(chain, middlewares...)
	chain = append(chain, r.wraphandler(handler))
	r.rg.Handle(method, path, chain...)
}

func (r *router) GET(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("GET", path, handler, middlewares...)
}

func (r *router) POST(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("POST", path, handler, middlewares...)
}

func (r *router) DELETE(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("DELETE", path, handler, middlewares...)
}

func (r *router) PATCH(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("PATCH", path, handler, middlewares...)
}

func (r *router) PUT(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("PUT", path, handler, middlewares...)
}

func (r *router) OPTIONS(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("OPTIONS", path, handler, middlewares...)
}

func (r *router) HEAD(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("HEAD", path, handler, middlewares...)
}

func (r *router) wraphandler(f Handler) gin.HandlerFunc {
	return convertHandler(f)
}

func newReqInstance(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr:
		return newReqInstance(t.Elem())
	case reflect.Interface:
		return nil
	default:
		return reflect.New(t).Interface()
	}
}

type sRequest interface {
	Parse(*gin.Context) serror.Error
	Validate() serror.Error
}

func convertHandler(f Handler) gin.HandlerFunc {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("handler should be a function")
	}
	numIn := t.NumIn()
	requestFields := make([]int, 0)
	for i := 0; i < numIn; i++ {
		if t.In(i).Implements(reflect.TypeOf((*sRequest)(nil)).Elem()) {
			requestFields = append(requestFields, i)
		}
	}
	if len(requestFields) > 1 {
		panic("handler should only have one request")
	}
	return func(c *gin.Context) {
		realParams := make([]reflect.Value, len(requestFields))
		for i, field := range requestFields {
			if req := newReqInstance(t.In(field)); req != nil {
				err := BindJsonReq(c, req)
				if err != nil {
					panic(err.Error)
				}
				realParams[i] = reflect.ValueOf(req)
			}
		}
		ret := reflect.ValueOf(f).Call(realParams)
		if len(ret) > 0 {
			i := ret[0].Interface()
			switch i.(type) {
			case Response:
				if i != nil {
					i.(Response).Render(c)
				}
			case string:
				c.String(http.StatusOK, i.(string))
			}
		}
	}
}

func ConvertHandler(f Handler) gin.HandlerFunc {
	return convertHandler(f)
}
