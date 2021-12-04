package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
}

func NewUserModule() *UserModule {
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
