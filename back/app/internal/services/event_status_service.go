package services

import "context"

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../serviceimpl/mocks/event_status_service_mock.go -package=mockservices

// EventStatusService - EventStatusサービス
type EventStatusService interface {
	// ShouldcreateInvoice - 請求書の作成をする必要があるか判定する
	ShouldcreateInvoice(ctx context.Context, eventDocID string) (bool, error)
	// SetInvoiceCreationChecked - 請求書開始判定済にする
	SetInvoiceCreationChecked(ctx context.Context, eventDocID string) error
}
