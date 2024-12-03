package service

import (
	"encoding/json"
	http2 "sportbit.com/racechip/backends/http"
)

func (s *HttpBackendService) TESTJA(req http2.ExampleBackendRequest) (*http2.ExampleBackendResponse, error) {
	var model http2.ExampleBackendResponse
	var builder, err = s.newGetRequestBuilder(s.Config.BackendAPI.ExampleUrl, nil, req)

	if err != nil {
		return nil, err
	}
	builder.
		SetResponseModel(&model).
		SetSecurityTransport().
		SetContentType("application/json")

	err = s.IHttpBackendService.DoRequest(builder)
	if err != nil {
		errX := json.Unmarshal([]byte(err.Error()), &model)
		if errX == nil {
			return nil, err
		}
		return &model, err
	}

	return &model, nil
}
