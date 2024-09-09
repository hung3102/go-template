package usecases

import (
	"context"
	"gcim/example/internal/domain/repositories"
	"gcim/example/internal/usecases/dto/input"
	"gcim/example/internal/usecases/dto/output"
)

type IGetTaskUsecase interface {
	Execute(ctx context.Context, r *input.GetTaskInput) (*output.GetTaskOutput, error)
}

type GetTaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewGetTaskUsecase(
	taskRepository repositories.TaskRepository,
) IGetTaskUsecase {
	return &GetTaskUsecase{
		taskRepository: taskRepository,
	}
}

func (u *GetTaskUsecase) Execute(ctx context.Context, input *input.GetTaskInput,
) (*output.GetTaskOutput, error) {
	m, err := u.taskRepository.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	output := &output.GetTaskOutput{
		Task: m,
	}

	return output, nil
}
