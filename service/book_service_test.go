package service

import (
	"books_service/domain"
	"books_service/repository/mocks"
	"books_service/web"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBookSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := mocks.NewBookRepository(t)
	svc := &BookServiceImpl{repo: repo, db: db}

	ctx := context.Background()
	request := &web.Request{Title: "Testing"}
	expected := &domain.Domain{Title: "Testing"}

	repo.On("Create", ctx, mock.AnythingOfType("*sql.Tx"), expected).Return(expected, nil)

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	result, err := svc.Create(ctx, request)
	if err != nil {
		t.Fatal("diharapkan tidak ada error, tapi mendapat:", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.Title, result.Title)
	repo.AssertExpectations(t)
}

func TestCreateBookFailed(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := mocks.NewBookRepository(t)
	svc := NewBookServiceImpl(repo, db)

	ctx := context.Background()
	request := &web.Request{Title: ""}
	expected := &domain.Domain{Title: ""}
	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()
	expectedErr := errors.New("gagal menyimpan data")
	repo.On("Create", ctx, mock.AnythingOfType("*sql.Tx"), expected).Return(nil, expectedErr)

	result, err := svc.Create(ctx, request)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "gagal menyimpan data", err.Error())
	repo.AssertExpectations(t)
}

func TestUpdateBookSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := mocks.NewBookRepository(t)
	svc := NewBookServiceImpl(repo, db)

	ctx := context.Background()
	request := &web.Request{Author: "Tester", Title: "Testing"}
	expected := &domain.Domain{Author: "Tester", Title: "Testing"}

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	repo.On("Update", ctx, mock.AnythingOfType("*sql.Tx"), expected).Return(expected, nil)

	result, err := svc.Update(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}
