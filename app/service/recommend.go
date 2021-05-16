package service

import (
	"context"
	"github.com/T-Graduation-Project/book-server/protobuf"
)

func (s *BookApi) RecommendBook(
	ctx context.Context, req *protobuf.RecommendBookReq, rsp *protobuf.RecommendBookRsp) error {
	rsp.Code = 1

	log.Info(rsp)
	rsp.Code = 0
	return nil
}

func cosineSimilarity() {

}
