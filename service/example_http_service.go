package service

import (
	"encoding/json"

	"github.com/getlantern/deepcopy"
	"github.com/google/uuid"
	"sportbit.com/racechip/backends/http"
	"sportbit.com/racechip/backends/http/service"
	"sportbit.com/racechip/logger"
)

type ExampleHttpBackendService struct {
	baseService
	service.IHttpBackend
	logger.AppLog

	beRequest  http.ExampleBackendRequest
	beResponse *http.ExampleBackendResponse

	serviceRequest  ExampleRequest
	serviceResponse ExampleResponse
}

type ExampleRequest struct {
	DPKName string `json:"dPKName"`
	Data    string `json:"data"`
}

type ExampleResponse []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (s *ExampleHttpBackendService) Parse() error {
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

func (s *ExampleHttpBackendService) InputMapping() error {
	var hd = s.beRequest.RequestHeader
	hd.UserId = s.Request.Header.UserId
	hd.RqDt = s.Request.Header.RqDt
	hd.FuncNm = "TESTJA"
	hd.RqAppId = "0000"
	hd.RqUID = uuid.New().String()
	err := deepcopy.Copy(&s.beRequest.ExampleBackendBodyRequest, s.serviceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleHttpBackendService) OutputMapping() error {
	err := deepcopy.Copy(&s.serviceResponse, s.beResponse)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleHttpBackendService) getResponse() ResMsg {
	s.baseService.getResponse()
	s.Response.Body = s.serviceResponse
	return s.Response
}

func (s *ExampleHttpBackendService) Business() error {
	var err error
	s.beResponse, err = s.IHttpBackend.TESTJA(s.beRequest)
	if err != nil {
		return err
	}
	return nil
}
