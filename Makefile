

migration:
	migrate create -ext sql -dir=db/migrations/ -seq $(name)

roll-back:
	migrate -path db/migrations -database "postgresql://postgres:rfz@localhost:5431/rfz?sslmode=disable" force $(version)

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:rfz@localhost:5431/rfz?sslmode=disable" -verbose down


run:
	cd cmd/api && go run main.go

sqlc generate:
	cd internal/stocks && sqlc generate