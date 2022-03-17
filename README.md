# go-back

## Descriptions

- Following [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- Inspired by
    - [マイクロサービスにおける決済トランザクション管理 | メルカリエンジニアリング](https://engineering.mercari.com/blog/entry/2019-06-07-155849/) | [Payment Transaction Management in Microservices | Mercari Engineering](https://engineering.mercari.com/en/blog/entry/20210831-2019-06-07-155849/)
    - [mercari/mercari-microservices-example](https://github.com/mercari/mercari-microservices-example)

## Microservices

### gateway
- gRPC -> JSON proxy service by [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
### dtm
- distributed transaction management service by [dtm-labs/dtf](https://github.com/dtm-labs/dtf)
### auth
- JWT authentication service
### live
- live houses and live events service
### payment
- payment (reserving live event seat) service
### scraper
- scrapping requests handling service
- enqueues scrapping jobs
- consume scrapped live events from message queue
### faktory
- job queue service by [contribsys/faktory](https://github.com/contribsys/faktory)
### faktory-workers
- live events (of [ビルボードライブ東京](http://www.billboard-live.com/pg/shop/index.php?mode=top&shop=1)) scrapping worker
- products scrapped live events as messages in message queue

## Operations

### Helm chart
- [NeptuneG/go-back-manifests](https://github.com/NeptuneG/go-back-manifests)
### Service mesh
- Istio
### CI/CD
- GitHub Actions
- ArgoCD
<!--
## Descriptions
- go: 1.17.3
- postgres: 14.1
- requires sqlc
    - [Installing sqlc — sqlc 1.11.0 documentation](https://docs.sqlc.dev/en/latest/overview/install.html)
    - [Configuration file (version 1) — sqlc 1.10.0 documentation](https://docs.sqlc.dev/en/stable/reference/config.html)
- setup
    ```bash
    make build-images-all
    docker-compose build
    docker-compose up -d
    make db-create-all
    make db-migrate-all
    ```
- generate db migration
    ```bash
    make svc-generate-migrate svc=#{service_name} NAME=#{snake_case_migration_name}
    ```
- db migrate
    ```bash
    make svc-db-migrate svc=#{service_name}
    ```
- db rollback
    ```bash
    make svc-db-rollback svc=#{service_name} STEP=n # default STEP is 1
    ```
- generate codebase for a new service
    ```bash
    make generate-svc svc_name=#{service_name} # service_name must be snake_case
    ```

## DB Diagram

- [go-back - dbdiagram.io](https://dbdiagram.io/d/619f9ec18c901501c0d2b534)

## Todo

- [ ] auto `updated_at`
    - [PostgreSQL アップデート時のタイムスタンプ自動更新 - IT技術で仕事を減らしたい！](https://timesaving.hatenablog.com/entry/2020/08/29/210000)
- [ ] slug
- [ ] test
    - [GoConvey - Go testing in the browser](http://goconvey.co/)
    - [fortio/fortio: Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats.](https://github.com/fortio/fortio)
- [x] logging
    - [uber-go/zap: Blazing fast, structured, leveled logging in Go.](https://github.com/uber-go/zap)
- [x] k8s
    - https://github.com/NeptuneG/go-back-manifests
- [ ] pagination
- [ ] service discovery
- [ ] config center
- [X] health check
- [x] cache
- [ ] rate limit
- [ ] hot reload debug
    - air does not work well due to 2345 may not be avaiable
- [x] timeout & retry
- [ ] grpc responses
- [ ] TLS

## Notes

- `docker-compose run --rm app go mod init github.com/NeptuneG/go-back`
- `docker-compose run --rm app air init`
- redis pubsub
    - subcriber will not consume payload automatically after restart
    - multiple subcriber will consume payload repeatedly
    - cannot persist data
    - no ack
- redis list LPUSH & BRPop
    - able to persist data
    - no ack?
- `protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:proto`
    - [bufbuild/buf: A new way of working with Protocol Buffers.](https://github.com/bufbuild/buf)
- context deadline exceeded
    - [Why am I seeing `context deadline exceeded` errors – HashiCorp Help Center](https://support.hashicorp.com/hc/en-us/articles/4404634420755-Why-am-I-seeing-context-deadline-exceeded-errors)

## 🤯🤯🤯

- create a live event elegantly against available seats

## References

### Fundamentals
- [Backend master class Series' Articles - DEV Community 👩‍💻👨‍💻](https://dev.to/techschoolguru/series/7172)
    - [dbdiagram.io - Database Relationship Diagrams Design Tool](https://dbdiagram.io/home)
- [Develop a Go app with Docker Compose | FireHydrant](https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/)
- [My Ultimate Makefile for Golang Projects | by Thomas Poignant | Better Programming](https://betterprogramming.pub/my-ultimate-makefile-for-golang-projects-fcc8ca20c9bb)

### gin
- [ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)

### Debug
- [Setup Go with VSCode in Docker and Air for debugging - DEV Community 👩‍💻👨‍💻](https://dev.to/andreidascalu/setup-go-with-vscode-in-docker-for-debugging-24ch)

### grpc
- [gRPC-go 入门（1）：Hello World - 知乎](https://zhuanlan.zhihu.com/p/258879142)

### logging
- [一文告诉你如何用好uber开源的zap日志库 | Tony Bai](https://tonybai.com/2021/07/14/uber-zap-advanced-usage/#comment-7590)
-->
