// Code generated by goctl. DO NOT EDIT!
// Source: search.proto

package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"datacenter/search/rpc/internal/config"
	//	searchlogic "datacenter/search/rpc/internal/logic"
	"datacenter/search/rpc/internal/server"
	"datacenter/search/rpc/internal/svc"
	"datacenter/search/rpc/search"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/search.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSearchServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		search.RegisterSearchServer(grpcServer, srv)
	})
	// 后台定时搜索服务
	go func(srv *server.SearchServer) {
		for {
			srv.ArticleInit(context.Background(), &search.Request{
				Once: true,
			})
			logx.Info("定时计划。。。")
			time.Sleep(10 * time.Second)
		}
	}(srv)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
