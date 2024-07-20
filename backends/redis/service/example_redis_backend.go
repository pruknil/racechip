package service

import (
	"encoding/json"
	http2 "sportbit.com/racechip/backends/http"
)

func (s *RedisBackendService) TestRedis(req http2.ExampleBackendRequest) (*http2.ExampleBackendResponse, error) {
	var model http2.ExampleBackendResponse

	err := s.Set("bbb", "&model", 0)
	if err != nil {
		errX := json.Unmarshal([]byte(err.Error()), &model)
		if errX == nil {
			return nil, err
		}
		return &model, err
	}
	return &model, nil
}
