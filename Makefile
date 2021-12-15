.PHONY: build-gateway-svc
build-gateway-svc:
	docker build -t neptuneg/go-back/gateway-service:latest --file ./services/gateway/Dockerfile .

.PHONY: build-user-svc
build-user-svc:
	docker build -t neptuneg/go-back/user-service:latest --file ./services/user/Dockerfile .

.PHONY: user-generate-migrate
user-generate-migrate:
	docker exec user-service migrate create -dir db/migrations -ext sql $(NAME)

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

.PHONY: build-live-svc
build-live-svc:
	docker build -t neptuneg/go-back/live-service:latest --file ./services/live/Dockerfile .

.PHONY: live-generate-migrate
live-generate-migrate:
	docker exec live-service migrate create -dir db/migrations -ext sql $(NAME)

.PHONY: live-db-create
live-db-create:
	docker exec live-db createdb --username=dev --owner=dev live_development
	docker exec live-db createdb --username=dev --owner=dev live_test

.PHONY: live-db-drop
live-db-drop:
	docker exec live-db dropdb --username=dev -f live_development
	docker exec live-db dropdb --username=dev -f live_test

.PHONY: live-db-seed
live-db-seed:
	cat services/live/db/seeds.sql | xargs -0 docker exec live-db psql -U dev -d live_development -c

.PHONY: live-db-migrate
live-db-migrate:
	docker exec -it live-service migrate \
	-database postgresql://dev@live-db/live_development?sslmode=disable \
	-path db/migrations \
	-verbose up

.PHONY: live-db-rollback
live-db-rollback:
	docker exec -it live-service migrate \
	-database postgresql://dev@live-db/live_development?sslmode=disable \
	-path db/migrations \
	-verbose down $(or $(STEP), 1)

.PHONY: live-sqlc-generate
live-sqlc-generate:
	cd services/live && sqlc generate

.PHONY: build-scraper-svc
build-scraper-svc:
	docker build -t neptuneg/go-back/scraper-service:latest --file ./services/scraper/Dockerfile .

.PHONY: db-create
db-create:
	docker exec user-db createdb --username=dev --owner=dev user_development
	docker exec user-db createdb --username=dev --owner=dev user_test
	docker exec live-db createdb --username=dev --owner=dev live_development
	docker exec live-db createdb --username=dev --owner=dev live_test

.PHONY: db-migrate
db-migrate:
	docker exec -it user-service migrate \
	-database postgresql://dev@user-db/user_development?sslmode=disable \
	-path db/migrations \
	-verbose up
	docker exec -it live-service migrate \
	-database postgresql://dev@live-db/live_development?sslmode=disable \
	-path db/migrations \
	-verbose up

.PHONY: db-seed
db-seed:
	cat services/live/db/seeds.sql | xargs -0 docker exec live-db psql -U dev -d live_development -c

.PHONY: build-services
build-services:
	docker build -t neptuneg/go-back/gateway-service:latest --file ./services/gateway/Dockerfile .
	docker build -t neptuneg/go-back/user-service:latest --file ./services/user/Dockerfile .
	docker build -t neptuneg/go-back/live-service:latest --file ./services/live/Dockerfile .
	docker build -t neptuneg/go-back/scraper-service:latest --file ./services/scraper/Dockerfile .
