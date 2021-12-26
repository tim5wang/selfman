package dig

import (
	"fmt"

	"github.com/tim5wang/selfman/app/api/controller"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/common/env"
	"github.com/tim5wang/selfman/common/web"
	"github.com/tim5wang/selfman/dao/filedao"
	"github.com/tim5wang/selfman/dao/userdao"
	"github.com/tim5wang/selfman/service/fileservice"
	"github.com/tim5wang/selfman/service/userservice"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

var (
	appName  = "selfman-api"
	confPath = "./app/api/conf/"
)
var (
	ApiContainer *dig.Container
)

func init() {
	var (
		err error
	)
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("container load error: %v", err))
		}
	}()
	ApiContainer = dig.New()
	{ // controller
		err = ApiContainer.Provide(controller.NewUserModule, web.ModuleGroup)
		err = ApiContainer.Provide(controller.NewDocModule, web.ModuleGroup)
		err = ApiContainer.Provide(controller.NewImageModule, web.ModuleGroup)

	}
	{ // common
		err = ApiContainer.Provide(NewApiConfig)
		err = ApiContainer.Provide(NewApiDB)
		err = ApiContainer.Provide(database.NewMigration)
	}
	{ // service
		err = ApiContainer.Provide(userservice.NewUserService)
		err = ApiContainer.Provide(fileservice.NewFileService)
	}
	{ // dao
		err = ApiContainer.Provide(userdao.NewUserDao)
		err = ApiContainer.Provide(filedao.NewFileDao)
	}
}

func NewApiConfig() *configservice.ConfigService {
	options := &configservice.Options{
		Engines: make([]configservice.KVEngine, 0),
	}
	if env.Env() == env.Env_dev {
		options.Engines = append(options.Engines,
			configservice.NewYamlConfig(confPath, "config-dev", "yaml"))
	} else {
		options.Engines = append(options.Engines,
			configservice.NewYamlConfig(confPath, "config-live", "yaml"))
	}
	options.Engines = append(options.Engines,
		configservice.NewYamlConfig(confPath, "config", "yaml"))
	return configservice.NewConfigService(options)
}

func NewApiDB(config *configservice.ConfigService) *gorm.DB {
	file := config.GetString("gorm.sqlite.path")
	if file == "" {
		panic("sqlite file not config")
	}
	return database.CreateSQLite(file)
}
