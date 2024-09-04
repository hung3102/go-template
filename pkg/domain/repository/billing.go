package repository

import (
	"context"
	firestoreEntity "gcim/example/pkg/adapter/db/fiestore/entity"
	"gcim/example/pkg/adapter/db/fiestore/repository"
	"gcim/example/pkg/domain/entity"
)

func GetBillings(ctx context.Context) ([]entity.Billing, error) {
	result := []entity.Billing{}

	billings, err := repository.GetBillings(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range billings {
		result = append(result, convertBilling(ctx, v))
	}

	return result, nil
}

func convertBilling(_ context.Context, billing firestoreEntity.Billing) entity.Billing {
	return entity.Billing{
		Id:   billing.DocumentId,
		Cost: billing.Cost,
	}
}
