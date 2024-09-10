VOLCAGO_VERSION := 1.11.1
OAPI_CODEGEN_VERSION := 2.2.0

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

dev:
	docker compose up

.PHONY: bootstrap_volcago
bootstrap_volcago:
	mkdir -p ./bin
	GOBIN=$(PWD)/bin go install github.com/go-generalize/volcago/cmd/volcago@v$(VOLCAGO_VERSION)

.PHONY: bootstrap_oapi
bootstrap_oapi:
	mkdir -p ./bin
	GOBIN=${PWD}/bin/ go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v$(OAPI_CODEGEN_VERSION)

.PHONY: gomock
gomock:
	mkdir -p ./bin
	GOBIN=$(PWD)/bin go install go.uber.org/mock/mockgen@latest

graphql-gen:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.49
	go get github.com/99designs/gqlgen@v0.17.49
	go run github.com/99designs/gqlgen generate

migrate-create:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate create -ext sql -dir ./migrations/ $(name)

migrate-run:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://localhost:3306/mydatabase up
