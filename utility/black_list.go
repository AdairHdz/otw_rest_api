package utility

import (
	"context"
	"time"	
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options {
		Addr: "host.docker.internal:6379",
		Password: "",
		DB: 0,
	})
}

func Save(key string) error {
	ctx := context.Background()
	status := rdb.Set(ctx, key, true, 15 * time.Minute)
	_, err := status.Result()
	return err
}

func Get(key string) (string, bool) {
	ctx := context.Background()
	status := rdb.Get(ctx, key)
	result, err := status.Result()
	
	// If key is not present, or if there was an error while trying to read the key,
	// return "false", meaning that the key could not be found. That's because
	// an error in the redis remote service shouldn't return an error, stoping
	// the auth middleware from validating the token
	if (err != nil && err == redis.Nil) || err != nil {
		return "", false
	}


	return result, true
}