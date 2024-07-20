package service

import (
	redis "sportbit.com/racechip/backends/redis"
)

type RedisBackendService struct {
	redis.IRedisBackendService
	redis.Config
}

func New(c redis.Config, service redis.IRedisBackendService) IRedisBackendService {
	return &RedisBackendService{
		IRedisBackendService: service,
		Config:               c,
	}
}
