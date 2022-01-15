.PHONY: generate-svc
generate-svc:
	./.scaffold/script/scaffold.sh $(svc_name)

.PHONY: svc-build-image
svc-build-image:
	docker build -t neptuneg/$(svc)-service:latest --file ./build/docker/$(svc)/Dockerfile .

.PHONY: svc-push-image
svc-push-image:
	docker push neptuneg/$(svc)-service:latest

.PHONY: svc-generate-migrate
svc-generate-migrate:
	docker exec $(svc)-service migrate create -dir migrations/$(svc) -ext sql $(NAME)

.PHONY: svc-db-migrate
svc-db-migrate:
	docker exec -it $(svc)-service migrate \
	-database postgresql://dev@db/$(svc)_development?sslmode=disable \
	-path ../../migrations/$(svc) \
	-verbose up

.PHONY: svc-db-rollback
svc-db-rollback:
	docker exec -it $(svc)-service migrate \
	-database postgresql://dev@db/$(svc)_development?sslmode=disable \
	-path ../../migrations/$(svc) \
	-verbose down $(or $(STEP), 1)

.PHONY: svc-db-seed
svc-db-seed:
	cat seeds/$(svc)/seeds.sql | xargs -0 docker exec go-back-db psql -U dev -d $(svc)_development -c

.PHONY: svc-sqlc-generate
svc-sqlc-generate:
	cd internal/$(svc) && sqlc generate

.PHONY: db-create-all
db-create-all:
	docker exec go-back-db createdb --username=dev --owner=dev live_development
	docker exec go-back-db createdb --username=dev --owner=dev live_test
	docker exec go-back-db createdb --username=dev --owner=dev payment_development
	docker exec go-back-db createdb --username=dev --owner=dev payment_test
	docker exec go-back-db createdb --username=dev --owner=dev auth_development
	docker exec go-back-db createdb --username=dev --owner=dev auth_test

.PHONY: db-drop-all
db-drop-all:
	docker exec go-back-db dropdb --username=dev -f live_development
	docker exec go-back-db dropdb --username=dev -f live_test
	docker exec go-back-db dropdb --username=dev -f payment_development
	docker exec go-back-db dropdb --username=dev -f payment_test
	docker exec go-back-db dropdb --username=dev -f auth_development
	docker exec go-back-db dropdb --username=dev -f auth_test

.PHONY: db-migrate-all
db-migrate-all:
	make svc-db-migrate svc=live
	make svc-db-migrate svc=payment
	make svc-db-migrate svc=auth

.PHONY: db-seed-all
db-seed-all:
	make svc-db-seed svc=live

.PHONY: sqlc-generate-all
sqlc-generate-all:
	make svc-sqlc-generate svc=auth
	make svc-sqlc-generate svc=live
	make svc-sqlc-generate svc=payment

.PHONY: build-images-all
build-images-all:
	make svc-build-image svc=auth
	make svc-build-image svc=live
	make svc-build-image svc=gateway
	make svc-build-image svc=scraper
	make svc-build-image svc=payment
	docker build -t neptuneg/faktory-workers:latest --file ./build/docker/faktory-workers/Dockerfile ./cmd/faktory-workers

.PHONY: push-images-all
push-images-all:
	make svc-push-image svc=auth
	make svc-push-image svc=live
	make svc-push-image svc=gateway
	make svc-push-image svc=scraper
	make svc-push-image svc=payment
	docker push neptuneg/faktory-workers:latest

