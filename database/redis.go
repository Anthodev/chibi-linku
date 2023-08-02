package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	host := os.Getenv("REDIS_HOST")

	if host == "" {
		host = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(host),
		Password: "",
		DB:       dbNo,
	})

	return rdb
}

func FlushAll(client *redis.Client) error {
	err := client.FlushAll(Ctx).Err()

	if err != nil {
		return err
	}

	return nil
}

func SaveUrl(client *redis.Client, key string, url string, e int) error {
	expiration := 0

	if e > 0 {
		expiration = e
	}

	err := client.Set(Ctx, key, url, time.Duration(expiration)).Err()

	if err != nil {
		return err
	}

	err = client.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUrl(client *redis.Client, key string) (string, error) {
	val, err := client.Get(Ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
