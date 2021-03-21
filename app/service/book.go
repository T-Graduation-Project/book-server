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

func GetBookList() (*protobuf.GetBookListRsp, error) {
	rsp := new(protobuf.GetBookListRsp)
	err := db.Table("book_info").Scan(&rsp.Books)
	return rsp, err
}

func AddBooks(req *protobuf.AddBooksReq) (*protobuf.AddBooksRsp, error) {
	rsp := &protobuf.AddBooksRsp{
		Msg: "Add book err",
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
	rsp.Msg = "Add books success"
	return rsp, nil
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

// 检查书籍是否存在，判断方式：
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
