package usecases

import (
	"context"
	"gcim/example/internal/domain/model"
	"gcim/example/internal/domain/repositories"
	"gcim/example/internal/usecases/dto/input"
	"gcim/example/internal/usecases/dto/output"
)

type ICreateTaskInteractor interface {
	Execute(ctx context.Context, r *input.CreateTaskInput) (*output.CreateTaskOutput, error)
}

type CreateTaskInteractor struct {
	taskRepository repositories.TaskRepository
}

func NewCreateTaskInteractor(
	taskRepository repositories.TaskRepository,
) ICreateTaskInteractor {
	return &CreateTaskInteractor{
		taskRepository: taskRepository,
	}
}

func (u *CreateTaskInteractor) Execute(ctx context.Context, input *input.CreateTaskInput,
) (*output.CreateTaskOutput, error) {
	_, err := u.taskRepository.Insert(ctx, &model.Task{
		ID:   input.ID,
		Desc: input.Description,
	})
	if err != nil {
		return nil, err
	}

	output := &output.CreateTaskOutput{
		Task: &model.Task{
			ID:   input.ID,
			Desc: input.Description,
		},
	}

	return output, nil
}
