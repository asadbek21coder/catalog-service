package service

import (
	"context"

	"github.com/asadbek21coder/catalog/service/config"
	bs "github.com/asadbek21coder/catalog/service/genproto/book_service"
	"github.com/asadbek21coder/catalog/service/pkg/logger"
	"github.com/asadbek21coder/catalog/service/storage"
)

type service struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	bs.UnimplementedServiceServer
}

func NewServiceServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *service {
	return &service{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *service) GetAll(ctx context.Context, req *bs.GetAllRequest) (*bs.Books, error) {
	resp, err := s.strg.Service_I().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Get all books", logger.Error(err))
		return nil, err
	}
	return resp, nil
}
func (s *service) Create(ctx context.Context, req *bs.Book) (*bs.Book, error) {
	resp, err := s.strg.Service_I().Create(ctx, req)
	if err != nil {
		s.log.Error("Create Book", logger.Error(err))
		return nil, err
	}

	return &bs.Book{
		Id:         resp.Id,
		Name:       req.Name,
		CategoryId: req.CategoryId,
		Price:      req.Price,
		Author:     req.Author,
	}, nil
}

func (s *service) GetById(ctx context.Context, req *bs.Id) (*bs.Book, error) {
	resp, err := s.strg.Service_I().GetById(ctx, req)
	if err != nil {
		s.log.Error("Get Book by given id", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *service) Update(ctx context.Context, req *bs.Book) (*bs.Book, error) {
	resp, err := s.strg.Service_I().Update(ctx, req)
	if err != nil {
		s.log.Error("Update Book", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *service) Delete(ctx context.Context, req *bs.Id) (*bs.Id, error) {
	id, err := s.strg.Service_I().Delete(ctx, req)
	if err != nil {
		s.log.Error("Delete Book by given id", logger.Error(err))
		return nil, err
	}
	return &bs.Id{Id: id}, nil
}

func (s *service) GetAllCategories(ctx context.Context, req *bs.GetAllRequest) (*bs.Categories, error) {
	resp, err := s.strg.Service_I().GetAllCategories(ctx, req)
	if err != nil {
		s.log.Error("Get all category", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *service) CreateCategory(ctx context.Context, req *bs.Category) (*bs.Category, error) {
	resp, err := s.strg.Service_I().CreateCategory(ctx, req)
	if err != nil {
		s.log.Error("Create category", logger.Error(err))
		return nil, err
	}

	return &bs.Category{
		Id:   resp.Id,
		Name: req.Name,
	}, nil
}

func (s *service) GetCategoryById(ctx context.Context, req *bs.Id) (*bs.Category, error) {
	resp, err := s.strg.Service_I().GetCategoryById(ctx, req)
	if err != nil {
		s.log.Error("Get Category by given id", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *service) UpdateCategory(ctx context.Context, req *bs.Category) (*bs.Category, error) {
	resp, err := s.strg.Service_I().UpdateCategory(ctx, req)
	if err != nil {
		s.log.Error("Update Category", logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *service) DeleteCategory(ctx context.Context, req *bs.Id) (*bs.Id, error) {
	id, err := s.strg.Service_I().DeleteCategory(ctx, req)
	if err != nil {
		s.log.Error("Delete Category by given id", logger.Error(err))
		return nil, err
	}
	return &bs.Id{Id: id}, nil
}
