package service

import (
	"context"

	"github.com/asadbek21coder/catalog/service/config"
	pb "github.com/asadbek21coder/catalog/service/genproto/book_service"
	"github.com/asadbek21coder/catalog/service/pkg/logger"
	"github.com/asadbek21coder/catalog/service/storage"
)

type service struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedServiceServer
}

func NewServiceServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *service {
	return &service{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *service) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Books, error) {
	return nil, nil
}

func (s *service) GetById(ctx context.Context, req *pb.Id) (*pb.Book, error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, req *pb.Id) (*pb.Book, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	return nil, nil
}
