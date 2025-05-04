package service

import (
	"books_service/domain"
	"books_service/web"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request *web.Request) (*domain.Domain, error)
	Update(ctx context.Context, request *web.Request) (*domain.Domain, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*domain.Domain, error)
	FindAll(ctx context.Context) ([]*domain.Domain, error)
}
