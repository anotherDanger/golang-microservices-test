package service

import (
	"books_service/domain"
	"books_service/helper"
	"books_service/repository"
	"books_service/web"
	"context"
	"database/sql"
	"log"
)

type BookServiceImpl struct {
	repo repository.BookRepository
	db   *sql.DB
}

func NewBookServiceImpl(repo repository.BookRepository, db *sql.DB) BookService {
	return &BookServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (svc *BookServiceImpl) Create(ctx context.Context, request *web.Request) (*domain.Domain, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	result, err := svc.repo.Create(ctx, tx, (*domain.Domain)(request))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	helper.NewTx(tx, &err)

	return result, nil
}

func (svc *BookServiceImpl) Update(ctx context.Context, request *web.Request) (*domain.Domain, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	result, err := svc.repo.Update(ctx, tx, (*domain.Domain)(request))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	helper.NewTx(tx, &err)

	return result, nil
}

func (svc *BookServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return err
	}

	err = svc.repo.Delete(ctx, tx, id)
	if err != nil {
		log.Print(err)
		return err
	}
	helper.NewTx(tx, &err)

	return nil
}

func (svc *BookServiceImpl) FindById(ctx context.Context, id int) (*domain.Domain, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	result, err := svc.repo.FindById(ctx, tx, id)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	helper.NewTx(tx, &err)

	return result, nil
}

func (svc *BookServiceImpl) FindAll(ctx context.Context) ([]*domain.Domain, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	results, err := svc.repo.FindAll(ctx, tx)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	helper.NewTx(tx, &err)

	return results, nil
}
