package services

import "context"

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../serviceimpl/mocks/event_status_service_mock.go -package=mockservices

// EventStatusService - EventStatusサービス
type EventStatusService interface {
	// IsInvoiceCreatable - 請求書の作成ができる状態か判定する
	IsInvoiceCreatable(ctx context.Context, eventID string) (bool, error)
	// SetBillable - 請求書開始判定済にする
	SetBillable(ctx context.Context, eventID string) error
}
