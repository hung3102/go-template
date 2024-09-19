SWAGGER_CLI_CONTAINER_VERSION := 1.0.0

.PHONY: pkg_go_generate
pkg_go_generate:
	cd pkg && go generate ./...

.PHONY: app_go_generate
app_go_generate:
	cd app && go generate ./...

.PHONY: generate
generate: pkg_go_generate app_go_generate server_generate

.PHONY: open_api_bundle_general
open_api_bundle_general:
	docker run \
		--mount "type=bind,source=${PWD}/../docs/openapi/general/,target=/openapi/" \
		--mount "type=bind,source=${PWD}/../back/app/general/internal/initialize/,target=/openapi_bundle/" \
		swagger_cli:$(SWAGGER_CLI_CONTAINER_VERSION)

.PHONY: add_schemes_to_yaml_general
add_schemes_to_yaml_general:
	cd ../scripts && go run ./add_schemas/add_schemas.go \
		--openapi "../docs/openapi/general/openapi.yaml" \
		--schemas-dir "../docs/openapi/general/schemas/" \
		--examples-dir "../docs/openapi/general/examples/" > \
		../docs/openapi/general/openapi-merged.yaml

.PHONY: server_generate_general
server_generate_general: add_schemes_to_yaml_general open_api_bundle_general
	../bin/oapi-codegen -config ../docs/openapi/general/config/oapi.yaml ./app/general/internal/initialize/openapi.json

.PHONY: server_generate
server_generate: server_generate_general

.PHONY: app_lint_golangci_lint
app_lint_golangci_lint:
	../bin/golangci-lint run --config="../.github/.golangci.yml" --fast ./app/...

.PHONY: pkg_lint_golangci_lint
pkg_lint_golangci_lint:
	../bin/golangci-lint run --config="../.github/.golangci.yml" --fast ./pkg/...

.PHONY: lint_revive
lint_revive:
	../bin/revive -config ../.github/.revive.toml -formatter friendly -exclude services/... ./...

.PHONY: lint
lint: lint_revive app_lint_golangci_lint pkg_lint_golangci_lint

.PHONY: app_test
app_test:
	go test -v ./app/... | grep -v 'no test files'

.PHONY: pkg_test
pkg_test:
	go test -v ./pkg/... | grep -v 'no test files'

.PHONY: test
test: pkg_test app_test

.PHONY: gowork
gowork:
	go work init ./app ./pkg && go work sync
