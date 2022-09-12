package main

import (
	"github.com/asadbek21coder/catalog/service/config"
	"github.com/asadbek21coder/catalog/service/pkg/logger"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	// pgStore, err := postgres.NewPostgres(fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresDatabase,
	// ), cfg)
	// if err != nil {
	// 	panic(err)
	// }

	// grpcServer := grpc.SetUpServer(cfg, log, pgStore)

}
