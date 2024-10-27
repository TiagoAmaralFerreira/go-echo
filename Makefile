createmigration:
	migrate create -ext=sqlc -dir=sqlc/migrations -seq init

migrate:
	migrate -path=sqlc/migrations -database "mysql://root:root@tcp(localhost:3306)/finance" -verbose up

migratedown:
	migrate -path=sqlc/migrations -database "mysql://root:root@tcp(localhost:3306)/finance" -verbose down

.PHONY: migrate migratedown createmigration