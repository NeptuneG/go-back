.PHONY: build-gateway-svc
build-gateway-svc:
	docker build -t neptuneg/go-back/gateway-service:latest --file ./services/gateway/Dockerfile .

.PHONY: build-user-svc
build-user-svc:
	docker build -t neptuneg/go-back/user-service:latest --file ./services/user/Dockerfile .

.PHONY: generate-migrate
generate-migrate:
	docker exec go-back-app migrate create -dir db/migrations -ext sql $(NAME)

.PHONY: user-db-create
user-db-create:
	docker exec user-db createdb --username=dev --owner=dev user_development
	docker exec user-db createdb --username=dev --owner=dev user_test

.PHONY: user-db-drop
user-db-drop:
	docker exec user-db dropdb --username=dev -f user_development
	docker exec user-db dropdb --username=dev -f user_test

.PHONY: user-db-migrate
user-db-migrate:
	docker exec -it user-service migrate \
	-database postgresql://dev@user-db/user_development?sslmode=disable \
	-path db/migrations \
	-verbose up

.PHONY: user-db-rollback
user-db-rollback:
	docker exec -it user-service migrate \
	-database postgresql://dev@user-db/user_development?sslmode=disable \
	-path db/migrations \
	-verbose down $(or $(STEP), 1)

.PHONY: user-sqlc-generate
user-sqlc-generate:
	cd services/user && sqlc generate
