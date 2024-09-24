package services

import "context"

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../serviceimpl/mocks/event_status_service_mock.go -package=mockservices

// EventStatusService - EventStatusサービス
type EventStatusService interface {
	// ShouldCreateInvoice - 請求書の作成をする必要があるか判定する
	ShouldCreateInvoice(ctx context.Context, eventID string) (bool, error)
	// SetBillable - 請求書開始判定済にする
	SetBillable(ctx context.Context, eventID string) error
}
