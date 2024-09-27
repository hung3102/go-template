package invoiceunitlist

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"golang.org/x/xerrors"
)

// List - 団体ごとのCSPごと、と事業者ごとのCSPごとのリストを取得する
func (u *Usecase) List(ctx context.Context, input *Input) ([]*entities.InvoiceUnit, error) {
	invoiceUnits, err := u.deps.ORGCSPAccountRepository.ListInvoiceUnitByEventID(ctx, input.EventID)
	if err != nil {
		return nil, xerrors.Errorf("error in List: %w", err)
	}

	return invoiceUnits, nil
}
