package service

import (
	"fmt"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
)

func GetBookList() (*protobuf.GetBookListRsp, error) {
	rsp := new(protobuf.GetBookListRsp)
	db := g.DB("default")
	err := db.Table("book_info").Scan(&rsp.Books)
	return rsp, err
}

func AddBooks(req *protobuf.AddBooksReq) (*protobuf.AddBooksRsp, error) {
	rsp := new(protobuf.AddBooksRsp)
	db := g.DB("default")
	bookList := req.Books
	fmt.Println("booklist", bookList)
	// INSERT INTO `user`(`name`) VALUES('john_1'),('john_2'),('john_3')
	r, err := db.Table("book_info").Data(bookList).Insert()
	fmt.Println(r, err)
	return rsp, err
}

func UpdateBooks(req *protobuf.UpdateBooksReq) (*protobuf.UpdateBooksRsp, error) {
	rsp := new(protobuf.UpdateBooksRsp)
	//db := g.DB("default")
	return rsp, nil
}

func DeleteBooks(req *protobuf.DeleteBooksReq) (*protobuf.DeleteBooksRsp, error) {
	rsp := new(protobuf.DeleteBooksRsp)
	//db := g.DB("default")
	return rsp, nil
}