package postgres

import (
	"context"
	"errors"
	"fmt"

	bs "github.com/asadbek21coder/catalog/service/genproto/book_service"
	"github.com/asadbek21coder/catalog/service/pkg/helper"
	"github.com/asadbek21coder/catalog/service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceRepo struct {
	db *pgxpool.Pool
}

func NewServiceRepo(db *pgxpool.Pool) storage.Service_I {
	return &serviceRepo{
		db: db,
	}
}

func (r *serviceRepo) GetAll(ctx context.Context, req *bs.GetAllRequest) (*bs.Books, error) {
	var (
		resp   bs.Books
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM books WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning books %w", err)
	}

	query := `SELECT
				id,
				name,
				author,
				category,
				price
			FROM books
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var book bs.Book

		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.Author,
			&book.CategoryId,
			&book.Price,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning profession err: %w", err)
		}

		resp.Books = append(resp.Books, &book)
	}

	return &resp, nil
}

func (r *serviceRepo) GetById(ctx context.Context, req *bs.Id) (*bs.Book, error) {
	var resp bs.Book
	getBookByIdQuery := `SELECT * FROM books WHERE id=$1`
	row := r.db.QueryRow(ctx, getBookByIdQuery, req.Id)
	err := row.Scan(
		&resp.Id,
		&resp.Name,
		&resp.Author,
		&resp.Price,
		&resp.CategoryId,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning book %w", err)
	}
	return &resp, nil
}

func (r *serviceRepo) Create(ctx context.Context, req *bs.Book) (res *bs.Id, err error) {
	createBookQuery := `INSERT INTO books (name,author,category, price) values ($1,$2,$3,$4) returning id`
	row := r.db.QueryRow(ctx, createBookQuery, req.Name, req.Author, req.CategoryId, req.Price)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	var Id bs.Id
	err = row.Scan(
		&Id.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning book err: %w", err)
	}
	res = &Id
	return res, nil
}

func (r *serviceRepo) Update(ctx context.Context, req *bs.Book) (*bs.Book, error) {
	updateBookQuery := `update books set name=$1, author=$2, category=$3, price=$4 where id=$5 returning *`
	row := r.db.QueryRow(ctx, updateBookQuery, req.Name, req.Author, req.CategoryId, req.Price, req.Id)

	var book bs.Book
	err := row.Scan(
		&book.Id,
		&book.Name,
		&book.Author,
		&book.CategoryId,
		&book.Price,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning book err: %w", err)
	}
	return &book, nil
}

func (r *serviceRepo) Delete(ctx context.Context, req *bs.Id) (int32, error) {
	deleteBookQuery := `DELETE FROM books WHERE id=$1`
	row, err := r.db.Exec(ctx, deleteBookQuery, req.Id)
	if err != nil {
		return 0, fmt.Errorf("error while deleting book err: %w", err)
	}
	err = errors.New("no book with such id")
	if row.RowsAffected() == 0 {
		return 0, fmt.Errorf("err: %w", err)
	}
	return req.Id, nil
}
