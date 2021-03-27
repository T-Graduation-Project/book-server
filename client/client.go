package main

import (
	"context"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:8001"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewBooksClient(conn)
	//// 获取书籍列表
	//r, err := c.GetBookList(context.Background(), &protobuf.GetBookListReq{})
	//log.Println(r.Books)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r)

	// 添加书籍
	book := &protobuf.Book{
		Name:         "test",
		Author:       "tianhao1",
		Publish:      "hhh",
		Introduction: "hhh",
		Number:       5,
		ISBN:         "112",
	}
	book2 := &protobuf.Book{
		Name:         "test2",
		Author:       "tianhao1",
		Publish:      "hhh",
		Introduction: "hhh",
		Number:       5,
		ISBN:         "113",
	}
	//req := &protobuf.AddBooksReq{
	//	Books: []*protobuf.Book{
	//		book,
	//		book2,
	//	},
	//}
	//r2, err := c.AddBooks(context.Background(), req)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r2)

	// 更新书籍列表
	book.Name = "update"
	book2.Name = "update"
	req3 := &protobuf.UpdateBooksReq{
		Books: []*protobuf.Book{
			book,
			book2,
		},
	}
	r3, err := c.UpdateBooks(context.Background(), req3)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r3)

	//// 删除书籍列表
	//req4 := &protobuf.DeleteBooksReq{
	//	Books: []*protobuf.Book{
	//		book,
	//		book2,
	//	},
	//}
	//r4, err := c.DeleteBooks(context.Background(), req4)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r4)
}
