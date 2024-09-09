package presenters

import (
	"gcim/example/internal/api"
	"gcim/example/internal/usecases/dto/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITaskPresenter interface {
	PresentCreateTask(c echo.Context, output *output.CreateTaskOutput) error
}

type TaskPresenter struct{}

func NewTaskPresenter() ITaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) PresentCreateTask(c echo.Context, output *output.CreateTaskOutput) error {
	response := api.CreateTaskResponse{
		ID:   output.Task.ID,
		Desc: output.Task.Desc,
	}

	return c.JSON(http.StatusOK, response)
}
