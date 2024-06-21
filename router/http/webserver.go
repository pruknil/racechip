package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"net/http"
	"os"
	"sportbit.com/racechip/logger"
	"sportbit.com/racechip/service"
	"time"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type WebServer struct {
	srv *http.Server
	//srvSSL      *http.Server
	config      Config
	httpService service.IHttpService
	router      *gin.Engine
	log         logger.AppLog
	container   *dig.Container
}

func NewGin(cfg Config, service service.IHttpService, logg logger.AppLog) *WebServer {
	return &WebServer{
		config:      cfg,
		httpService: service,
		log:         logg,
	}
}

func (g *WebServer) initializeRoutes() {
	hn, _ := os.Hostname()
	//g.router.Use(static.Serve("/", static.LocalFile("views/static", true)))
	g.router.POST("/api", g.serviceLocator)

	g.router.GET("/health", func(c *gin.Context) {
		c.String(200, hn)
	})
}

func (g *WebServer) serviceLocator(c *gin.Context) {
	var reqMsg service.ReqMsg
	err := c.BindJSON(&reqMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resMsg := g.httpService.DoService(reqMsg, g.log)
	c.JSON(http.StatusOK, resMsg)

}

func (g *WebServer) Start() {
	//g.container = container
	g.router = gin.Default()

	g.router.Use(GenRsUID())
	g.router.Use(LogRequest(&g.log))
	g.router.Use(LogResponse(&g.log))

	g.initializeRoutes()

	go func() {
		g.srv = &http.Server{
			Addr:         ":" + g.config.Port,
			Handler:      g.router,
			ReadTimeout:  g.config.ReadTimeout,
			WriteTimeout: g.config.WriteTimeout,
			IdleTimeout:  g.config.IdleTimeout,
		}
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (g *WebServer) Shutdown() {
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
