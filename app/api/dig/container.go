package dig

import (
	"fmt"

	"github.com/tim5wang/selfman/app/api/controller"
	"github.com/tim5wang/selfman/common/web"
	"go.uber.org/dig"
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
}
