package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
)

// EventsRepository - イベントリポジトリ
type EventsRepository interface {
	// GetByID - IDから取得
	GetByID(ctx context.Context, id string) (*entities.Event, error)
}
