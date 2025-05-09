package controller

import (
	"books_service/domain"
	"books_service/service/mocks"
	"books_service/web"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
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

func TestBookCreateFailed(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	request := httptest.NewRequest("POST", "http://localhost:8080/v1/book", strings.NewReader("{invalid-json}"))
	recorder := httptest.NewRecorder()

	svc.On("Create", mock.Anything, mock.AnythingOfType("*web.Request")).Return(nil, errors.New("error"))

	ctrl.Create(recorder, request, nil)

	var result web.Response[domain.Domain]
	json.NewDecoder(recorder.Body).Decode(&result)

	assert.Equal(t, http.StatusBadRequest, result.Code)
	assert.Equal(t, "error", result.Message)
	assert.Equal(t, 0, result.Data.Id)
	assert.Equal(t, "Error", result.Status)

}

func TestBookUpdateSuccess(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	reqBody := `{"Author":"Test 2", "Title": "Testing 2"}`
	request := httptest.NewRequest("PUT", "http://localhost:8080/v1/book/1", strings.NewReader(reqBody))
	recorder := httptest.NewRecorder()
	params := httprouter.Params{
		{Key: "id", Value: "1"},
	}
	expected := &domain.Domain{Id: 1, Author: "Test 2", Title: "Testing 2"}

	svc.On("Update", mock.Anything, mock.Anything).Return(expected, nil)

	ctrl.Update(recorder, request, params)

	var result web.Response[domain.Domain]
	json.NewDecoder(recorder.Body).Decode(&result)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, "OK", result.Status)
	assert.Equal(t, expected.Id, result.Data.Id)
	assert.Equal(t, expected.Author, result.Data.Author)
	assert.Equal(t, expected.Title, result.Data.Title)

}

func TestBookUpdateFailed(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	reqBody := `{"Author":"Test","Title":"Testing"}`
	request := httptest.NewRequest("PUT", "http://localhost:8080/v1/book/2", strings.NewReader(reqBody))
	recorder := httptest.NewRecorder()

	params := httprouter.Params{
		{Key: "id", Value: "2"},
	}

	svc.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("id not found"))

	ctrl.Update(recorder, request, params)

	var result web.Response[domain.Domain]
	json.NewDecoder(recorder.Body).Decode(&result)
	assert.Equal(t, 400, result.Code)
	assert.Equal(t, "Error", result.Status)
	assert.Equal(t, "id not found", result.Message)
	assert.Equal(t, domain.Domain{Id: 0, Author: "", Title: ""}, result.Data)
}

func TestBookDeleteSuccess(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/v1/book/3", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{
		{Key: "id", Value: "2"},
	}

	svc.On("Delete", mock.Anything, mock.Anything).Return(nil)

	ctrl.Delete(recorder, request, params)

	assert.Equal(t, 204, recorder.Result().StatusCode)
}

func TestBookDeleteFailed(t *testing.T) {
	svc := mocks.NewBookService(t)
	ctrl := NewBookController(svc)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/v1/book/4", nil)
	recorder := httptest.NewRecorder()

	params := httprouter.Params{
		{
			Key: "id", Value: "4",
		},
	}

	svc.On("Delete", mock.Anything, mock.Anything).Return(errors.New("id not found"))

	ctrl.Delete(recorder, request, params)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}
