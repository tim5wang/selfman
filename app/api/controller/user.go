package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/util"
)

type UserModule struct{}

func NewUserModule() util.Module {
	return &UserModule{}
}

func (m *UserModule) Init(r gin.IRouter) {
	g := r.Group("api/user")
	{
		g.Handle("GET", "/:id", m.GetUserByID)
	}
}

type GetUserByIDReq struct {
	ID   uint64 `form:"id" json:"id" uri:"id"`
	Name string `form:"name"`
}

func (m *UserModule) GetUserByID(ctx *gin.Context) {
	//req := &GetUserByIDReq{}
	//err := util.BindJsonReq(ctx, req)
	//if err != nil {
	//	_ = ctx.AbortWithError(1002, errors.New("hello gin !"))
	//}
	//ctx.JSON(http.StatusOK, fmt.Sprintf("hello %v, %v", req.ID, req.Name))
	req := map[string]interface{}{}
	err := util.BindJsonReq(ctx, req)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, req)
}
