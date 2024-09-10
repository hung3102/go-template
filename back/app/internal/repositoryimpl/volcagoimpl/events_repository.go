package volcagoimpl

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
)

type eventImpl struct {
	infra infrastructures.EventRepository
}

// NewEvent - Event リポジトリを生成
func NewEvent(client *firestore.Client) repositories.BaseRepository[entities.Event] {
	return &eventImpl{infra: infrastructures.NewEventRepository(client)}
}

// Create - イベントを作成
func (e *eventImpl) Create(ctx context.Context, event *entities.Event) error {
	_, err := e.infra.Insert(ctx, &volcago.Event{
		ID:           event.ID(),
		BillingMonth: event.BillingMonth(),
		Status:       event.Status(),
		Meta: volcago.Meta{
			CreatedAt: event.Meta().CreatedAt(),
			CreatedBy: event.Meta().CreatedBy(),
			UpdatedAt: event.Meta().UpdatedAt(),
			UpdatedBy: event.Meta().UpdatedBy(),
			DeletedAt: event.Meta().DeletedAt(),
			DeletedBy: event.Meta().DeletedBy(),
		},
	})
	if err != nil {
		return repositoryerrors.NewUnknownError("failed to create event", err)
	}

	return nil
}

// GetByID - ID から取得
func (e *eventImpl) GetByID(ctx context.Context, id string) (*entities.Event, error) {
	event, err := e.infra.Get(ctx, id)
	if err != nil {
		if errors.Is(err, infrastructures.ErrNotFound) {
			return nil, repositoryerrors.NewNotFoundError(err)
		}
		return nil, repositoryerrors.NewUnknownError("failed to get event", err)
	}

	return entities.NewEvent(&entities.NewEventParam{
		ID:           event.ID,
		BillingMonth: event.BillingMonth,
		Status:       event.Status,
		Meta: entities.NewMeta(&entities.NewMetaParam{
			CreatedAt: event.CreatedAt,
			CreatedBy: event.CreatedBy,
			UpdatedAt: event.UpdatedAt,
			UpdatedBy: event.UpdatedBy,
			DeletedAt: event.DeletedAt,
			DeletedBy: event.DeletedBy,
		}),
	}), nil
}
