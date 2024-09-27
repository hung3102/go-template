package invoiceunitlist

import (
	"context"

	"golang.org/x/xerrors"
)

// List - 団体ごとのCSPごと、と事業者ごとのCSPごとのリストを取得する
func (u *Usecase) List(ctx context.Context, input *Input) ([]*Output, error) {
	invoiceUnits, err := u.deps.ORGCSPAccountRepository.ListInvoiceUnitByEventID(ctx, input.EventID)
	if err != nil {
		return nil, xerrors.Errorf("error in List: %w", err)
	}

	output := make([]*Output, 0)
	for _, v := range invoiceUnits {
		output = append(output, &Output{
			IsPaymentAgent: v.IsPaymentAgent(),
			Subject:        v.Subject(),
			CSP:            v.CSP(),
		})
	}

	return output, nil
}
