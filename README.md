# go-back

## Descriptions

- go: 1.17.3
- postgres: 14.1
- requires sqlc
    - [Installing sqlc — sqlc 1.11.0 documentation](https://docs.sqlc.dev/en/latest/overview/install.html)
    - [Configuration file (version 1) — sqlc 1.10.0 documentation](https://docs.sqlc.dev/en/stable/reference/config.html)
- setup
    ```bash
    make build-services
    docker-compose build
    docker-compose up -d
    make db-create
    make db-migrate
    ```
- generate db migration
    ```bash
    make generate-migrate NAME=#{snake_case_migration_name}
    ```
- db migrate
    ```bash
    make db-migrate
    ```
- db rollback
    ```bash
    make db-rollback STEP=n # default STEP is 1
    ```

## DB Diagram

- [go-back - dbdiagram.io](https://dbdiagram.io/d/619f9ec18c901501c0d2b534)

## Todo

- [ ] auto `updated_at`
    - [PostgreSQL アップデート時のタイムスタンプ自動更新 - IT技術で仕事を減らしたい！](https://timesaving.hatenablog.com/entry/2020/08/29/210000)
- [ ] slug
- [ ] test
- [ ] logging
- [ ] k8s
- [ ] pagination
- [ ] service discovery
- [ ] config center
- [ ] health check
- [ ] hot reaload debug
    - air does not work well due to 2345 is occupied

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

## 🤯🤯🤯

- create order elegantly against available seats

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
