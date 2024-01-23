package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DBConf struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	CacheConf cache.CacheConf
}
