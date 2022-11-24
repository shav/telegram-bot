OS=$(shell uname -s)
OSARCH=$(shell uname -m)
USRBIN=/usr/local/bin
CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
DBMIGRATE=${BINDIR}/goose_${GOVER}
DATA_ACCESS=777
BUF=${USRBIN}/buf
BUFVER=1.9.0
BUFSRC="http://github.com/bufbuild/buf/releases/download/v${BUFVER}/buf-${OS}-${OSARCH}"

# TODO: Получать настройки подключения к БД из конфига или docker-compose файла
POSTGRES_HOST=localhost
POSTGRES_DB=telegram_bot
POSTGRES_DB_PORT=5432
POSTGRES_TEST_DB=telegram_bot_test
POSTGRES_TEST_DB_PORT=5433
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_CONN_STRING="host=${POSTGRES_HOST} port=${POSTGRES_DB_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable"
POSTGRES_TEST_CONN_STRING="host=${POSTGRES_HOST} port=${POSTGRES_TEST_DB_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_TEST_DB} sslmode=disable"
TEST_CACHE_CONN_STRING="localhost:6380"

PACKAGE=github.com/shav/telegram-bot/cmd/bot
PACKAGE_REPORT=github.com/shav/telegram-bot/cmd/report

# **********************************************************************************************************************
# Разработка (сборка, тесты, линтер)
# **********************************************************************************************************************

all: format build test lint

build: bindir
	go build -o ${BINDIR}/bot ${PACKAGE}
	go build -o ${BINDIR}/report ${PACKAGE_REPORT}

test: docker-run-test-infrastructure test-database-up
	sudo env PATH="$$PATH" TELEGRAM_BOT_DB_CONNECTION_STRING=${POSTGRES_TEST_CONN_STRING} TELEGRAM_BOT_CACHE_CONNECTION_STRING=${TEST_CACHE_CONN_STRING} go test -count=1 ${CURDIR}/...

test-ci:
	go test -count=1 ./...

run:
	go run ${PACKAGE}

run-report:
	go run ${PACKAGE_REPORT}

generate:
	#sudo env PATH="$$PATH" go generate ${CURDIR}/...
	cd ${CURDIR}/internal/modules/core/transport/grpc && sudo env PATH="$$PATH" buf generate

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/**/mocks

# **********************************************************************************************************************
#  Установка сторонних зависимостей
# **********************************************************************************************************************

install-gen: install-mock-gen install-json-gen install-grpc-gen

install-mock-gen:
	go get github.com/gojuno/minimock/v3/cmd/minimock
	go install github.com/gojuno/minimock/v3/cmd/minimock

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

install-dbmigrate: bindir
	test -f ${DBMIGRATE} || \
		(GOBIN=${BINDIR} go install github.com/pressly/goose/v3/cmd/goose@latest && \
		mv ${BINDIR}/goose ${DBMIGRATE})

install-psql:
	sudo apt-get install postgresql-client-common
	sudo apt-get install postgresql-client

install-json-gen:
	go get github.com/mailru/easyjson
	go install github.com/mailru/easyjson/...@latest

install-grpc-gen:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	sudo test -f ${BUF} || (sudo curl ${BUFSRC} -o ${BUF} && sudo chmod +x ${BUF})

# **********************************************************************************************************************
# База данных
# **********************************************************************************************************************

database-up: install-psql install-dbmigrate
	PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${POSTGRES_HOST} -p ${POSTGRES_DB_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -a -f "${CURDIR}/db/reset_db.sql"
	${DBMIGRATE} -s -dir "${CURDIR}/db/migrations" postgres ${POSTGRES_CONN_STRING} up

database-convert: install-dbmigrate
	${DBMIGRATE} -s -dir "${CURDIR}/db/migrations" postgres ${POSTGRES_CONN_STRING} up

test-database-up: install-psql install-dbmigrate
	PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${POSTGRES_HOST} -p ${POSTGRES_TEST_DB_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_TEST_DB} -a -f "${CURDIR}/db/reset_db.sql"
	${DBMIGRATE} -s -dir "${CURDIR}/db/migrations" postgres ${POSTGRES_TEST_CONN_STRING} up

test-database-convert: install-dbmigrate
	${DBMIGRATE} -s -dir "${CURDIR}/db/migrations" postgres ${POSTGRES_TEST_CONN_STRING} up

# **********************************************************************************************************************
# Деплой приложения
# **********************************************************************************************************************

docker-build: build
	yes | cp -f /etc/ssl/certs/ca-certificates.crt ${BINDIR}
	sudo docker build -t telegram-bot-finance:1.0 -f ${CURDIR}/docker/bot/Dockerfile ${BINDIR}
	sudo docker build -t telegram-bot-report:1.0 -f ${CURDIR}/docker/report/Dockerfile ${BINDIR}

docker-run: bot report

docker-run-all: docker-network docker-run-infrastructure docker-run

docker-stop: bot-stop report-stop

docker-delete: bot-down report-down

docker-delete-images:
	docker image rm telegram-bot-finance:1.0
	docker image rm telegram-bot-report:1.0

docker-delete-all: docker-delete docker-delete-images

.PHONY: bot
bot:
	cd ${CURDIR}/docker/bot && sudo docker compose up -d

.PHONY: bot-stop
bot-stop:
	cd ${CURDIR}/docker/bot && sudo docker compose stop

.PHONY: bot-down
bot-down:
	cd ${CURDIR}/docker/bot && sudo docker compose down

.PHONY: report
report:
	cd ${CURDIR}/docker/report && sudo docker compose up -d

.PHONY: report-stop
report-stop:
	cd ${CURDIR}/docker/report && sudo docker compose stop

.PHONY: report-down
report-down:
	cd ${CURDIR}/docker/report && sudo docker compose down

# **********************************************************************************************************************
# Деплой инфраструктуры
# **********************************************************************************************************************

docker-network:
	sudo docker network create tgbot-network || true

docker-network-down:
	sudo docker network rm tgbot-network || true

docker-pull:
	sudo docker pull postgres:12.2-alpine
	sudo docker pull redis:7.0
	sudo docker pull jaegertracing/all-in-one:1.18
	sudo docker pull prom/prometheus
	sudo docker pull grafana/grafana-oss
	sudo docker pull wurstmeister/kafka
	sudo docker pull wurstmeister/zookeeper

docker-run-infrastructure: db cache tracing metrics message-queue

docker-stop-infrastructure: db-stop cache-stop tracing-stop metrics-stop message-queue-stop

docker-delete-infrastructure: db-down cache-down tracing-down metrics-down message-queue-down docker-network-down

docker-run-test-infrastructure: db-test cache-test

docker-stop-test-infrastructure: db-test-stop cache-test-stop

docker-delete-test-infrastructure: db-test-down cache-test-down

.PHONY: db
db:
	cd ${CURDIR}/docker/db/dev && \
 	mkdir -p data && sudo chmod -R ${DATA_ACCESS} data && \
 	sudo docker compose up -d

.PHONY: db-stop
db-stop:
	cd ${CURDIR}/docker/db/dev && sudo docker compose stop

.PHONY: db-down
db-down:
	cd ${CURDIR}/docker/db/dev && sudo docker compose down

.PHONY: db-test
db-test:
	cd ${CURDIR}/docker/db/test && \
 	mkdir -p data && sudo chmod -R ${DATA_ACCESS} data && \
 	sudo docker compose up -d

.PHONY: db-test-stop
db-test-stop:
	cd ${CURDIR}/docker/db/test && sudo docker compose stop

.PHONY: db-test-down
db-test-down:
	cd ${CURDIR}/docker/db/test && sudo docker compose down

.PHONY: cache
cache:
	cd ${CURDIR}/docker/cache/dev && \
	mkdir -p data && sudo chmod -R ${DATA_ACCESS} data && \
	sudo docker compose up -d

.PHONY: cache-stop
cache-stop:
	cd ${CURDIR}/docker/cache/dev && sudo docker compose stop

.PHONY: cache-down
cache-down:
	cd ${CURDIR}/docker/cache/dev && sudo docker compose down

.PHONY: cache-test
cache-test:
	cd ${CURDIR}/docker/cache/test && \
	mkdir -p data && sudo chmod -R ${DATA_ACCESS} data && \
	sudo docker compose up -d

.PHONY: cache-test-stop
cache-test-stop:
	cd ${CURDIR}/docker/cache/test && sudo docker compose stop

.PHONY: cache-test-down
cache-test-down:
	cd ${CURDIR}/docker/cache/test && sudo docker compose down

.PHONY: tracing
tracing:
	cd ${CURDIR}/docker/tracing && sudo docker compose up -d

.PHONY: tracing-stop
tracing-stop:
	cd ${CURDIR}/docker/tracing && sudo docker compose stop

.PHONY: tracing-down
tracing-down:
	cd ${CURDIR}/docker/tracing && sudo docker compose down

.PHONY: metrics
metrics:
	cd ${CURDIR}/docker/metrics && \
 	mkdir -p data && sudo chmod -R ${DATA_ACCESS} data && \
 	sudo docker compose up -d

.PHONY: metrics-stop
metrics-stop:
	cd ${CURDIR}/docker/metrics && sudo docker compose stop

.PHONY: metrics-down
metrics-down:
	cd ${CURDIR}/docker/metrics && sudo docker compose down

.PHONY: message-queue
message-queue:
	cd ${CURDIR}/docker/message_queue && \
	mkdir -p kafka-data && sudo chmod -R ${DATA_ACCESS} kafka-data && \
	mkdir -p zk-data && sudo chmod -R ${DATA_ACCESS} zk-data && \
	mkdir -p zk-tx-logs && sudo chmod -R ${DATA_ACCESS} zk-tx-logs && \
 	sudo docker compose up -d

.PHONY: message-queue-stop
message-queue-stop:
	cd ${CURDIR}/docker/message_queue && sudo docker compose stop

.PHONY: message-queue-down
message-queue-down:
	cd ${CURDIR}/docker/message_queue && sudo docker compose down
