package event

import (
	"log"

	"github.com/go-utils/structs"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
)

// Usecase - ユースケース
type Usecase struct {
	deps Dependencies
}

// Dependencies - Usecase が依存するもの
type Dependencies struct {
	EventsRepository repositories.BaseRepository[entities.Event]
}

// NewUsecase - Usecase のコンストラクタ
func NewUsecase(deps Dependencies) *Usecase {
	if nilFields := structs.GetNilFields(deps); len(nilFields) > 0 {
		log.Fatalf("%+v in Dependencies is nil", nilFields)
	}
	return &Usecase{deps: deps}
}
