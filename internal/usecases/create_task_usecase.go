package usecases

import (
	"context"
	"gcim/example/internal/domain/model"
	"gcim/example/internal/domain/repositories"
	"gcim/example/internal/usecases/dto/input"
	"gcim/example/internal/usecases/dto/output"
)

type ICreateTaskUsecase interface {
	Execute(ctx context.Context, r *input.CreateTaskInput) (*output.CreateTaskOutput, error)
}

type CreateTaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewCreateTaskUsecase(
	taskRepository repositories.TaskRepository,
) ICreateTaskUsecase {
	return &CreateTaskUsecase{
		taskRepository: taskRepository,
	}
}

func (u *CreateTaskUsecase) Execute(ctx context.Context, input *input.CreateTaskInput,
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
