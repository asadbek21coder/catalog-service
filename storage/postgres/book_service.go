package postgres

import (
	"context"

	pb "github.com/asadbek21coder/catalog/service/genproto/book_service"
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

func (r *serviceRepo) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Books, error) {
	return nil, nil
}

func (r *serviceRepo) GetById(ctx context.Context, req *pb.Id) (*pb.Book, error) {
	return nil, nil
}
func (r *serviceRepo) Create(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	return nil, nil
}
func (r *serviceRepo) Update(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	return nil, nil
}
func (r *serviceRepo) Delete(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	return nil, nil
}
