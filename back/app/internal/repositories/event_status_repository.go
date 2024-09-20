package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/event_status_repository_mock.go -package=mockrepositories

// EventStatusRepository - event_statusリポジトリ
type EventStatusRepository interface {
	// GetByEventDocIDAndStatus - 指定したevent_doc_id, statusのレコードを取得する
	GetByEventDocIDAndStatus(ctx context.Context, eventDocID string, status int) (*entities.EventStatus, error)
	// Create - レコードを登録する
	Create(ctx context.Context, eventStatus *entities.EventStatus) error
}
