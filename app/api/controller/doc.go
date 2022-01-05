package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/configservice"
	"github.com/tim5wang/selfman/common/web"
	"github.com/tim5wang/selfman/model"
	"github.com/tim5wang/selfman/service/docservice"
)

type DocModule struct {
	config     *configservice.ConfigService
	docService *docservice.DocService
}

func NewDocModule(config *configservice.ConfigService, docService *docservice.DocService) web.Module {
	return &DocModule{config: config, docService: docService}
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
	model.Doc
}

func (m *DocModule) GetDoc(ctx *gin.Context, req *GetDocReq) {
	err, doc := m.docService.GetDoc(req.DocID)
	web.GeneralResponse(ctx, err, doc)
}

func (m *DocModule) SaveDoc(ctx *gin.Context, req *DocReqRsp) {
	err, doc := m.docService.SaveDoc(&req.Doc)
	web.GeneralResponse(ctx, err, doc)
}

func (m *DocModule) DocList(ctx *gin.Context, req *model.DicListReq) {
	err, data := m.docService.DocList(req)
	web.GeneralResponse(ctx, err, data)
}
