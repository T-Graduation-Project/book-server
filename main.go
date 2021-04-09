package main

import (
	"github.com/T-Graduation-Project/book-server/app/service"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
)

var (
	log = g.Log()
)

func main() {
	server := micro.NewService(
		micro.Name("book"),
		micro.Version("latest"),
	)
	server.Init()
	protobuf.RegisterBooksHandler(server.Server(), new(service.BookApi))
	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
