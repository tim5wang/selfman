package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/app/api/dig"
	"github.com/tim5wang/selfman/common/middleware"
	"github.com/tim5wang/selfman/common/web"
)

func main() {
	s := gin.New()
	s.Use(middleware.Log, middleware.Log)
	r := web.NewRouter("/v1", s)
	err := web.BindModuleFromContainer(r, dig.ApiContainer)
	if err != nil {
		panic(fmt.Errorf("bind modules error: %w", err))
	}
	err = s.Run(":8080")
	if err != nil {
		panic(fmt.Errorf("start service error: %w", err))
	}
}
