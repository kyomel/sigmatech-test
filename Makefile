.PHONY: generate

migrate-up:
	migrate -path pkg/db/migration -database "postgresql://postgres:password@localhost:5432/sigmatech-test?sslmode=disable" -verbose up

migrate-down:
	migrate -path pkg/db/migration -database "postgresql://postgres:password@localhost:5432/sigmatech-test?sslmode=disable" -verbose down