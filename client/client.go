package main

import (
	"context"
	"fmt"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
)

var (
	log = g.Log()
)

func main() {
	service := micro.NewService(
		micro.Name("book.client"),
	)
	service.Init()
	client := protobuf.NewBooksService("book", service.Client())
	rsp, err := client.GetBookList(context.TODO(), &protobuf.GetBookListReq{
		Name: "test",
	})
	if err != nil {
		log.Info(err)
	}
	fmt.Println(rsp.Books)

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
