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
			return nil, fmt.Errorf("error while scanning books err: %w", err)
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

func (r *serviceRepo) GetAllCategories(ctx context.Context, req *bs.GetAllRequest) (*bs.Categories, error) {
	var (
		resp   bs.Categories
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM book_categories WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning books %w", err)
	}

	query := `SELECT
				category_id,
				name
			FROM book_categories
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
		var category bs.Category

		err = rows.Scan(
			&category.Id,
			&category.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning category err: %w", err)
		}

		resp.Categories = append(resp.Categories, &category)
	}

	return &resp, nil
}

func (r *serviceRepo) GetCategoryById(ctx context.Context, req *bs.Id) (*bs.Category, error) {
	var resp bs.Category
	getCategoryByIdQuery := `SELECT * FROM book_categories WHERE category_id=$1`
	row := r.db.QueryRow(ctx, getCategoryByIdQuery, req.Id)
	err := row.Scan(
		&resp.Id,
		&resp.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning category %w", err)
	}
	return &resp, nil
}

func (r *serviceRepo) CreateCategory(ctx context.Context, req *bs.Category) (res *bs.Id, err error) {
	createCategoryQuery := `INSERT INTO book_categories (name) values ($1) returning category_id`
	row := r.db.QueryRow(ctx, createCategoryQuery, req.Name)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	var Id bs.Id
	err = row.Scan(
		&Id.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning category err: %w", err)
	}
	res = &Id
	return res, nil
}

func (r *serviceRepo) UpdateCategory(ctx context.Context, req *bs.Category) (*bs.Category, error) {
	updateBookQuery := `update book_categories set name=$1 where category_id=$2 returning *`
	row := r.db.QueryRow(ctx, updateBookQuery, req.Name, req.Id)

	var category bs.Category
	err := row.Scan(
		&category.Id,
		&category.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("error while scanning category err: %w", err)
	}
	return &category, nil
}

func (r *serviceRepo) DeleteCategory(ctx context.Context, req *bs.Id) (int32, error) {
	deleteBookQuery := `DELETE FROM book_categories WHERE category_id=$1`
	row, err := r.db.Exec(ctx, deleteBookQuery, req.Id)
	if err != nil {
		return 0, fmt.Errorf("error while deleting category err: %w", err)
	}
	err = errors.New("no category with such id")
	if row.RowsAffected() == 0 {
		return 0, fmt.Errorf("err: %w", err)
	}
	return req.Id, nil
}
