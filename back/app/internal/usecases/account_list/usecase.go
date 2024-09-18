package accountlist

import (
	"log"

	"github.com/go-utils/structs"
	gcasapi "github.com/topgate/gcim-temporary/back/app/internal/api/gcas_api"
	gcasdashboardapi "github.com/topgate/gcim-temporary/back/app/internal/api/gcas_dashboard_api"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
)

// Usecase - アカウントリストのユースケース
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
