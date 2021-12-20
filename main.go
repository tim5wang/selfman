package main

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/app/api/dig"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/common/middleware"
	"github.com/tim5wang/selfman/common/web"
	"github.com/tim5wang/selfman/dao/entity"
)

var (
	config *configservice.ConfigService

	//go:embed static/*
	embedFS embed.FS
)

func beforeStart() (err error) {
	err = dig.ApiContainer.Invoke(func(c *configservice.ConfigService) {
		config = c
	})
	if err != nil {
		return
	}
	err = dig.ApiContainer.Invoke(func(migrate *database.Migration) {
		err = migrate.Migrate(
			&entity.User{},
		)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	return
}

func main() {
	err := beforeStart()
	if err != nil {
		panic(err)
	}
	s := gin.New()
	s.Use(
		middleware.ConsoleLog,
		web.EmbedServer(embedFS, "static/", "/sf"),
		web.StaticServer(config.GetString("gin.static.path"), "/", "/v1/api"),
	)

	r := web.NewRouter("/v1", s)
	err = web.BindModuleFromContainer(r, dig.ApiContainer)
	if err != nil {
		panic(fmt.Errorf("bind modules error: %w", err))
	}

	err = s.Run(":8080")
	if err != nil {
		panic(fmt.Errorf("start service error: %w", err))
	}
}
