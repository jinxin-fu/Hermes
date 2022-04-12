package main

import (
	"Hermes/rpc/transform/inter/config"
	"Hermes/rpc/transform/inter/server"
	"Hermes/rpc/transform/inter/svc"
	"Hermes/rpc/transform/pb/transform"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//c.Name = "transform.rpc"
	//c.ListenOn = "127.0.0.1:8080"
	//c.Etcd = discov.EtcdConf{
	//	Hosts: []string{
	//		"127.0.0..1:2379",
	//	},
	//	Key: "transform.rpc",
	//}
	//c.DataSource = "root:123456@tcp(192.168.2.64:3306)/gozero"
	//c.Table = "hermesd"
	//rdsConfig := redis.RedisConf{
	//	Host: "192.168.2.64:6379",
	//}
	//c.Cache = cache.CacheConf{
	//	cache.NodeConf{rdsConfig, 100},
	//}

	ctx := svc.NewServiceContext(c)
	svr := server.NewTransformerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformerServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
