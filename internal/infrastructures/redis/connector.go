package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
	"time"
)

type Connector struct {
	Client *redis.Client
}

func NewRedisConnector() *Connector {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "password",
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		time.Sleep(5 * time.Second)
		err = rdb.Ping(ctx).Err()
		if err != nil {
			logging.Fatal(err, "cloud not connect to redis")
			panic(err)
		}
	}

	return &Connector{
		rdb,
	}
}
