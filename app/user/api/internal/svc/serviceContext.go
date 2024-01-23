package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ticket-server/app/user/api/internal/config"
	"ticket-server/app/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
