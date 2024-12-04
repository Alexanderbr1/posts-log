package repository

import (
	"context"
	"github.com/Alexanderbr1/posts-log/internal/config"
	"github.com/Alexanderbr1/posts-log/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	// Вставляем документ
	insertResult, err := r.db.Collection(r.cfg.DB.Collection).InsertOne(ctx, item)
	if err != nil {
		return err
	}

	// Получаем идентификатор вставленного документа
	insertedID := insertResult.InsertedID

	// Ищем вставленный документ по идентификатору
	var insertedItem domain.LogItem
	err = r.db.Collection(r.cfg.DB.Collection).FindOne(ctx, bson.M{"_id": insertedID}).Decode(&insertedItem)
	if err != nil {
		return err
	}

	// Выводим документ в лог
	log.Printf("Вставленный документ: %+v\n", insertedItem)

	return nil
}
