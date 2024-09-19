package config

import (
	"context"

	"github.com/heetch/confita"
	"golang.org/x/xerrors"
)

// Config - config base for reading from outside
type Config struct {
	CORSAllowOrigins           string `config:"CORS_ALLOW_ORIGINS"`            // CORS Allow Origins
	SkipJWT                    bool   `config:"SKIP_JWT"`                      // authentication skip
	JWTSecret                  string `config:"JWT_SECRET"`                    // JWT生成用のシークレット
	JWTIssuer                  string `config:"JWT_ISSUER"`                    // JWT生成時のIssuer名
	FromEmailAddress           string `config:"FROM_EMAIL_ADDRESS"`            // 送信元メールアドレス
	FirestoreProjectOnEmulator string `config:"FIRESTORE_PROJECT_ON_EMULATOR"` // on emulator
}

var defaultConfig = Config{
	FirestoreProjectOnEmulator: "test-project",
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
