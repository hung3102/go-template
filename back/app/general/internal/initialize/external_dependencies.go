package initialize

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/gcp-kit/gcpen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/topgate/gcim-temporary/back/app/general/internal/config"
	"github.com/topgate/gcim-temporary/back/pkg/environ"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
)

// ExternalDependencies - external dependencies
type ExternalDependencies struct {
	firestoreClient *firestore.Client
	storageClient   *storage.Client
	pubsubClient    *pubsub.Client
	openapi         *openapi3.T
	sesService      *sesv2.Client
}

// NewExternalDependencies - initialize external dependencies
func NewExternalDependencies(ctx context.Context, cfg config.Config) (*ExternalDependencies, error) {
	ed := new(ExternalDependencies)
	var err error

	{
		loader := openapi3.NewLoader()
		spec, err := loader.LoadFromData(OpenAPISpecBin)
		if err != nil {
			return nil, xerrors.Errorf("failed to get openapi specs: %w", err)
		}
		ed.openapi = spec
	}

	// AWS
	{
		if environ.IsLocal() {
			awsCfg, err := awsConfig.LoadDefaultConfig(
				ctx,
				// リージョンは指定しておく
				awsConfig.WithRegion("ap-northeast-1"),
				// switch先ロールのprofileを指定
				awsConfig.WithSharedConfigProfile(""),
				// MFAトークン取得経路を指定
				awsConfig.WithAssumeRoleCredentialOptions(func(options *stscreds.AssumeRoleOptions) {
					options.TokenProvider = func() (string, error) {
						return stscreds.StdinTokenProvider()
					}
				}),
			)
			if err != nil {
				return nil, xerrors.Errorf("failed to initialize aws client: %w", err)
			}
			sesService := sesv2.NewFromConfig(awsCfg)
			ed.sesService = sesService
		}
	}

	// Google Cloud
	{
		projectID := gcpen.ProjectID
		options := make([]option.ClientOption, 0)
		if environ.IsLocal() {
			projectID = cfg.FirestoreProjectOnEmulator
			options = append(options, option.WithoutAuthentication())
		}

		ed.firestoreClient, err = firestore.NewClient(ctx, projectID, options...)
		if err != nil {
			return nil, xerrors.Errorf("failed to initialize firestore client: %w", err)
		}

		options = make([]option.ClientOption, 0)
		if environ.IsLocal() {
			options = append(options, option.WithoutAuthentication())
			options = append(options, option.WithEndpoint("http://"+cfg.StorageEmulatorHost))
		}
		ed.storageClient, err = storage.NewClient(ctx, options...)
		if err != nil {
			return nil, xerrors.Errorf("failed to initialize storage client: %w", err)
		}

		ed.pubsubClient, err = pubsub.NewClient(ctx, projectID)

	}

	return ed, nil
}

// OpenAPISpec - get openapi specs
func (e *ExternalDependencies) OpenAPISpec() *openapi3.T {
	return e.openapi
}
