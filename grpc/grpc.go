package grpc

import (
	"github.com/asadbek21coder/catalog/service/config"
	"github.com/asadbek21coder/catalog/service/genproto/book_service"
	"github.com/asadbek21coder/catalog/service/grpc/service"
	"github.com/asadbek21coder/catalog/service/pkg/logger"
	"github.com/asadbek21coder/catalog/service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	book_service.RegisterServiceServer(grpcServer, service.NewServiceServer(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}
