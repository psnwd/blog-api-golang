package utils

import (
	"context"
	"os"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/redis/go-redis/v9"

	"go.uber.org/zap"
)

func NewLogger(service string) *zap.Logger {
	env := os.Getenv("ENV")
	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("env", env),
		zap.String("service", service),
	))

	if env == "" || env == "development" {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}

func NewRedisLocksClient(ctx context.Context, maxConns int) (*redis.Client, error) {
	return newRedisClient(ctx, "REDIS_LOCKS_URL", maxConns)
}

func NewRedisQueueClient(ctx context.Context, maxConns int) (*redis.Client, error) {
	return newRedisClient(ctx, "REDIS_QUEUE_URL", maxConns)
}

func newRedisClient(ctx context.Context, env string, maxConns int) (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv(env))
	if err != nil {
		return nil, err
	}
	opt.PoolSize = maxConns

	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	client.AddHook(redisotel.NewTracingHook())

	return client, nil
}
