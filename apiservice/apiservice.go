package main

import (
	"flag"
	"fmt"

	"sectran/apiservice/internal/config"
	"sectran/apiservice/internal/handler"
	"sectran/apiservice/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// var configFile = flag.String("f", "D:\\本地项目\\xx\\Sectran\\apiservice\\etc\\sectran-api.yaml", "the config file")
var configFile = flag.String("f", "etc/sectran-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
