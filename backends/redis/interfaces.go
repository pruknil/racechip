package redis

import "time"

type IRedisBackendService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}
