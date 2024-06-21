package service

import (
	"context"

	"sportbit.com/racechip/logger"
)

type IHttpService interface {
	DoService(ReqMsg, logger.AppLog) ResMsg
}

type IServiceTemplate interface {
	Parse() error
	Validate() error
	OutputMapping() error
	InputMapping() error
	Business() error
	setRequest(ReqMsg) error
	getResponse() ResMsg
	DoService(req ReqMsg, service IServiceTemplate) (ResMsg, error)
	setLog(appLog logger.AppLog)
	setContext(context.Context)
	getContext() context.Context
	//ConvertStringEncoding(fromEncoding string, toEncoding string, input string) (string, error)
}
