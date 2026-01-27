dev:
	docker compose up

migrate-create:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate create -ext sql -dir ./migrations/ $(name)

migrate-run:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://localhost:3306/mydatabase up

run:
	go run main.go
	
install-oapi-codegen:
	@which oapi-codegen > /dev/null || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

generate-oapi: install-oapi-codegen
	oapi-codegen -generate types -package api ./internal/api/openapi.yaml > ./internal/api/types_gen.go
	oapi-codegen -generate server -package api ./internal/api/openapi.yaml > ./internal/api/server_gen.go
