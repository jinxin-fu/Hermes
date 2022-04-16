package main

import (
	"Hermes/api/inter/config"
	"Hermes/api/inter/handler"
	"Hermes/api/inter/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/hermes-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	c.Name = "hermes-api"
	c.Host = "0.0.0.0"
	c.Port = 8888
	c.Transform.Etcd = discov.EtcdConf{
		Hosts: []string{
			"192.168.2.62:2379",
		},
		Key: "transform.rpc",
	}
	c.Log.Path = "logs"
	//conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
