package service

import (
	"sportbit.com/racechip/backends/http"
)

//import "sportbit.com/racechip/backends/http"

type IHttpBackend interface {
	//DecryptData(http.DecryptDataRequest) (*http.DecryptDataResponse, error)
	//ExampleBackend(http.ExampleBackendRequest) (*http.ExampleBackendResponse, error)
	TESTJA(http.ExampleBackendRequest) (*http.ExampleBackendResponse, error)
}
