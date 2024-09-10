package initialize

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
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
	openapi         *openapi3.T
	sesService      *ses.SES
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
			sessOpts := session.Options{
				Config: aws.Config{
					// リージョンは指定しておく
					Region: aws.String("ap-northeast-1"),
				},

				// switch先ロールのprofileを指定
				Profile: "",

				// MFAトークン取得経路を指定
				AssumeRoleTokenProvider: stscreds.StdinTokenProvider,

				// 有効にすることで、いい感じに設定が取れるらしい
				SharedConfigState: session.SharedConfigEnable,
			}
			s := session.Must(session.NewSessionWithOptions(sessOpts))
			sesService := ses.New(s)
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
			options = append(options, option.WithEndpoint("http://localhost:8080"))
		}

		ed.firestoreClient, err = firestore.NewClient(ctx, projectID, options...)
		if err != nil {
			return nil, xerrors.Errorf("failed to initialize firestore client: %w", err)
		}
	}

	return ed, nil
}

// OpenAPISpec - get openapi specs
func (e *ExternalDependencies) OpenAPISpec() *openapi3.T {
	return e.openapi
}
