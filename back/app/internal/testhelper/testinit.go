package testhelper

import (
	"context"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// LoadEnv - .envを読み込む
func LoadEnv(t *testing.T) {
	t.Helper()

	wd, err := os.Getwd()
	i := strings.Index(wd, "/back/")

	wd2 := wd[i:]
	j := strings.Count(wd2, "/")

	env := ".env"
	for k := 0; k < j; k++ {
		env = "../" + env
	}

	err = godotenv.Load(env)
	if err != nil {
		t.Fatalf("読み込み出来ませんでした: %v", err)
	}
}

func FirestoreClient(t *testing.T) *firestore.Client {
	t.Helper()

	ctx := context.Background()
	projectID := os.Getenv("FIRESTORE_PROJECT_ON_EMULATOR")
	options := make([]option.ClientOption, 0)
	options = append(options, option.WithoutAuthentication())
	result, err := firestore.NewClient(ctx, projectID, options...)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	return result
}
