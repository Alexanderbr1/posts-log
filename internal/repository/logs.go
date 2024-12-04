package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"posts-log/internal/config"
	"posts-log/pkg/domain"
)

type LogsRepository struct {
	cfg *config.Config
	db  *mongo.Database
}

func NewLogsRepository(cfg *config.Config, db *mongo.Database) *LogsRepository {
	return &LogsRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *LogsRepository) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := r.db.Collection(r.cfg.DB.Collection).InsertOne(ctx, item)

	return err
}
