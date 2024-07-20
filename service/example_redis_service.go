package service

import (
	"encoding/json"
	"sportbit.com/racechip/backends/http"

	"github.com/getlantern/deepcopy"
	"github.com/google/uuid"
	"sportbit.com/racechip/backends/redis/service"
	"sportbit.com/racechip/logger"
)

type ExampleRedisBackendService struct {
	baseService
	service.IRedisBackendService
	logger.AppLog

	beRequest  http.ExampleBackendRequest
	beResponse *http.ExampleBackendResponse

	serviceRequest  ExampleRequest
	serviceResponse ExampleResponse
}

func (s *ExampleRedisBackendService) Parse() error {
	jsonString, err := json.Marshal(s.Request.Body)
	if err != nil {
		return err
	}
	s.Log.Trace.Debug(jsonString)
	s.serviceRequest = ExampleRequest{}
	err = json.Unmarshal(jsonString, &s.serviceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleRedisBackendService) InputMapping() error {
	var hd = s.beRequest.RequestHeader
	hd.UserId = s.Request.Header.UserId
	hd.RqDt = s.Request.Header.RqDt
	hd.FuncNm = "TestRedis"
	hd.RqAppId = "0000"
	hd.RqUID = uuid.New().String()
	err := deepcopy.Copy(&s.beRequest.ExampleBackendBodyRequest, s.serviceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleRedisBackendService) OutputMapping() error {
	err := deepcopy.Copy(&s.serviceResponse, s.beResponse)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleRedisBackendService) getResponse() ResMsg {
	s.baseService.getResponse()
	s.Response.Body = s.serviceResponse
	return s.Response
}

func (s *ExampleRedisBackendService) Business() error {
	var err error
	err = s.IRedisBackendService.Set("key", "value", 0)
	if err != nil {
		return err
	}
	return nil
}
