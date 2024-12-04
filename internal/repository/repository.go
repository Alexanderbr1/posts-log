package repository

import (
	"context"
	"github.com/Alexanderbr1/posts-log/internal/config"
	"github.com/Alexanderbr1/posts-log/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
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
