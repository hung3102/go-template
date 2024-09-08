package volcagoimpl

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	volcago "github.com/topgate/gcim-temporary/back/app/internal/volcago/repositories"
)

type eventImpl struct {
	repo volcago.EventRepository
}

// NewEvent - Event リポジトリを生成
func NewEvent(client *firestore.Client) repositories.EventsRepository {
	return &eventImpl{repo: volcago.NewEventRepository(client)}
}

// GetByID - ID から取得
func (e *eventImpl) GetByID(ctx context.Context, id string) (*entities.Event, error) {
	event, err := e.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, volcago.ErrNotFound) {
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
