db/start:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DATABASE=account -d postgres:latest

db/init:
	docker exec -it postgres psql -U admin -c "CREATE DATABASE account;"

db/migrate:
	migrate -path migrations -database "postgresql://admin:admin@localhost:5432/account?sslmode=disable" -verbose up

db/rollback:
	migrate -path migrations -database "postgresql://admin:admin@localhost:5432/account?sslmode=disable" -verbose down