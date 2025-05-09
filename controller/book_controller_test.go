package controller

import (
	"books_service/domain"
	"books_service/service/mocks"
	"books_service/web"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookCreateSuccess(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	reqBody := `{"Author": "Test", "Title": "Testing"}`
	request := httptest.NewRequest("POST", "http://localhost:8080/v1/book", strings.NewReader(reqBody))
	recorder := httptest.NewRecorder()

	expected := &domain.Domain{Id: 1, Author: "Test", Title: "Testing"}
	svc.On("Create", mock.Anything, mock.Anything).Return(expected, nil)

	ctrl.Create(recorder, request, nil)

	var result web.Response[domain.Domain]
	json.NewDecoder(recorder.Body).Decode(&result)

	assert.Equal(t, 201, result.Code)
	assert.Equal(t, "OK", result.Status)
	assert.Equal(t, expected.Id, result.Data.Id)
	assert.Equal(t, expected.Author, result.Data.Author)
	assert.Equal(t, expected.Title, result.Data.Title)
}
