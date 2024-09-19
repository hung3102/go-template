package props

import (
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/event"
)

// ControllerProps - controller props
type ControllerProps struct {
	EventUseCase *event.Usecase // イベントユースケース
}
