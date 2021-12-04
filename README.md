# go-back

## Descriptions

- go: 1.17.3
- postgres: 14.1
- setup
    ```bash
    docker-compose build
    docker-compose up -d
    make create-db
    make db-migrate
    ```
- generate db migration
    ```bash
    make generate-migrate NAME=#{migration_name}
    ```
- db migrate
    ```bash
    make db-migrate
    ```
- db rollback
    ```bash
    make db-rollback STEP=n # default STEP is 1
    ```

## Steup steps

- `docker-compose run --rm app go mod init github.com/NeptuneG/go-back`
- `docker-compose run --rm app air init`
- `sqlc init`
    - sqlc installed on localhost
    - [Installing sqlc — sqlc 1.11.0 documentation](https://docs.sqlc.dev/en/latest/overview/install.html)
    - [Configuration file (version 1) — sqlc 1.10.0 documentation](https://docs.sqlc.dev/en/stable/reference/config.html)

## DB Diagram

- [go-back - dbdiagram.io](https://dbdiagram.io/d/619f9ec18c901501c0d2b534)

## Todo

- [ ] auto `updated_at`
    - [PostgreSQL アップデート時のタイムスタンプ自動更新 - IT技術で仕事を減らしたい！](https://timesaving.hatenablog.com/entry/2020/08/29/210000)
- [ ] slug
- [ ] test
- [ ] logging
- [ ] k8s

## References

### Fundamentals
- [Backend master class Series' Articles - DEV Community 👩‍💻👨‍💻](https://dev.to/techschoolguru/series/7172)
    - [dbdiagram.io - Database Relationship Diagrams Design Tool](https://dbdiagram.io/home)
- [Develop a Go app with Docker Compose | FireHydrant](https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/)
- [My Ultimate Makefile for Golang Projects | by Thomas Poignant | Better Programming](https://betterprogramming.pub/my-ultimate-makefile-for-golang-projects-fcc8ca20c9bb)

### gin
- [ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)
