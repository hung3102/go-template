package serviceimpl

import (
	"context"
	"errors"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/services"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
	"golang.org/x/xerrors"
)

var _ services.EventStatusService = (*EventStatusServiceImpl)(nil)

type EventStatusServiceImpl struct {
	EventStatusRepository repositories.EventStatusRepository
	uuid                  uuid.UUID
}

// shouldcreateInvoice - 請求書の作成をする必要があるか判定する
func (i *EventStatusServiceImpl) ShouldCreateInvoice(ctx context.Context, eventID string) (bool, error) {
	if _, err := i.EventStatusRepository.GetByEventIDAndStatus(ctx, eventID, entities.EventStatusStart); err != nil {
		var rerr repositoryerrors.RepositoryError[repositoryerrors.NotFoundError]
		if errors.As(err, &rerr) {
			return false, nil
		}
		return false, xerrors.Errorf("error in EventStatusServiceImpl.ShouldcreateInvoice: %w", err)
	}
	if _, err := i.EventStatusRepository.GetByEventIDAndStatus(ctx, eventID, entities.EventStatusInvoiceCreationChecked); err != nil {
		var rerr repositoryerrors.RepositoryError[repositoryerrors.NotFoundError]
		if errors.As(err, &rerr) {
			return true, nil
		}
		return false, xerrors.Errorf("error in EventStatusServiceImpl.ShouldcreateInvoice: %w", err)
	}
	return false, nil
}

// SetBillable - 請求書開始判定済にする
func (i *EventStatusServiceImpl) SetBillable(ctx context.Context, eventID string) error {
	uuid, err := i.uuid.GetUUID()
	if err != nil {
		return xerrors.Errorf("error in EventStatusServiceImpl.SetBillable: %w", err)
	}
	err = i.EventStatusRepository.Create(ctx, entities.NewEventStatus(&entities.NewEventStatusParam{
		ID:      uuid,
		EventID: eventID,
		Status:  entities.EventStatusInvoiceCreationChecked,
	}))
	if err != nil {
		return xerrors.Errorf("error in EventStatusServiceImpl.SetBillable: %w", err)
	}
	return nil
}
