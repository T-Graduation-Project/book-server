package service

import (
	"github.com/T-Graduation-Project/book-server/protobuf"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBookService(t *testing.T) {

	Convey("Test GetBookList", t, func() {

		Convey("调用函数", func() {
			rsp, err := GetBookList()
			if err != nil {
				log.Fatalf("Get book list error %v", err)
			}

			Convey("判断返回码 == 0", func() {
				So(rsp.Code, ShouldEqual, 0)
			})
		})
	})

	Convey("Test AddBooks", t, func() {
		book1 := &protobuf.Book{
			Name:         "test1",
			Author:       "tianhao1",
			Publish:      "hhh",
			Introduction: "hhh",
			Number:       5,
			ISBN:         "1",
		}
		book2 := &protobuf.Book{
			Name:         "test2",
			Author:       "tianhao2",
			Publish:      "hhh",
			Introduction: "hhh",
			Number:       5,
			ISBN:         "2",
		}
		req := &protobuf.AddBooksReq{
			Books: []*protobuf.Book{
				book1,
				book2,
			},
		}

		Convey("调用函数", func() {
			rsp, err := AddBooks(req)
			if err != nil {
				log.Fatalf("Get book list error %v", err)
			}

			Convey("判断返回码 == 0", func() {
				So(rsp.Code, ShouldEqual, 0)
			})
		})
	})
}
