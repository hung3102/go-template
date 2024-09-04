package billing

import (
	"context"
	"gcim/example/pkg/domain/entity"
	"gcim/example/pkg/domain/repository"
)

func GetBillings(ctx context.Context) ([]entity.Billing, error) {
	billings, err := repository.GetBillings(ctx)
	if err != nil {
		return nil, err
	}

	return billings, nil
}
