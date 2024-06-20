package service

import (
	"encoding/json"
	http2 "sportbit.com/racechip/backends/http"
)

func (s *HttpBackendService) TESTJA(req http2.SFTP0002I01Request) (*http2.SFTP0002I01Response, error) {
	var model http2.SFTP0002I01Response
	var builder, err = s.newPostRequestBuilder(s.CCMSAPI.LocalUrl, req)
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
