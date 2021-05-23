package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/tim5wang/selfman/app/user/api/internal/config"
	"github.com/tim5wang/selfman/app/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
