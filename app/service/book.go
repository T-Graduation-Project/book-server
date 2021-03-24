package service

import (
	"fmt"
	"github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var (
	db  = g.DB("default")
	log = glog.New()
)

// 获取书籍列表
func GetBookList() (*protobuf.GetBookListRsp, error) {
	rsp := &protobuf.GetBookListRsp{
		Code: 1,
	}
	err := db.Table("book_info").Scan(&rsp.Books)
	log.Println(rsp)
	rsp.Code = 0
	return rsp, err
}

// 添加书籍（列表）
// 先校验所有要添加的书籍再一次性添加
func AddBooks(req *protobuf.AddBooksReq) (*protobuf.AddBooksRsp, error) {
	rsp := &protobuf.AddBooksRsp{
		Code: 1,
		Msg:  "Add books err",
	}
	bookList := req.Books
	fmt.Println("booklist", bookList)
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return rsp, err
		} else {
			if flag == true {
				rsp.Msg = "There is some book existed"
				return rsp, nil
			}
		}
	}
	r, err := db.Table("book_info").Data(bookList).Insert()
	log.Println("db info:", r)
	if err != nil {
		return rsp, err
	}
	rsp.Code = 0
	rsp.Msg = "Add books success"
	return rsp, nil
}

// 更新书籍信息（列表）
// 先校验所有要更新的书籍再一次性更新
func UpdateBooks(req *protobuf.UpdateBooksReq) (*protobuf.UpdateBooksRsp, error) {
	rsp := &protobuf.UpdateBooksRsp{
		Code: 1,
		Msg:  "Update books err",
	}
	bookList := req.Books
	fmt.Println("booklist", bookList)
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return rsp, err
		} else {
			if flag == false {
				rsp.Msg = "There is any book not existed"
				return rsp, nil
			}
		}
	}
	for _, book := range bookList {
		r, err := db.Table("book_info").Data(book).Where(
			"ISBN=?", book.ISBN).Update()
		if err != nil {
			return rsp, err
		}
		log.Println("db info:", r)
	}
	rsp.Code = 0
	rsp.Msg = "Update books success"
	return rsp, nil
}

// 删除书籍（列表）
func DeleteBooks(req *protobuf.DeleteBooksReq) (*protobuf.DeleteBooksRsp, error) {
	rsp := &protobuf.DeleteBooksRsp{
		Code: 1,
		Msg:  "Delete books err",
	}
	bookList := req.Books
	fmt.Println("booklist", bookList)
	for _, book := range bookList {
		isbnCode := book.ISBN
		flag, err := isExisted(isbnCode)
		if err != nil {
			return rsp, err
		} else {
			if flag == false {
				rsp.Msg = "There is any book not existed"
				return rsp, nil
			}
		}
	}
	for _, book := range bookList {
		r, err := db.Table(
			"book_info").Data(book).Delete("ISBN=?", book.ISBN)
		if err != nil {
			return rsp, err
		}
		log.Println("db info:", r)
	}
	rsp.Code = 0
	rsp.Msg = "Delete books success"
	return rsp, nil
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
