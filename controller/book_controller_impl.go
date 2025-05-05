package controller

import (
	"books_service/domain"
	"books_service/service"
	"books_service/web"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	svc service.BookService
}

func NewBookController(svc service.BookService) BookController {
	return &BookControllerImpl{
		svc: svc,
	}
}

func (ctrl *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqBody := domain.Domain{}
	json.NewDecoder(r.Body).Decode(&reqBody)

	result, err := ctrl.svc.Create(r.Context(), (*web.Request)(&reqBody))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(web.Response[domain.Domain]{
			Code:    400,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(web.Response[domain.Domain]{
		Code:    201,
		Status:  "OK",
		Message: "Success",
		Data:    *result,
	})
}

func (ctrl *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")
	bookId, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		log.Print(err)
		return
	}

	reqBody := web.Request{}
	json.NewDecoder(r.Body).Decode(&reqBody)
	reqBody.Id = bookId
	result, err := ctrl.svc.Update(r.Context(), &reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(web.Response[domain.Domain]{
			Code:    400,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(web.Response[domain.Domain]{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    *result,
	})
}

func (ctrl *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")
	bookId, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		log.Print(err)
		return
	}

	err = ctrl.svc.Delete(r.Context(), bookId)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(204)
}

func (ctrl *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")
	bookId, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		log.Print(err)
		return
	}

	result, err := ctrl.svc.FindById(r.Context(), bookId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(web.Response[domain.Domain]{
			Code:    400,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(web.Response[domain.Domain]{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    *result,
	})
}

func (ctrl *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	result, err := ctrl.svc.FindAll(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(web.Response[domain.Domain]{
			Code:    400,
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(web.Response[[]*domain.Domain]{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    result,
	})
}
