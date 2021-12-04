package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/util"
)

type UserModule struct {
}

func NewUserModule() util.Module {
	return &UserModule{}
}

func (m *UserModule) Init(r gin.IRouter) {
	g := r.Group("api/user")
	{
		g.Handle("GET", "/:id", m.GetUserByID)
	}
}

func (m *UserModule) GetUserByID(ctx *gin.Context) {
	_ = ctx.AbortWithError(1000, errors.New("hello gin !"))
}
