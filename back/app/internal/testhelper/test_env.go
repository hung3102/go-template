package testhelper

import (
	"os"
	"testing"
)

// SetEnv - 環境変数を設定する
func SetEnv(t *testing.T) {
	t.Helper()
	os.Setenv("IS_LOCAL", "true")
	os.Setenv("BUCKET_NAME", "test-project.appspot.com")
	os.Setenv("FIRESTORE_PROJECT_ON_EMULATOR", "test-project")
}
