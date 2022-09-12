package postgres

import (
	"context"

	"github.com/asadbek21coder/catalog/service/config"
	"github.com/asadbek21coder/catalog/service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	service storage.Service_I
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pool,
	}, err
}

func (s *Store) Service_I() storage.Service_I {
	if s.service == nil {
		s.service = NewServiceRepo(s.db)
	}

	return s.service
}
