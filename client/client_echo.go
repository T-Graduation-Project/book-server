package main

import (
	"github.com/T-Graduation-Project/book-server/protobuf"
	"golang.org/x/net/context"
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

	r, err := c.GetBookList(context.Background(), &protobuf.GetBookListReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r)

	book := &protobuf.Book{
		Name:         "test",
		Author:       "tianhao",
		Publish:      "hhh",
		Introduction: "hhh",
		Number:       5,
		ISBN:         "321",
	}
	req := &protobuf.AddBooksReq{
		Books: []*protobuf.Book{
			book,
		},
	}
	r2, err := c.AddBooks(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r2)

	//r2, err := c.ReturnBook(context.Background(), &protobuf.ReturnBookReq{UserId: 1, BookId: 1})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r2)
}
