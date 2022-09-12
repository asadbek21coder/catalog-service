package storage

import (
	"context"

	pb "github.com/asadbek21coder/catalog/service/genproto/book_service"
)

type StorageI interface {
	Service_I() Service_I
}

type Service_I interface {
	GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Books, error)
	GetById(ctx context.Context, req *pb.Id) (*pb.Book, error)
	Update(ctx context.Context, req *pb.Id) (*pb.Id, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Id, error)
	Create(ctx context.Context, req *pb.Book) (*pb.Id, error)
}
