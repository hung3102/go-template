package config

import (
	"context"

	"github.com/heetch/confita"
	"golang.org/x/xerrors"
)

// Config - config base for reading from outside
type Config struct {
	CORSAllowOrigins           string `config:"CORS_ALLOW_ORIGINS"`             // CORS Allow Origins
	SkipJWT                    bool   `config:"SKIP_JWT"`                       // authentication skip
	JWTSecret                  string `config:"JWT_SECRET"`                     // JWT生成用のシークレット
	JWTIssuer                  string `config:"JWT_ISSUER"`                     // JWT生成時のIssuer名
	BucketName                 string `config:"BUCKET_NAME"`                    // Storageのバケット名
	FirestoreProjectOnEmulator string `config:"FIRESTORE_PROJECT_ON_EMULATOR"`  // on emulator
	StorageEmulatorHost        string `config:"FIREBASE_STORAGE_EMULATOR_HOST"` // on emulator
}

var defaultConfig = Config{
	FirestoreProjectOnEmulator: "test-project",
	StorageEmulatorHost:        "firebase:9199",
}

// ReadConfig - read config from environment variables
func ReadConfig(ctx context.Context) (*Config, error) {
	loader := confita.NewLoader()

	cfg := defaultConfig
	if err := loader.Load(ctx, &cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
