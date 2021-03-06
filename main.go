package main

import (
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/T-Graduation-Project/book-server/service"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

var (
	log = g.Log()
)

func main() {
	server := micro.NewService(
		micro.Name("book"),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs(g.Cfg().GetString("registry_addr")),
		)),
	)
	server.Init()
	protobuf.RegisterBooksHandler(server.Server(), new(service.BookApi))
	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
