package main

import (
	"github.com/T-Graduation-Project/book-server/app/api"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	address := ":8001"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		g.Log().Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	protobuf.RegisterBooksServer(s, api.Books)
	g.Log().Printf("grpc server starts listening on %s", address)
	if err := s.Serve(listen); err != nil {
		g.Log().Fatalf("failed to serve: %v", err)
	}
}
