package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/config"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/internal/middlewares"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/pkg/dashboard"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/template"
)

type echoServer struct {
    conf *config.Config
    app *echo.Echo
}

var (
    server *echoServer
    once sync.Once
)

func NewEchoServer(conf *config.Config) *echoServer {
    echoApp := echo.New()
    echoApp.Logger.SetLevel(log.DEBUG)

    once.Do(func() {
        server = &echoServer{
            app: echoApp,
            conf: config.ConfigGetting(),
        }
    })

    return server
}

func (s * echoServer) Start() {
    timeOutMiddleware := middlewares.TimeOutMiddleware(s.conf.Server.Timeout)
    corsMiddleware := middlewares.CORSMiddleware(s.conf.Server.AllowOrigins)

    s.app.Use(middleware.Recover())
    s.app.Use(middleware.Logger())

    s.app.Use(timeOutMiddleware)
    s.app.Use(corsMiddleware)

    s.app.Static("/dist", ".dist")
    s.app.Static("/assets", "assets")

    // for health check
    s.app.GET("/v1/health", s.healthCheck)

    // Register template templ
    template.NewTemplateRenderer(s.app)

    baseGroup := s.app.Group("")
    dashboard.NewDashboardFrontend(baseGroup)

    // Graceful Shutdown
    quitCh := make(chan os.Signal, 1)
    signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
    go s.gracefullyShutdown(quitCh)

    s.httpListening()
}

func (s *echoServer) gracefullyShutdown(quitCh <-chan os.Signal) {
    ctx := context.Background()

    <-quitCh
    s.app.Logger.Info("Shutting down the service...")

    if err := s.app.Shutdown(ctx); err != nil {
        s.app.Logger.Fatalf("Error: %s", err.Error())
    }
}

func (s *echoServer) httpListening() {
    url := fmt.Sprintf(":%d", s.conf.Server.Port)
    if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
        s.app.Logger.Fatalf("shutting down the server: %v", err)
    }
}

func (s *echoServer) healthCheck(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}
