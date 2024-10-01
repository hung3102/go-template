package services

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../serviceimpl/mocks/event_status_service_mock.go -package=mockservices

// EventStatusService - EventStatusサービス
type EventStatusService interface {
	// IsInvoiceCreatable - 請求書の作成ができる状態か判定する
	IsInvoiceCreatable(ctx context.Context, eventID valueobjects.EventID) (bool, error)
	// SetBillable - 請求書開始判定済にする
	SetBillable(ctx context.Context, eventID valueobjects.EventID) error
}
