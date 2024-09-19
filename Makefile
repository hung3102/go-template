GOLANGCI_LINT_VERSION := 1.52.2
MOCKGEN_VERSION := 0.4.0
OAPI_CODEGEN_VERSION := 2.2.0
REVIVE_VERSION := 1.2.5
SWAGGER_CLI_CONTAINER_VERSION := 1.0.0
VOLCAGO_VERSION := 1.11.1

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

dev:
	docker compose up

.PHONY: init
init: bootstrap
	test -f .env || cp .env.template .env
	cd back && make gowork

.PHONY: bootstrap
bootstrap: bootstrap_golangci_lint bootstrap_mockgen bootstrap_oapi bootstrap_revive bootstrap_swagger-cli

.PHONY: bootstrap_volcago
bootstrap_volcago:
	mkdir -p ./bin
	GOBIN=$(PWD)/bin go install github.com/go-generalize/volcago/cmd/volcago@v$(VOLCAGO_VERSION)

.PHONY: bootstrap_oapi
bootstrap_oapi:
	mkdir -p ./bin
	GOBIN=${PWD}/bin/ go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v$(OAPI_CODEGEN_VERSION)

.PHONY: bootstrap_swagger-cli
bootstrap_swagger-cli:
	docker build -f ./docker/swagger_cli/Dockerfile -t swagger_cli:$(SWAGGER_CLI_CONTAINER_VERSION) .

.PHONY: bootstrap_golangci_lint
bootstrap_golangci_lint:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION)

.PHONY: bootstrap_mockgen
bootstrap_mockgen:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install go.uber.org/mock/mockgen@v$(MOCKGEN_VERSION)

.PHONY: bootstrap_revive
bootstrap_revive:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/mgechev/revive@v$(REVIVE_VERSION)

.PHONY: gomock
gomock:
	mkdir -p ./bin
	GOBIN=$(PWD)/bin go install go.uber.org/mock/mockgen@latest

migrate-create:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate create -ext sql -dir ./migrations/ $(name)

migrate-run:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://localhost:3306/mydatabase up

run:
	go run main.go
