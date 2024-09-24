package event

import (
	"context"
)

// GetByID - イベントIDからイベントを取得する
func (u *Usecase) GetEventByID(ctx context.Context, id string) (*Event, error) {
	event, err := u.deps.EventsRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:           event.ID(),
		BillingMonth: event.BillingMonth(),
		CreatedAt:    event.Meta().CreatedAt(),
	}, nil
}
