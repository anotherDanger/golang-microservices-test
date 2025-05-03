package repository

import (
	"books_service/domain"
	"context"
	"database/sql"
)

type BookRepository interface {
	Create(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
	Update(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
	Delete(ctx context.Context, sql *sql.Tx, id int) error
	FindById(ctx context.Context, sql *sql.Tx, id int) (*domain.Domain, error)
	FindAll(ctx context.Context, sql *sql.Tx) ([]*domain.Domain, error)
}
