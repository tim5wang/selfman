package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Module interface {
	Init(r Router)
}

func BindModule(r Router, handlers ...Module) {
	for _, h := range handlers {
		h.Init(r)
	}
}

var ModuleGroup = dig.Group("module")

type modules struct {
	dig.In
	Modules []Module `group:"module"`
}

func BindModuleFromContainer(r Router, container *dig.Container) (err error) {
	err = container.Invoke(func(param modules) {
		for _, module := range param.Modules {
			module.Init(r)
		}
	})
	return
}

type Validation interface {
	Validate() error
}

func BindJsonReq(ctx *gin.Context, reqs ...interface{}) (err error) {
	for _, req := range reqs {
		err = ctx.BindHeader(req)
		if err != nil {
			return
		}
		err = ctx.BindUri(req)
		if err != nil {
			return
		}
		err = ctx.BindQuery(req)
		if err != nil {
			return
		}
		if strings.Contains(ctx.GetHeader("Content-Type"), "json") {
			err = ctx.BindJSON(req)
			if err != nil {
				return
			}
		}

		v, ok := req.(Validation)
		if ok {
			err = v.Validate()
			if err != nil {
				return
			}
		}
	}
	return
}
