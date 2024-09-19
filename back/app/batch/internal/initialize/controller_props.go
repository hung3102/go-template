package initialize

import (
	"github.com/topgate/gcim-temporary/back/app/batch/internal/config"
	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/props"
)

// NewControllerProps - コントローラでつかうプロパティのコンストラクタ
func NewControllerProps(_ *config.Config, useCases UseCases) *props.ControllerProps {
	cp := &props.ControllerProps{
		EventUseCase: useCases.EventUsecase,
	}

	return cp
}
