package service

import (
	"books_service/repository/mocks"
	"testing"
)

func TestCreateBookSuccess(t *testing.T) {
	mock := new(mocks.BookRepository)
	svc := &BookServiceImpl{repo: mock}
}
