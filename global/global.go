package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/samuel/go-zookeeper/zk"
)

var (
	// 服务启动端口配置
	ADDR          = ":9090"
	SYSTEM_NAME   = "wawa 低代码客户端"
	REDIS_ADDR    = "127.0.0.1:6379"
	ZookeeperAddr = "127.0.0.1:2181"
	ZookeeperConn *zk.Conn
	RedisDB       *redis.Client
)
