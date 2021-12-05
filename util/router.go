package util

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/codegangsta/inject"
	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/serror"
)

// Handler can be any type of function.
type Handler interface{}

// Router is application router
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
	path     string
	injector inject.Injector
	rg       *gin.RouterGroup
}

// Group creates a new router. You should add all the routes that have common middlwares or the same path prefix.
func (r *router) Group(relativePath string, middlewares ...gin.HandlerFunc) Router {
	if !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}
	return &router{
		path:     r.path + relativePath,
		injector: r.injector,
		rg:       r.rg.Group(relativePath, middlewares...),
	}
}

// Handle registers a new request handle and middleware with the given path and method.
func (r *router) Handle(method, path string, handler Handler, middlewares ...gin.HandlerFunc) {
	//addHandlerInfo(method, strings.Replace(r.path+path, "//", "/", -1), handler, middlewares)
	var chain []gin.HandlerFunc
	for _, middleware := range middlewares {
		chain = append(chain, middleware)
	}
	chain = append(chain, r.wraphandler(handler))
	r.rg.Handle(method, path, chain...)
}

// GET is a shortcut for router.Handle("GET", path, handler, middlewares...).
func (r *router) GET(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("GET", path, handler, middlewares...)
}

// POST is a shortcut for router.Handle("POST", path, handler, middlewares...).
func (r *router) POST(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("POST", path, handler, middlewares...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handler, middlewares...).
func (r *router) DELETE(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("DELETE", path, handler, middlewares...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handler, middlewares...).
func (r *router) PATCH(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("PATCH", path, handler, middlewares...)
}

// PUT is a shortcut for router.Handle("PUT", path, handler, middlewares...).
func (r *router) PUT(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("PUT", path, handler, middlewares...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handler, middlewares...).
func (r *router) OPTIONS(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("OPTIONS", path, handler, middlewares...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handler, middlewares...).
func (r *router) HEAD(path string, handler Handler, middlewares ...gin.HandlerFunc) {
	r.Handle("HEAD", path, handler, middlewares...)
}

// wraphandler turns a normal Handler into a gin-handler compatible
func (r *router) wraphandler(f Handler) gin.HandlerFunc {
	return convertHandler(f, r.injector)
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

type ginwebRequest interface {
	Parse(*gin.Context) serror.Error
	Validate() serror.Error
	ResponseSchema() (interface{}, string) // only for api doc
}

func convertHandler(f Handler, parentInjector inject.Injector) gin.HandlerFunc {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("handler should be a function")
	}
	//switch t.NumOut() {
	//case 0:
	//case 1:
	//	//outTyp := t.Out(0)
	//	//if outTyp.Kind() != reflect.String && !outTyp.Implements(reflect.TypeOf((*Response)(nil)).Elem()) {
	//	//	panic("handler output parameter type should be `string` or `ginweb.Response`")
	//	//}
	//default:
	//	panic("handler output parameter count should be 0 or 1")
	//}

	numIn := t.NumIn()
	requestFields := []int{}
	for i := 0; i < numIn; i++ {
		if t.In(i).Implements(reflect.TypeOf((*ginwebRequest)(nil)).Elem()) {
			requestFields = append(requestFields, i)
		}
	}
	if len(requestFields) > 1 {
		panic("handler should only have one request")
	}
	return func(c *gin.Context) {
		injector := inject.New()
		if parentInjector != nil {
			injector.SetParent(parentInjector)
		}
		context := contextPool.Get().(*Context)
		context.ctx = c
		for _, field := range requestFields {
			if req := newReqInstance(t.In(field)); req != nil {
				rp := newRequestParser(req)
				if err := rp.parse(c); err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, &jsonResposneData{ErrCode: err.Code(), ErrMsg: err.Msg()})
					return
				}
				gr := req.(ginwebRequest)
				if err := gr.Parse(c); err != nil && err != serror.Success {
					c.AbortWithStatusJSON(http.StatusBadRequest, &jsonResposneData{ErrCode: err.Code(), ErrMsg: err.Msg()})
					return
				}
				if err := gr.Validate(); err != nil && err != serror.Success {
					c.AbortWithStatusJSON(http.StatusBadRequest, &jsonResposneData{ErrCode: err.Code(), ErrMsg: err.Msg()})
					return
				}
				injector.Map(req)
			}
		}
		injector.Map(context)
		injector.Map(c)
		ret, err := injector.Invoke(f)
		if err != nil {
			panic(err)
		}

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

// ConvertHandler converts a ginweb handler to gin handler.
func ConvertHandler(f Handler) gin.HandlerFunc {
	return convertHandler(f, nil)
}
