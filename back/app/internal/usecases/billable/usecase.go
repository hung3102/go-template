package billable

import (
	"log"

	"github.com/go-utils/structs"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/services"
)

// Usecase - 請求書作成の開始判定のユースケース
type Usecase struct {
	createdBy string
	deps      Dependencies
}

// Dependencies - Usecase が依存するもの
type Dependencies struct {
	GCASDashboardAPI      gcasdashboardapi.GCASDashboardAPI
	GCASAPI               gcasapi.GCASAPI
	EventStatusService    services.EventStatusService
	GCASCSPCostRepository repositories.GCASCSPCostRepository
}

// NewUsecase - Usecase のコンストラクタ
func NewUsecase(deps Dependencies) *Usecase {
	if nilFields := structs.GetNilFields(deps); len(nilFields) > 0 {
		log.Fatalf("%+v in Dependencies is nil", nilFields)
	}
	return &Usecase{
		createdBy: "billable.Usecase", // TODO CreatedByの値要検討
		deps:      deps,
	}
}
