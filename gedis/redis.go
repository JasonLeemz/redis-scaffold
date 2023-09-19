package gedis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var (
	redisClient      *redis.Client
	redisClient_Once sync.Once
)

func Redis() *redis.Client {
	redisClient_Once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     "192.168.31.92:6379",
			Username: "",
			Password: "MhxzKhl*", // no password set
			DB:       0,          // use default DB

			// 连接池容量及闲置连接数量
			PoolSize:     15,  // 连接池数量
			MinIdleConns: 10,  // 最小空闲连接数
			MaxIdleConns: 100, // 最大空闲连接数

			ConnMaxIdleTime: 0,
			ConnMaxLifetime: 0,

			// 超时
			DialTimeout:  5 * time.Second, // 连接建立超时时间
			ReadTimeout:  3 * time.Second, // 读 超时,默认3s, -1表示取消读超时
			WriteTimeout: 3 * time.Second, // 写 超时,默认等于读超时
			PoolTimeout:  4 * time.Second, // 当所有连接都处于繁忙状态时,客户端等待可用连接的最大等待时长

			// 命令执行失败时的重试策略
			MaxRetries:      0,                      // 命令执行失败时,最多重试多少次,默认0不重试
			MinRetryBackoff: 8 * time.Millisecond,   // 每次计算重试间隔时间的下线,默认8毫秒,-1表示取消间隔
			MaxRetryBackoff: 512 * time.Millisecond, // 每次计算重试间隔时间的上限,默认512毫秒,-1表示取消间隔

			ClientName:          "",
			Dialer:              nil,
			OnConnect:           nil,
			Protocol:            0,
			CredentialsProvider: nil,

			PoolFIFO:  false,
			TLSConfig: nil,
			Limiter:   nil,

			ContextTimeoutEnabled: false,
		})

		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("connect error:%s", err.Error()))
		}
		log.Println(pong)
	})

	return redisClient
}
