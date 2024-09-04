dev:
	docker compose up

graphql-gen:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.49
	go get github.com/99designs/gqlgen@v0.17.49
	go run github.com/99designs/gqlgen generate

migrate-create:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate create -ext sql -dir ./migrations/ $(name)

migrate-run:
	docker run -v ./pkg/adapter/db/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://localhost:3306/mydatabase up