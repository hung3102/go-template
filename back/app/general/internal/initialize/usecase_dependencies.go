package initialize

import (
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/topgate/gcim-temporary/back/app/general/internal/config"
	"github.com/topgate/gcim-temporary/back/app/internal/authentication"
	jwt "github.com/topgate/gcim-temporary/back/app/internal/authentication/jwtimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/pkg/environ"
	"github.com/topgate/gcim-temporary/back/pkg/mail"
	"github.com/topgate/gcim-temporary/back/pkg/storage"
	"github.com/topgate/gcim-temporary/back/pkg/storage/gcs"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
)

// UseCaseDependencies - 初期化されたユースケースの依存の集合体
type UseCaseDependencies struct {
	EventRepository       repositories.BaseRepository[entities.Event]
	SessionRepository     repositories.BaseRepository[entities.UserSession]
	AuthenticationService authentication.Provider
	MailService           mail.Mail
	StorageService        storage.Provider
	UUID                  uuid.UUID
}

// NewUseCaseDependencies - ユースケースに依存するものの初期化
func NewUseCaseDependencies(cfg config.Config, externalDeps ExternalDependencies) *UseCaseDependencies {
	eventRepository := volcagoimpl.NewEvent(externalDeps.firestoreClient)

	sessionRepository := volcagoimpl.NewUserSession(externalDeps.firestoreClient)

	authenticationService := jwt.NewJWTProvider(
		jwt.NewJWTProviderParam{
			Secret:        []byte(cfg.JWTSecret),
			SigningMethod: gojwt.SigningMethodHS512,
			Issuer:        cfg.JWTIssuer,
		},
	)

	mailService := mail.NewMailSES(&mail.NewMailSESParams{
		SesService:  externalDeps.sesService,
		FromAddress: cfg.FromEmailAddress,
	})

	storageService := gcs.NewProvider(&gcs.NewProviderParams{
		Client:     externalDeps.storageClient,
		BucketName: cfg.BucketName,
		IsLocal:    environ.IsLocal(),
	})

	uuid := uuid.UUID{}

	return &UseCaseDependencies{
		EventRepository:       eventRepository,
		SessionRepository:     sessionRepository,
		AuthenticationService: authenticationService,
		MailService:           mailService,
		StorageService:        storageService,
		UUID:                  uuid,
	}
}
