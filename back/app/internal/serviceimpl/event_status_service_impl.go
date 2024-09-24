package serviceimpl

import (
	"context"
	"errors"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/services"
	"golang.org/x/xerrors"
)

var _ services.EventStatusService = (*EventStatusServiceImpl)(nil)

type EventStatusServiceImpl struct {
	EventStatusRepository repositories.EventStatusRepository
}

// IsInvoiceCreatable - 請求書の作成ができる状態か判定する
func (i *EventStatusServiceImpl) IsInvoiceCreatable(ctx context.Context, eventID string) (bool, error) {
	_, err := i.EventStatusRepository.GetByEventIDAndStatus(ctx, &repositories.GetByEventIDAndStatusParam{
		EventID: eventID,
		Status:  entities.EventStatusStart,
	})
	if err != nil {
		var rerr repositoryerrors.RepositoryError[repositoryerrors.NotFoundError]
		if errors.As(err, &rerr) {
			return false, nil
		}
		return false, xerrors.Errorf("error in EventStatusServiceImpl.IsInvoiceCreatable: %w", err)
	}
	_, err = i.EventStatusRepository.GetByEventIDAndStatus(ctx, &repositories.GetByEventIDAndStatusParam{
		EventID: eventID,
		Status:  entities.EventStatusInvoiceCreationChecked,
	})
	if err != nil {
		var rerr repositoryerrors.RepositoryError[repositoryerrors.NotFoundError]
		if errors.As(err, &rerr) {
			return true, nil
		}
		return false, xerrors.Errorf("error in EventStatusServiceImpl.IsInvoiceCreatable: %w", err)
	}
	return false, nil
}

// SetBillable - 請求書開始判定済にする
func (i *EventStatusServiceImpl) SetBillable(ctx context.Context, eventID string) error {
	err := i.EventStatusRepository.Create(ctx, entities.NewEventStatus(&entities.NewEventStatusParam{
		EventID: eventID,
		Status:  entities.EventStatusInvoiceCreationChecked,
	}))
	if err != nil {
		return xerrors.Errorf("error in EventStatusServiceImpl.SetBillable: %w", err)
	}
	return nil
}
