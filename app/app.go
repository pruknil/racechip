package app

import (
	http2 "sportbit.com/racechip/backends/http"
	"sportbit.com/racechip/logger"
	"sportbit.com/racechip/router/http"
	"sportbit.com/racechip/service"
)

type Config struct {
	logger.AppLog
	Backend
	Router
	Service
}

type Router struct {
	Http http.Config
}

type Service struct {
	Http service.Config
}

type Backend struct {
	Http http2.Config
}
