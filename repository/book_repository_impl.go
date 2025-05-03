package repository

import (
	"books_service/domain"
	"context"
	"database/sql"
	"log"
)

type BookRepositoryImpl struct{}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	return &BookRepositoryImpl{}
}

func (repo *BookRepositoryImpl) Create(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "insert into books(author, title) values(?, ?)"
	result, err := sql.ExecContext(ctx, query, entity.Author, entity.Title)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	response := domain.Domain{
		Id:     int(lastId),
		Author: entity.Author,
		Title:  entity.Title,
	}

	return &response, nil
}

func (repo *BookRepositoryImpl) Update(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "update books set "
	args := []interface{}{}

	if entity.Author != "" {
		query += "author = ?, "
		args = append(args, entity.Author)
	}

	if entity.Title != "" {
		query += "title = ?, "
		args = append(args, entity.Title)
	}

	query = query[:len(query)-2]

	query += " where id = ?"
	args = append(args, entity.Id)

	_, err := sql.ExecContext(ctx, query, args...)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return entity, nil
}

func (repo *BookRepositoryImpl) Delete(ctx context.Context, sql *sql.Tx, id int) error {
	query := "delete from books where id = ?"
	_, err := sql.ExecContext(ctx, query, id)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (repo *BookRepositoryImpl) FindById(ctx context.Context, sql *sql.Tx, id int) (*domain.Domain, error) {
	query := "select * from books where id = ?"
	row := sql.QueryRowContext(ctx, query, id)

	var rowSet domain.Domain
	err := row.Scan(&rowSet.Id, &rowSet.Author, &rowSet.Title)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &rowSet, nil
}

func (repo *BookRepositoryImpl) FindAll(ctx context.Context, sql *sql.Tx) ([]*domain.Domain, error) {
	rows, err := sql.QueryContext(ctx, "select * from books")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var books []*domain.Domain

	for rows.Next() {
		var book domain.Domain
		err := rows.Scan(&book.Id, &book.Author, &book.Title)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
