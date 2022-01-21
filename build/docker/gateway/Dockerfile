FROM golang:1.17.3-alpine3.14 as dev

ENV LANG="ja_JP.UTF-8"
RUN set -x && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.14/main' >> /etc/apk/repositories && \
    apk add --update --no-cache tzdata curl build-base && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /go/src/github.com/NeptuneG/go-back
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .

WORKDIR /go/src/github.com/NeptuneG/go-back/cmd/gateway

FROM golang:1.17.3-alpine3.14 as builder

ENV LANG="ja_JP.UTF-8"
RUN set -x && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.14/main' >> /etc/apk/repositories && \
    apk add --update --no-cache tzdata curl build-base && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /go/src/github.com/NeptuneG/go-back

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o ./bin/gateway ./cmd/gateway

FROM alpine:3.14
RUN apk --no-cache add ca-certificates && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /go/src/github.com/NeptuneG/go-back/bin/gateway .
ENTRYPOINT ["./gateway"]
