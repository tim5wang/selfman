package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/web"
)

type DocModule struct {
	config *configservice.ConfigService
}

func NewDocModule(config *configservice.ConfigService) web.Module {
	return &DocModule{config: config}
}

func (m *DocModule) Init(r web.Router) {
	g := r.Group("api/doc")
	{
		g.GET("/config", m.GetConfig)
	}
}

func (m *DocModule) GetConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, m.config.GetString("gorm.path"))
}
