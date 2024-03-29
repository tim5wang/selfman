package main

import (
	"embed"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/app/api/dig"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/common/middleware"
	"github.com/tim5wang/selfman/common/util"
	"github.com/tim5wang/selfman/common/web"
	"github.com/tim5wang/selfman/dao/entity"
)

var (
	config *configservice.ConfigService
	//go:embed static app/api/conf
	embedFS embed.FS
)

func beforeStart() (err error) {
	err = dig.ApiContainer.Invoke(func(c *configservice.ConfigService) {
		config = c
		err = config.FromEmbed(embedFS)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		return
	}
	err = dig.ApiContainer.Invoke(func(migrate *database.Migration) {
		err = migrate.Migrate(
			&entity.User{},
			&entity.Doc{},
			&entity.File{},
			&entity.ID{},
		)
		if err != nil {
			panic(err)
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
		web.EmbedServer(embedFS, "static/", "/s/", "/s/img"),
		web.StaticServer(config.GetString("gin.static.path"), "/", "/v1/api", "/s/"),
		web.StaticServer(config.GetString("upload.image.dir"), config.GetString("upload.image.path")),
		//middleware.APIDoc,
	)

	r := web.NewRouter("/v1", s)
	err = web.BindModuleFromContainer(r, dig.ApiContainer)
	if err != nil {
		panic(fmt.Errorf("bind modules error: %w", err))
	}
	timer := time.NewTimer(1 * time.Second)
	go func() {
		<-timer.C
		_ = util.Open("http://localhost" + config.GetString("port") + config.GetString("home"))
	}()
	err = s.Run(config.GetString("port"))
	if err != nil {
		panic(fmt.Errorf("start service error: %w", err))
	}
}
