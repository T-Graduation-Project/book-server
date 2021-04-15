package main

import (
	"context"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"time"
)

var (
	log = g.Log()
)

func main() {
	service := micro.NewService(
		micro.Name("book.client"),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs("121.5.238.116:2379"),
		)),
	)
	client := protobuf.NewBooksService("book", service.Client())
	rsp, err := client.GetBookList(context.TODO(), &protobuf.GetBookListReq{
		Name: "test",
	})
	if err != nil {
		log.Info(err)
	}
	log.Info(rsp.Books)
	time.Sleep(1000)
	book := &protobuf.Book{
		Name:         "testmicro",
		Author:       "tianhao1",
		Publish:      "hhh",
		Introduction: "hhh",
		Number:       5,
		ISBN:         "1",
	}
	req := &protobuf.AddBooksReq{
		Books: []*protobuf.Book{
			book,
		},
	}
	r2, err := client.AddBooks(context.TODO(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("####### get server Greeting response:", r2)
	time.Sleep(1000)
	// 更新书籍列表
	book.Name = "update"
	req3 := &protobuf.UpdateBooksReq{
		Books: []*protobuf.Book{
			book,
		},
	}
	r3, err := client.UpdateBooks(context.TODO(), req3)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("####### get server Greeting response:", r3)
	time.Sleep(1000)
	// 删除书籍列表
	req4 := &protobuf.DeleteBooksReq{
		Books: []*protobuf.Book{
			book,
		},
	}
	r4, err := client.DeleteBooks(context.Background(), req4)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("####### get server Greeting response:", r4)
}
