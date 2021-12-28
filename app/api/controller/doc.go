package controller

import (
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
		g.GET("/:doc_id", m.GetDoc)
		g.POST("/save", m.SaveDoc)
	}
}

type GetDocReq struct {
	DocID string `json:"doc_id" uri:"doc_id"`
}
type DocReqRsp struct {
	DocID string `json:"doc_id"`
}

func (m *DocModule) GetDoc(ctx *gin.Context, req *GetUserReq) {

}

func (m *DocModule) SaveDoc(ctx *gin.Context, req *GetUserReq) {

}
