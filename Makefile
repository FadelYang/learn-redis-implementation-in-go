DB_URL=postgres://postgres:postgres@localhost:5432/redis-go?sslmode=disable

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path db/migrations -database "$(DB_URL)" force $(version)

migrate-version:
	migrate -path db/migrations -database "$(DB_URL)" version