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

type eventStatusImpl struct {
	infra infrastructures.EventStatusRepository
}

// NewEventStatus - EventStatus リポジトリを生成
func NewEventStatus(client *firestore.Client) repositories.EventStatusRepository {
	return &eventStatusImpl{
		infra: infrastructures.NewEventStatusRepository(client),
	}
}

// // GetByEventIDAndStatus - 指定したevent_id, statusのレコードを取得する
func (e *eventStatusImpl) GetByEventIDAndStatus(ctx context.Context, param *repositories.GetByEventIDAndStatusParam) (*entities.EventStatus, error) {
	id := entities.ToEventStatusID(&entities.NewEventStatusParam{
		EventID: param.EventID,
		Status:  param.Status,
	})
	eventStatus, err := e.infra.Get(ctx, id)
	if err != nil {
		if errors.Is(err, infrastructures.ErrNotFound) {
			return nil, repositoryerrors.NewNotFoundError(err)
		}
		return nil, repositoryerrors.NewUnknownError("failed to get event", err)
	}

	return entities.NewEventStatus(&entities.NewEventStatusParam{
		EventID: eventStatus.EventID,
		Status:  eventStatus.Status,
		Meta: entities.NewMeta(&entities.NewMetaParam{
			CreatedAt: eventStatus.CreatedAt,
			CreatedBy: eventStatus.CreatedBy,
			UpdatedAt: eventStatus.UpdatedAt,
			UpdatedBy: eventStatus.UpdatedBy,
			DeletedAt: eventStatus.DeletedAt,
			DeletedBy: eventStatus.DeletedBy,
		}),
	}), nil
}

// Create - レコードを登録する
func (e *eventStatusImpl) Create(ctx context.Context, eventStatus *entities.EventStatus) error {
	_, err := e.infra.Insert(ctx, &volcago.EventStatus{
		ID:      eventStatus.ID(),
		EventID: eventStatus.EventID(),
		Status:  eventStatus.Status(),
		Meta: volcago.Meta{
			CreatedAt: eventStatus.Meta().CreatedAt(),
			CreatedBy: eventStatus.Meta().CreatedBy(),
			UpdatedAt: eventStatus.Meta().UpdatedAt(),
			UpdatedBy: eventStatus.Meta().UpdatedBy(),
			DeletedAt: eventStatus.Meta().DeletedAt(),
			DeletedBy: eventStatus.Meta().DeletedBy(),
		},
	})
	if err != nil {
		return repositoryerrors.NewUnknownError("failed to create event_status", err)
	}

	return nil
}
