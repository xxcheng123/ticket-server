package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ticket-server/app/user/rpc/internal/config"
	"ticket-server/app/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewSqlConn("mysql", ""), c.CacheConf),
	}
}
