package http

type IRedisBackendService interface {
	Set(RedisRequestBuilder) error
}
