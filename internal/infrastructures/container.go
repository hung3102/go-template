package infrastructure

import (
	"template/example/internal/controllers"
	"template/example/internal/domain/repositories"
	uploadexample "template/example/pkg/uploadExample"

	// "template/example/internal/infrastructures/repositories"
	"template/example/internal/presenters"
	"template/example/internal/usecases"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDB)
	container.Provide(NewStorage)

	// controllers
	container.Provide(controllers.NewTaskController)
	container.Provide(controllers.NewDownloadUrlController)

	// presenters
	container.Provide(presenters.NewTaskPresenter)
	container.Provide(presenters.NewDownloadUrlPresenter)
	container.Provide(presenters.NewErrorPresenter)

	// usecases
	container.Provide(usecases.NewCreateTaskUsecase)
	container.Provide(usecases.NewGetTaskUsecase)
	container.Provide(usecases.NewGetDownloadUrlUsecase)

	// repositories
	container.Provide(repositories.NewTaskRepository)

	// sample
	container.Provide(uploadexample.NewUploadExample)

	return container
}
