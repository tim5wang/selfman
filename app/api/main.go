package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/app/api/dig"
	"github.com/tim5wang/selfman/common/middleware"
	"github.com/tim5wang/selfman/util"
)

func main() {
	s := gin.New()
	//l := &controller.UserModule{}
	s.Use(middleware.Log, middleware.Log)
	//util.BindModule(s, l)
	err := util.BindModuleWithDig(s, dig.ApiContainer)
	if err != nil {
		panic(fmt.Errorf("bind modules error: %w", err))
	}
	err = s.Run(":8080")
	if err != nil {
		panic(fmt.Errorf("start service error: %w", err))
	}
}
