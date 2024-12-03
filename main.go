package main

import (
	"fmt"
	"github.com/spf13/viper"
	//"github.com/spf13/viper"
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"sportbit.com/racechip/app"
	httpbackend "sportbit.com/racechip/backends/http"
	httpsvc "sportbit.com/racechip/backends/http/service"
	redissvc "sportbit.com/racechip/backends/redis"
	"sportbit.com/racechip/logger"
	"sportbit.com/racechip/router"
	"sportbit.com/racechip/router/http"
	"sportbit.com/racechip/service"
	"syscall"
	"time"
)

func main() {
	container := buildContainer()
	err := invokeContainer(container)
	if err != nil {
		log.Fatal("Invoke Container error")
	}
}

func invokeContainer(container *dig.Container) error {
	return container.Invoke(func(route []router.IRouter) {
		for _, v := range route {
			v.Start()
		}

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-quit

		for _, v := range route {
			v.Shutdown()
		}
	})
}
func errorWrap(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func buildContainer() *dig.Container {
	container := dig.New()
	errorWrap(container.Provide(NewConfig))
	errorWrap(container.Provide(NewLogger))

	errorWrap(container.Provide(NewHttpService))
	errorWrap(container.Provide(NewHttp))
	errorWrap(container.Provide(NewRedisBackend))
	errorWrap(container.Provide(NewHttpBackend))

	//errorWrap(container.Provide(NewDataBase))
	//errorWrap(container.Provide(NewDataService))

	errorWrap(container.Provide(NewRouter))
	return container
}

func NewLogger() logger.AppLog {
	al := logger.New()
	al.Error = al.NewLog("error", os.Getenv("LOG_LEVEL"))
	al.Perf = al.NewLog("perf", os.Getenv("LOG_LEVEL"))
	al.Trace = al.NewLog("trace", os.Getenv("LOG_LEVEL"))
	al.Rest = al.NewLog("rest", os.Getenv("LOG_LEVEL"))
	al.Router = al.NewLog("router", os.Getenv("LOG_LEVEL"))
	al.Socket = al.NewLog("socket", os.Getenv("LOG_LEVEL"))
	al.Redis = al.NewLog("redis", os.Getenv("LOG_LEVEL"))
	return al
}

//================= Start BACKEND Section =================

func NewHttp(cfg app.Config, log logger.AppLog) httpbackend.IHttpBackendService {
	return httpbackend.New(cfg.Backend.Http, log)
}

func NewRedisBackend(cfg app.Config, log logger.AppLog) redissvc.IRedisBackendService {
	return redissvc.New(cfg.Backend.Redis, log)
}

func NewHttpBackend(cfg app.Config, s httpbackend.IHttpBackendService) httpsvc.IHttpBackend {
	return httpsvc.New(cfg.Backend.Http, s)
}

//================= End BACKEND Section =================

// Create all router here eg.. rest, socket, sftp
func NewRouter(httpService service.IHttpService, conf app.Config, log logger.AppLog) []router.IRouter {
	var route []router.IRouter
	route = append(route, http.NewGin(conf.Router.Http, httpService, log))
	return route
}

// service
// Http service
func NewHttpService(httpService httpsvc.IHttpBackend, redis redissvc.IRedisBackendService, cfg app.Config, log logger.AppLog) service.IHttpService {
	return service.HttpService{
		ExampleHttpService:  service.ExampleHttpBackendService{IHttpBackend: httpService, AppLog: log},
		ExampleRedisService: service.ExampleRedisBackendService{IRedisBackendService: redis, AppLog: log},
	}
}
func NewConfig() app.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	read, _ := time.ParseDuration(viper.GetString("server.timeout.read"))
	write, _ := time.ParseDuration(viper.GetString("server.timeout.write"))
	idle, _ := time.ParseDuration(viper.GetString("server.timeout.idle"))
	return app.Config{
		Backend: app.Backend{
			Http: httpbackend.Config{
				BackendAPI: httpbackend.BackendAPI{ExampleUrl: "https://jsonplaceholder.typicode.com/posts"},
				//CCMSAPI: httpbackend.CCMSAPI{
				//	DecryptDataUrl: os.Getenv("CCMS_DECRYPT"),
				//	EncryptDataUrl: os.Getenv("CCMS_ENCRYPT"),
				//	LocalUrl:       os.Getenv("CCMS_SFTP_URL"),
				//},
			},
			Redis: redissvc.Config{
				RedisCfg: redissvc.RedisConfig{
					Addr:     viper.GetString("redis.server"),
					Password: "",
					DB:       0,
				}},
		},
		Router: app.Router{
			Http: http.Config{
				Port:         viper.GetString("server.port"),
				ReadTimeout:  read,
				WriteTimeout: write,
				IdleTimeout:  idle,
			},
		},
		Service: app.Service{
			Http: service.Config{},
		},
	}
}
