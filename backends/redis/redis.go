package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	breaker "github.com/sony/gobreaker"
	"sportbit.com/racechip/logger"
	"time"
)

type Client struct {
	*breaker.CircuitBreaker
	Config
	logger.AppLog
	RedisRequestBuilder
}

var ctx = context.Background()

func New(c Config, log logger.AppLog) IRedisBackendService {

	var st breaker.Settings
	st.Name = "Redis"
	st.Timeout = 3 * time.Second
	st.ReadyToTrip = func(counts breaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	cb := breaker.NewCircuitBreaker(st)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rrb := RedisRequestBuilder{
		Req:           "",
		client:        rdb,
		responseModel: nil,
	}

	return &Client{CircuitBreaker: cb, Config: c, AppLog: log, RedisRequestBuilder: rrb}
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type Config struct {
	RedisCfg RedisConfig
}

func (s *Client) Set(key string, value interface{}, expiration time.Duration) error {
	defer func(t time.Time) {
		s.AppLog.Trace.Printf("%s", value)
		s.AppLog.Trace.Printf("Redis.Set elapsed time %.4f ms", float64(time.Since(t).Nanoseconds())/float64(time.Millisecond))

	}(time.Now())
	_, err := s.CircuitBreaker.Execute(func() (i interface{}, e error) {
		return nil, s.RedisRequestBuilder.client.Set(ctx, key, value, expiration).Err()
	})
	return err
}

func (s *Client) Get(key string) (string, error) {
	defer func(t time.Time) {
		s.AppLog.Trace.Printf("%s", key)
		s.AppLog.Trace.Printf("Redis.Get elapsed time %.4f ms", float64(time.Since(t).Nanoseconds())/float64(time.Millisecond))

	}(time.Now())
	val, err := s.CircuitBreaker.Execute(func() (i interface{}, e error) {

		return s.RedisRequestBuilder.client.Get(ctx, "key").Result()

	})
	return fmt.Sprint(val), err
}

type RedisRequestBuilder struct {
	Req           string
	client        *redis.Client
	responseModel interface{}
}
