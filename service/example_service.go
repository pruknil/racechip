package service

import (
	"encoding/json"

	"github.com/getlantern/deepcopy"
	"github.com/google/uuid"
	"sportbit.com/racechip/backends/http"
	"sportbit.com/racechip/backends/http/service"
	"sportbit.com/racechip/logger"
)

type ExampleService struct {
	baseService
	service.IHttpBackend
	logger.AppLog

	beRequest  http.EncryptDataRequest
	beResponse *http.EncryptDataResponse

	serviceRequest  ExampleRequest
	serviceResponse ExampleResponse
}

type ExampleRequest struct {
	DPKName string `json:"dPKName"`
	Data    string `json:"data"`
}

type ExampleResponse struct {
	DPKName string `json:"dPKName"`
	EData   string `json:"eData"`
}

func (s *ExampleService) Parse() error {
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

func (s *ExampleService) InputMapping() error {
	s.beRequest.RequestHeader.UserId = s.Request.Header.UserId
	s.beRequest.RequestHeader.RqDt = s.Request.Header.RqDt
	s.beRequest.RequestHeader.FuncNm = "TESTJA"
	s.beRequest.RequestHeader.RqAppId = "0000"
	s.beRequest.RequestHeader.RqUID = uuid.New().String()
	err := deepcopy.Copy(&s.beRequest.EncryptDataBodyRequest, s.serviceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleService) OutputMapping() error {
	err := deepcopy.Copy(&s.serviceResponse, s.beResponse.EncryptDataBodyResponse)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExampleService) getResponse() ResMsg {
	s.Response.Body = s.serviceResponse
	return s.Response
}

func (s *ExampleService) Business() error {
	var err error
	//s.beResponse, err = s.IHttpBackend.TESTJA(s.beRequest)

	s.beResponse = &http.EncryptDataResponse{
		ResponseHeader: http.ResponseHeader{
			FuncNm:     "1",
			RqUID:      "2",
			RsDt:       "3",
			RsAppId:    "4",
			StatusCode: "5",
		},
		EncryptDataBodyResponse: http.EncryptDataBodyResponse{
			DPKName: "abc",
			EData:   "def",
		},
	}
	if err != nil {
		return err
	}
	return nil
}
