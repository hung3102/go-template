package invoiceunitlist

import (
	"log"

	"github.com/go-utils/structs"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
)

// Usecase - Usecase
type Usecase struct {
	deps Dependencies
}

// Dependencies - Dependencies
type Dependencies struct {
	ORGCSPAccountRepository repositories.ORGCSPAccountCostRepository
}

// NewUsecase - Constructor of Usecase
func NewUsecase(deps Dependencies) *Usecase {
	if nilFields := structs.GetNilFields(deps); len(nilFields) > 0 {
		log.Fatalf("%+v in Dependencies is nil", nilFields)
	}

	return &Usecase{deps: deps}
}
