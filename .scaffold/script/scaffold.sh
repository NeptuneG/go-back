#!/bin/sh

set -e

if [ ! -z $1 ]
then
    export SERVICE_UNDERSCORE_NAME=$1
    export SERVICE_PASCALCASE_NAME=$(echo $SERVICE_UNDERSCORE_NAME | perl -pe 's/(^|_)./uc($&)/ge;s/_//g')

    mkdir -p api/proto/$SERVICE_UNDERSCORE_NAME/proto
    mkdir -p migrations/$SERVICE_UNDERSCORE_NAME
    mkdir -p build/docker/$SERVICE_UNDERSCORE_NAME
    mkdir -p services/$SERVICE_UNDERSCORE_NAME
    mkdir -p services/$SERVICE_UNDERSCORE_NAME/db/query
    mkdir -p services/$SERVICE_UNDERSCORE_NAME/db/sqlc
    mkdir -p services/$SERVICE_UNDERSCORE_NAME/server

    cp -r .scaffold/template/migrations/* migrations/$SERVICE_UNDERSCORE_NAME
    cp -r .scaffold/template/service/db/* services/$SERVICE_UNDERSCORE_NAME/db
    cp .scaffold/template/service/.air.toml services/$SERVICE_UNDERSCORE_NAME/.air.toml
    cp .scaffold/template/service/sqlc.yaml services/$SERVICE_UNDERSCORE_NAME/sqlc.yaml
    envsubst '\${SERVICE_UNDERSCORE_NAME} \${SERVICE_PASCALCASE_NAME}' < .scaffold/template/service/proto/service.proto.template  > api/proto/$SERVICE_UNDERSCORE_NAME/proto/$SERVICE_UNDERSCORE_NAME.proto
    envsubst '\${SERVICE_UNDERSCORE_NAME} \${SERVICE_PASCALCASE_NAME}' < .scaffold/template/service/server/server.go.template > services/$SERVICE_UNDERSCORE_NAME/server/server.go
    envsubst '\${SERVICE_UNDERSCORE_NAME} \${SERVICE_PASCALCASE_NAME}' < .scaffold/template/service/Dockerfile.template > build/docker/$SERVICE_UNDERSCORE_NAME/Dockerfile
    envsubst '\${SERVICE_UNDERSCORE_NAME} \${SERVICE_PASCALCASE_NAME}' < .scaffold/template/service/main.go.template > services/$SERVICE_UNDERSCORE_NAME/main.go

    buf generate

    # TODO: sed or awk to update Makefile and docker-compose.yml
    echo "Basic files for ${SERVICE_UNDERSCORE_NAME} are generated. Please add or update the followings:"
    echo "  - docker-compose.yml & .env"
    echo "  - Makefile"
    echo "  - db migrations then make svc-db-migrate svc=${SERVICE_UNDERSCORE_NAME}"
else
    echo "Usage: $0 <service_name> (must in underscore_case)"
    exit 1
fi
