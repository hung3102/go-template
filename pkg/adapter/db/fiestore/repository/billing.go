package repository

import (
	"context"
	"fmt"
	"gcim/example/pkg/adapter/db/fiestore/entity"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetBillings(ctx context.Context) ([]entity.Billing, error) {
	result := []entity.Billing{}

	// json := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	// jwtConfig, err := google.JWTConfigFromJSON([]byte(json))
	// if err != nil {
	// 	return nil, err
	// }

	// token := jwtConfig.TokenSource(ctx)

	// client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"), option.WithTokenSource(token))
	fmt.Println(os.Getenv("PROJECT_ID"))
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iter := client.Collection("billings").Documents(ctx)
	defer iter.Stop()

	for {
		var billing entity.Billing
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		doc.DataTo(billing)
		result = append(result, billing)
	}

	return result, nil
}
