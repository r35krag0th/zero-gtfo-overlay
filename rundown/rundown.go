package main

import (
	"flag"
	"fmt"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/config"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/handler"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "rundown/etc/rundown-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
