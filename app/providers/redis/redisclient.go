package redisclient

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var redisclient *redis.Client
var redisclientpipeline redis.Pipeliner

type Config struct {
	Host string
	Port int
}

type RedisClient struct {
	client   *redis.Client
	pipeline redis.Pipeliner
}

type RedisClientPipeline struct {
	clientPipeLine redis.Pipeliner
}

func InitRedisClient(ctx context.Context, config Config) error {
	addr := getAddr(config.Host, config.Port)
	redisclient = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	err := redisclient.Conn().Ping(ctx).Err()
	if err != nil {
		return err
	}

	redisclientpipeline = redisclient.Pipeline()

	fmt.Println("Redis connected successfully")
	return nil
}

func getAddr(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

func GetClient() RedisClient {
	return RedisClient{
		client:   redisclient,
		pipeline: redisclientpipeline,
	}
}
