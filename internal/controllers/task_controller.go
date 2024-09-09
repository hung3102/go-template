package controllers

import (
	"gcim/example/internal/api"
	"gcim/example/internal/presenters"
	"gcim/example/internal/usecases"
	"gcim/example/internal/usecases/dto/input"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	updateTaskNameUsecase usecases.ICreateTaskUsecase
	taskPresenter         presenters.ITaskPresenter
	errorPresenter        presenters.IErrorPresenter
}

func NewTaskController(
	updateTaskNameUsecase usecases.ICreateTaskUsecase,
	taskPresenter presenters.ITaskPresenter,
	errorPresenter presenters.IErrorPresenter,
) *TaskController {
	return &TaskController{
		updateTaskNameUsecase: updateTaskNameUsecase,
		taskPresenter:         taskPresenter,
		errorPresenter:        errorPresenter,
	}
}

func (c *TaskController) CreateTask(e echo.Context) error {
	req := api.CreateTaskRequest{}
	if err := e.Bind(&req); err != nil {
		return c.errorPresenter.PresentBadRequest(e, "invalid request")
	}

	input := &input.CreateTaskInput{
		ID:          *req.Id,
		Description: *req.Desc,
	}
	ctx := e.Request().Context()

	output, err := c.updateTaskNameUsecase.Execute(ctx, input)
	if err != nil {
		return c.errorPresenter.PresentInternalServerError(e, err)
	}

	return c.taskPresenter.PresentCreateTask(e, output)
}
