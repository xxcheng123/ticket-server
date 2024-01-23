package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ticket-server/app/user/rpc/internal/config"
	"ticket-server/app/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		c.DBConf.User,
		c.DBConf.Password,
		c.DBConf.Host,
		c.DBConf.Port,
		c.DBConf.DBName,
	)
	logx.Debugf("mysql dsn:%s\n", dsn)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewSqlConn("mysql", dsn), c.CacheConf),
	}
}
