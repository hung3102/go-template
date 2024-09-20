package billable

import (
	"log"

	"github.com/go-utils/structs"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
)

// Usecase - 請求書作成の開始判定のユースケース
type Usecase struct {
	deps Dependencies
}

// Dependencies - Usecase が依存するもの
type Dependencies struct {
	GCASDashboardAPI      gcasdashboardapi.GCASDashboardAPI
	GCASAPI               gcasapi.GCASAPI
	EventsRepository      repositories.BaseRepository[entities.Event]
	GCASAccountRepository repositories.GCASAccountRepository
	GCASCSPCostRepository repositories.GCASCSPCostRepository
	UUID                  uuid.UUID
}

// NewUsecase - Usecase のコンストラクタ
func NewUsecase(deps Dependencies) *Usecase {
	if nilFields := structs.GetNilFields(deps); len(nilFields) > 0 {
		log.Fatalf("%+v in Dependencies is nil", nilFields)
	}
	return &Usecase{deps: deps}
}
