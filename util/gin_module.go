package util

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Module interface {
	Init(r gin.IRouter)
}

func BindModule(r gin.IRouter, handlers ...Module) {
	for _, h := range handlers {
		h.Init(r)
	}
}

var ModuleGroup = dig.Group("module")

type inParams struct {
	dig.In
	Modules []Module `group:"module"`
}

func BindModuleWithDig(r gin.IRouter, container *dig.Container) (err error) {
	err = container.Invoke(func(params inParams) {
		for _, module := range params.Modules {
			module.Init(r)
		}
	})
	return
}
