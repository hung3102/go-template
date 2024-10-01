package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/event_status_repository_mock.go -package=mockrepositories

// EventStatusRepository - event_statusリポジトリ
type EventStatusRepository interface {
	// GetByEventIDAndStatus - 指定したevent_id, statusのレコードを取得する
	GetByEventIDAndStatus(ctx context.Context, param *GetByEventIDAndStatusParam) (*entities.EventStatus, error)
	// Create - レコードを登録する
	Create(ctx context.Context, eventStatus *entities.EventStatus) error
}

// GetByEventIDAndStatusParam - GetByEventIDAndStatusのパラメーター
type GetByEventIDAndStatusParam struct {
	EventID valueobjects.EventID // EventID
	Status  int                  // Status
}
