package redis

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/go-redis/redis/v8"
	"github.com/usual2970/gopkg/conf"
)

var client atomic.Value

func GetRedis() (*redis.Client, error) {

	lc := client.Load()
	if lc == nil {
		password := conf.GetString("redis.password")
		db := conf.GetInt("redis.db")

		addr := fmt.Sprintf("%s:%s", conf.GetString("redis.host"), conf.GetString("redis.port"))
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,       // use default DB
			PoolSize: conf.GetInt(`redis.pool_size`),
		})

		client.Store(rdb)
		return rdb, nil
	}

	rs, ok := lc.(*redis.Client)
	if !ok {
		return nil, errors.New("lc not a redis.client")
	}
	return rs, nil
}
