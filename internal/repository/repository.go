package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"posts-log/internal/config"
	"posts-log/pkg/domain"
)

type Logs interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type Repository struct {
	Logs Logs
}

func NewRepository(cfg *config.Config, db *mongo.Database) *Repository {
	return &Repository{
		Logs: NewLogsRepository(cfg, db),
	}
}
