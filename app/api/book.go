package api

import (
	"context"
	"github.com/T-Graduation-Project/book-server/app/service"
	"github.com/T-Graduation-Project/book-server/protobuf"
)

type BookApi struct {}

var Books = &BookApi{}

// 获取书籍列表
func (s *BookApi) GetBookList(
	ctx context.Context, req *protobuf.GetBookListReq) (*protobuf.GetBookListRsp, error) {
	rsp, err := service.GetBookList()
	return rsp, err
}

func (s *BookApi) AddBooks(
	ctx context.Context, req *protobuf.AddBooksReq) (*protobuf.AddBooksRsp, error) {
	rsp, err := service.AddBooks(req)
	return rsp, err
}

func (s *BookApi) UpdateBooks(
	ctx context.Context, req *protobuf.UpdateBooksReq) (*protobuf.UpdateBooksRsp, error) {
	rsp, err := service.UpdateBooks()
	return rsp, err
}

func (s *BookApi) DeleteBooks(
	ctx context.Context, req *protobuf.DeleteBooksReq) (*protobuf.DeleteBooksRsp, error) {
	rsp, err := service.DeleteBooks()
	return rsp, err
}


