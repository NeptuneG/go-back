.PHONY: svc-build-image
svc-build-image:
	docker build -t neptuneg/go-back/$(svc)-service:latest --file ./services/$(svc)/Dockerfile .

.PHONY: svc-generate-migrate
svc-generate-migrate:
	docker exec $(svc)-service migrate create -dir db/migrations -ext sql $(NAME)

.PHONY: svc-db-create
svc-db-create:
	docker exec $(svc)-db createdb --username=dev --owner=dev $(svc)_development
	docker exec $(svc)-db createdb --username=dev --owner=dev $(svc)_test

.PHONY: svc-db-drop
svc-db-drop:
	docker exec $(svc)-db dropdb --username=dev -f $(svc)_development
	docker exec $(svc)-db dropdb --username=dev -f $(svc)_test

.PHONY: svc-db-migrate
svc-db-migrate:
	docker exec -it $(svc)-service migrate \
	-database postgresql://dev@$(svc)-db/$(svc)_development?sslmode=disable \
	-path db/migrations \
	-verbose up

.PHONY: svc-db-rollback
svc-db-rollback:
	docker exec -it $(svc)-service migrate \
	-database postgresql://dev@$(svc)-db/$(svc)_development?sslmode=disable \
	-path db/migrations \
	-verbose down $(or $(STEP), 1)

.PHONY: svc-db-seed
svc-db-seed:
	cat services/$(svc)/db/seeds.sql | xargs -0 docker exec $(svc)-db psql -U dev -d $(svc)_development -c

.PHONY: svc-sqlc-generate
svc-sqlc-generate:
	cd services/$(svc) && sqlc generate

.PHONY: db-create-all
db-create-all:
	make svc-db-create svc=user
	make svc-db-create svc=live
	make svc-db-create svc=payment

.PHONY: db-drop-all
db-drop-all:
	make svc-db-drop svc=user
	make svc-db-drop svc=live
	make svc-db-drop svc=payment

.PHONY: db-migrate-all
db-migrate-all:
	make svc-db-migrate svc=user
	make svc-db-migrate svc=live
	make svc-db-migrate svc=payment

.PHONY: db-seed-all
db-seed-all:
	make svc-db-migrate svc=live

.PHONY: build-images-all
build-images-all:
	make svc-build-image svc=user
	make svc-build-image svc=live
	make svc-build-image svc=gateway
	make svc-build-image svc=scraper
	make svc-build-image svc=payment

