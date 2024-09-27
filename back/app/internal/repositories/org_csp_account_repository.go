package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/org_csp_account_repository_mock.go -package=mockrepositories

// ORGCSPAccountRepository - org_csp_account repository
type ORGCSPAccountRepository interface {
	// ListInvoiceUnitByEventID - List invoiceUnit by eventID
	ListInvoiceUnitByEventID(ctx context.Context, eventID string) ([]*entities.InvoiceUnit, error)
}
