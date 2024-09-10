package initialize

import (
	"github.com/topgate/gcim-temporary/back/app/general/internal/config"
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/event"
)

// UseCases - usecaseをまとめた構造体
type UseCases struct {
	//AuthenticationUsecase *authentication.Usecase
	EventUsecase *event.Usecase
	//SessionUsecase        *session.Usecase
}

// NewUseCases - usecaseのコンストラクタ
func NewUseCases(cfg config.Config, useCaseDependencies UseCaseDependencies) *UseCases {
	uc := new(UseCases)
	uc.EventUsecase = event.NewUsecase(
		event.Dependencies{
			EventsRepository: useCaseDependencies.EventRepository,
		},
	)
	// AuthenticationUsecase
	// SessionUsecase

	// 他の Usecase を追加していく

	return uc
}
