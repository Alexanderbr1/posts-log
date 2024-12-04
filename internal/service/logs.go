package service

import (
	"context"
	"posts-log/internal/repository"
	"posts-log/pkg/domain"
	logs "posts-log/pkg/proto"
)

type LogsService struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{
		repo: repo,
	}
}

func (s *LogsService) Insert(ctx context.Context, req *logs.LogRequest) (*logs.Empty, error) {
	item := domain.LogItem{
		Action:    req.GetActions().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return &logs.Empty{}, s.repo.Insert(ctx, item)
}