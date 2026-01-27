package presenters

import (
	"net/http"
	"template/example/internal/api"
	"template/example/internal/usecases/dto/output"

	"github.com/labstack/echo/v4"
)

type ITaskPresenter interface {
	PresentCreateTask(c echo.Context, output *output.CreateTaskOutput) error
	PresentGetTask(c echo.Context, output *output.GetTaskOutput) error
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

func (p *TaskPresenter) PresentGetTask(c echo.Context, output *output.GetTaskOutput) error {

	response := api.GetTaskResponse{
		TaskId: &output.Task.ID,
		Desc:   &output.Task.Desc,
	}

	return c.JSON(http.StatusOK, response)
}
