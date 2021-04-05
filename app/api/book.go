package api

import (
	"context"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var (
	db  = g.DB("default")
	log = glog.New()
)

type BookApi struct{}

func (s *BookApi) GetBookList(
	ctx context.Context, req *protobuf.GetBookListReq, rsp *protobuf.GetBookListRsp) error {
	rsp.Code = 1
	err := db.Table("book_info").Scan(&rsp.Books)
	if err != nil {
		return err
	}
	log.Info(rsp)
	rsp.Code = 0
	return nil
}

func (s *BookApi) AddBooks(
	ctx context.Context, req *protobuf.AddBooksReq, rsp *protobuf.AddBooksRsp) error {
	rsp.Code = 1
	rsp.Msg = "Add books err"
	bookList := req.Books
	log.Info("booklist", bookList)
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return err
		} else {
			if flag == true {
				rsp.Msg = "There is some book existed"
				return err
			}
		}
	}
	r, err := db.Table("book_info").Data(bookList).Insert()
	log.Println("db info:", r)
	if err != nil {
		return err
	}
	rsp.Code = 0
	rsp.Msg = "Add books success"
	return nil
}

func (s *BookApi) UpdateBooks(
	ctx context.Context, req *protobuf.UpdateBooksReq, rsp *protobuf.UpdateBooksRsp) error {
	rsp.Code = 1
	rsp.Msg = "Update books err"
	bookList := req.Books
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return err
		} else {
			if flag == false {
				rsp.Msg = "There is any book not existed"
				return err
			}
		}
	}
	for _, book := range bookList {
		r, err := db.Table("book_info").Data(book).Where(
			"ISBN=?", book.ISBN).Update()
		if err != nil {
			return err
		}
		log.Println("db info:", r)
	}
	rsp.Code = 0
	rsp.Msg = "Update books success"
	return nil
}

func (s *BookApi) DeleteBooks(
	ctx context.Context, req *protobuf.DeleteBooksReq, rsp *protobuf.DeleteBooksRsp) error {
	rsp.Code = 1
	rsp.Msg = "Delete error"
	bookList := req.Books
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return err
		} else {
			if flag == false {
				rsp.Msg = "There is any book not existed"
				return err
			}
		}
	}
	for _, book := range bookList {
		r, err := db.Table(
			"book_info").Data(book).Delete("ISBN=?", book.ISBN)
		if err != nil {
			return err
		}
		log.Println("db info:", r)
	}
	rsp.Code = 0
	rsp.Msg = "Delete books success"
	return nil
}

// 检查某一本书籍是否存在，判断方式：
// 1. 书名 + 作者
// 2. ISBN
func isExisted(isbnCode string) (bool, error) {
	results, err := db.Table("book_info").All("ISBN=?", isbnCode)
	if err != nil {
		log.Error(err)
		return true, err
	}
	if results != nil {
		log.Println("Exist book:", results)
		return true, nil
	}
	return false, nil
}
