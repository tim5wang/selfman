package util

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/serror"
)

// Request is ginweb's request structure that gets embedded in user defined request.
type Request struct{}

// Parse parses request from gin context.
func (r *Request) Parse(c *gin.Context) serror.Error {
	return serror.Success
}

// Validate checks the validation of the request.
func (r *Request) Validate() serror.Error {
	return serror.Success
}

type requestParser struct {
	req interface{}
	err error
}

func newRequestParser(req interface{}) *requestParser {
	return &requestParser{
		req: req,
	}
}

func (rp *requestParser) parse(c *gin.Context) serror.Error {
	err := BindJsonReq(c, rp.req)
	if err != nil {
		return serror.ErrorParseRequest.SetMsg(rp.err.Error())
	}
	//rp.err = c.ShouldBind(rp.req)
	//rp.bindContext(c, rp.req)
	//if rp.err != nil {
	//	return serror.ErrorParseRequest.SetMsg(rp.err.Error())
	//}
	return nil
}

func (rp *requestParser) bindContext(c *gin.Context, s interface{}) {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	// debugPrint(s, typ.String(), typ.Kind(), val)
	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			typeField := typ.Field(i)
			structField := val.Field(i)
			// debugPrint("--", typeField.Name, structField.Kind(), structField.CanSet())
			if !structField.CanSet() {
				continue
			}
			structFieldKind := structField.Kind()
			switch structFieldKind {
			case reflect.Ptr:
				v := reflect.ValueOf(newReqInstance(structField.Type()))
				if v.Elem().Kind() == reflect.Struct {
					structField.Set(v)
					//rp.bindContext(c, structField.Interface())
					_ = BindJsonReq(c, structField.Interface())
					continue
				}
			case reflect.Struct:
				//rp.bindContext(c, structField.Addr().Interface())
				_ = BindJsonReq(c, structField.Addr().Interface())
				continue
			}
			_ = typeField

			//for _, binder := range ctxbBinders {
			//	err := binder.Bind(c, &typeField, &structField)
			//	if err != nil {
			//		rp.err = err
			//		return
			//	}
			//}
		}
	}
}
