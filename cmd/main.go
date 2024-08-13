package main

import (
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/config"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/internal/server"
)

func main() {
    conf := config.ConfigGetting()
    s := server.NewEchoServer(conf)
    s.Start()
}
