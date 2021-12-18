package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/web"
)

type UserModule struct{}

func NewUserModule() web.Module {
	return &UserModule{}
}

func (m *UserModule) Init(r web.Router) {
	g := r.Group("api/user")
	{
		g.Handle("GET", "/:id", m.GetUserByID)
	}
}

type GetUserByIDReq struct {
	ID   uint64 `form:"id" json:"id" uri:"id"`
	Name string `form:"name"`
}

func (m *UserModule) GetUserByID(ctx *gin.Context, req *GetUserByIDReq) {
	ctx.JSON(http.StatusOK, fmt.Sprintf("hello %v, %v", req.ID, req.Name))
}
