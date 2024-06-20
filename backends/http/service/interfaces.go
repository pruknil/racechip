package service

import "sportbit.com/racechip/backends/http"

//import "sportbit.com/racechip/backends/http"

type IHttpBackend interface {
	//DecryptData(http.DecryptDataRequest) (*http.DecryptDataResponse, error)
	//EncryptData(http.EncryptDataRequest) (*http.EncryptDataResponse, error)
	TESTJA(http.SFTP0002I01Request) (*http.SFTP0002I01Response, error)
}
