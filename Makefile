migrate-up:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/roombooking?sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/roombooking?sslmode=disable" down

generate:
	go tool oapi-codegen --config=oapi-codegen.yaml api/openapi.yaml