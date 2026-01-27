package billing

import (
	"context"
	"template/example/pkg/domain/entity"
	"template/example/pkg/domain/repository"
)

func GetBillings(ctx context.Context) ([]entity.Billing, error) {
	billings, err := repository.GetBillings(ctx)
	if err != nil {
		return nil, err
	}

	return billings, nil
}
