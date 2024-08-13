package server

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/config"
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


