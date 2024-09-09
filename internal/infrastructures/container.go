package infrastructure

import (
	"gcim/example/internal/controllers"
	"gcim/example/internal/domain/repositories"
	"gcim/example/pkg/getdownloadurlexample"
	uploadexample "gcim/example/pkg/uploadExample"

	// "gcim/example/internal/infrastructures/repositories"
	"gcim/example/internal/presenters"
	"gcim/example/internal/usecases"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDB)

	// controllers
	container.Provide(controllers.NewTaskController)

	// presenters
	container.Provide(presenters.NewTaskPresenter)
	container.Provide(presenters.NewErrorPresenter)

	// usecases
	container.Provide(usecases.NewCreateTaskUsecase)

	// repositories
	container.Provide(repositories.NewTaskRepository)

	// sample
	container.Provide(getdownloadurlexample.NewGetDownloadURLExample)
	container.Provide(uploadexample.NewUploadExample)

	return container
}
