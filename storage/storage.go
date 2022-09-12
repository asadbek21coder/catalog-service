package storage

import (
	"context"

	bs "github.com/asadbek21coder/catalog/service/genproto/book_service"
)

type StorageI interface {
	Service_I() Service_I
}

type Service_I interface {
	GetAll(ctx context.Context, req *bs.GetAllRequest) (*bs.Books, error)
	GetById(ctx context.Context, req *bs.Id) (*bs.Book, error)
	Update(ctx context.Context, req *bs.Book) (*bs.Book, error)
	Delete(ctx context.Context, req *bs.Id) (int32, error)
	Create(ctx context.Context, req *bs.Book) (*bs.Id, error)
}
