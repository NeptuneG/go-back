.PHONY: generate-migrate
generate-migrate:
	docker exec user-service migrate create -dir /app/db/migrations -ext sql $(NAME)

.PHONY: db-create
db-create:
	docker exec user-db createdb --username=dev --owner=dev user_development
	docker exec user-db createdb --username=dev --owner=dev user_test

.PHONY: db-drop
db-drop:
	docker exec user-db dropdb --username=dev -f user_development
	docker exec user-db dropdb --username=dev -f user_test

.PHONY: db-migrate
db-migrate:
	docker exec -it user-service migrate \
	-database postgresql://dev@user-db/user_development?sslmode=disable \
	-path /app/db/migrations \
	-verbose up

.PHONY: db-rollback
db-rollback:
	docker exec -it user-service migrate \
	-database postgresql://dev@user-db/user_development?sslmode=disable \
	-path /app/db/migrations \
	-verbose down $(or $(STEP), 1)

.PHONY: db-seed
db-seed:
	cat app/db/seeds.sql | xargs -0 docker exec user-db psql -U dev -d user_development -c

.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate

.PHONY: go-mod-tidy
go-mod-tidy:
	docker exec user-service go mod tidy -compat=1.17
