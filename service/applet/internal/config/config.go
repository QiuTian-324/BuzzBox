package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Sms struct {
		AccessKeyId     string
		AccessKeySecret string
		SignName        string
		TemplateCode    string
		Endpoint        string
	}
	BizRedis redis.RedisConf
	UserRPC  zrpc.RpcClientConf
}
