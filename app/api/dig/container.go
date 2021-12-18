package dig

import (
	"fmt"

	"github.com/tim5wang/selfman/app/api/controller"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/env"
	"github.com/tim5wang/selfman/common/web"
	"go.uber.org/dig"
)

var (
	appName  = "selfman-api"
	confPath = "app/api/conf/"
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
	}
	err = ApiContainer.Provide(NewApiConfig)
}
func NewApiConfig() *configservice.ConfigService {
	options := &configservice.Options{
		Engines: make([]configservice.KVEngine, 0),
	}
	if env.Env() == env.Env_dev {
		options.Engines = append(options.Engines,
			configservice.NewYamlConfig(confPath+"config-dev.yml"))
	} else {
		options.Engines = append(options.Engines,
			configservice.NewYamlConfig(confPath+"config-live.yml"))
	}
	options.Engines = append(options.Engines,
		configservice.NewYamlConfig(confPath+"config.yml"))
	return configservice.NewConfigService(options)
}
