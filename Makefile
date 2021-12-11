.PHONY: generate-migrate
generate-migrate:
	docker exec go-back-app migrate create -dir /app/db/migrations -ext sql $(NAME)

.PHONY: db-create
db-create:
	docker exec go-back-db createdb --username=dev --owner=dev back_development
	docker exec go-back-db createdb --username=dev --owner=dev back_test

.PHONY: db-drop
db-drop:
	docker exec go-back-db dropdb --username=dev -f back_development
	docker exec go-back-db dropdb --username=dev -f back_test

.PHONY: db-migrate
db-migrate:
	docker exec -it go-back-app migrate \
	-database postgresql://dev@db/back_development?sslmode=disable \
	-path /app/db/migrations \
	-verbose up

.PHONY: db-rollback
db-rollback:
	docker exec -it go-back-app migrate \
	-database postgresql://dev@db/back_development?sslmode=disable \
	-path /app/db/migrations \
	-verbose down $(or $(STEP), 1)

.PHONY: db-seed
db-seed:
	cat app/db/seeds.sql | xargs -0 docker exec go-back-db psql -U dev -d back_development -c

.PHONY: sqlc-generate
sqlc-generate:
	cd app && sqlc generate

.PHONY: go-mod-tidy
go-mod-tidy:
	docker exec go-back-app go mod tidy -compat=1.17
