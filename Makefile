SOURCES := $(shell find . -name '*.go' -type f -not -path './vendor/*' -not -path '*/mocks/*' -not -path '*/migrations/*')

MYSQL_DB_USER ?= marketplace
MYSQL_DB_PASSWORD ?= secret
MYSQL_DB_PORT ?= 3306
MYSQL_DB_ADDRESS ?= 127.0.0.1:${MYSQL_DB_PORT}
MYSQL_DB_NAME ?= marketplace_test
MYSQL_DB_CONTAINER_NAME ?= marketplace-mysql

MONGO_DB_PORT ?= 27017
MONGO_DB_ADDRESS ?= 127.0.0.1:${MONGO_DB_PORT}
MONGO_DB_NAME ?= marketplace_test
MONGO_DB_CONTAINER_NAME ?= marketplace-mongo

ELASTICSEARCH_NETWORK_NAME ?= elasticsearch
ELASTICSEARCH_CONTAINER_NAME ?= marketplace-elasticsearch
ELASTICSEARCH_KIBANA_CONTAINER_NAME ?= marketplace-kibana

# Linter
.PHONY: prepare-lint
prepare-lint:
	@echo Install linter
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1

.PHONY: lint
lint:
	@echo Run linter
	@golangci-lint run $(LINT_OPTS)

# Dependency management
.PHONY: dep-upgrade
dep-upgrade:
	@echo Upgrading dependencies
	@go get -u all

.PHONY: dep-tidy
dep-tidy:
	@echo Tidying dependencies
	@go mod tidy

# Generate files
.PHONY: generate
generate:
	@echo Generating files
	@go generate ./...

# MySQL
.PHONY: docker-mysql-up
docker-mysql-up:
	@docker run --rm -d --name $(MYSQL_DB_CONTAINER_NAME) -p ${MYSQL_DB_PORT}:3306 -e MYSQL_DATABASE=$(MYSQL_DB_NAME) -e MYSQL_USER=$(MYSQL_DB_USER) -e MYSQL_PASSWORD=$(MYSQL_DB_PASSWORD) -e MYSQL_ROOT_PASSWORD=rootsecret mysql:5.7.27 && docker logs -f $(MYSQL_DB_CONTAINER_NAME)

.PHONY: docker-mysql-down
docker-mysql-down:
	@docker stop $(MYSQL_DB_CONTAINER_NAME)

# Mongo
.PHONY: docker-mongo-up
docker-mongo-up:
	@docker run --rm -d --name $(MONGO_DB_CONTAINER_NAME) -p ${MONGO_DB_PORT}:27017 -e MONGO_INIT_ROOT_USERNAME=root -e MONGO_INIT_ROOT_PASSWORD=rootsecret mongo:4.2.0 && docker logs -f $(MONGO_DB_CONTAINER_NAME)

.PHONY: docker-mongo-down
docker-mongo-down:
	@docker stop $(MONGO_DB_CONTAINER_NAME)

# Elasticsearch
.PHONY: docker-elasticsearch-up
docker-elasticsearch-up:
	@docker network create $(ELASTICSEARCH_NETWORK_NAME) || true
	@docker run --rm -d --name $(ELASTICSEARCH_CONTAINER_NAME) --net $(ELASTICSEARCH_NETWORK_NAME) -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.3.1 && docker logs -f $(ELASTICSEARCH_CONTAINER_NAME)

.PHONY: docker-elasticsearch-down
docker-elasticsearch-down:
	@docker stop $(ELASTICSEARCH_CONTAINER_NAME) || true
	@docker network rm $(ELASTICSEARCH_NETWORK_NAME)

# Kibana
.PHONY: docker-kibana-up
docker-kibana-up:
	@docker run --rm -d --name $(ELASTICSEARCH_KIBANA_CONTAINER_NAME) --net $(ELASTICSEARCH_NETWORK_NAME) --link $(ELASTICSEARCH_CONTAINER_NAME):elasticsearch -p 5601:5601 docker.elastic.co/kibana/kibana:7.3.1 && docker logs -f $(ELASTICSEARCH_KIBANA_CONTAINER_NAME)

.PHONY: docker-kibana-down
docker-kibana-down:
	@docker stop $(ELASTICSEARCH_KIBANA_CONTAINER_NAME)

# MySQL migrations
.PHONY: mysql-migrate-up
mysql-migrate-up:
	@migrate -database "mysql://$(MYSQL_DB_USER):$(MYSQL_DB_PASSWORD)@tcp($(MYSQL_DB_ADDRESS))/$(MYSQL_DB_NAME)?multiStatements=true" -path=internal/mysql/migrations up

.PHONY: mysql-migrate-down
mysql-migrate-down:
	@migrate -database "mysql://$(MYSQL_DB_USER):$(MYSQL_DB_PASSWORD)@tcp($(MYSQL_DB_ADDRESS))/$(MYSQL_DB_NAME)?multiStatements=true" -path=internal/mysql/migrations down

.PHONY: mysql-migrate-drop
mysql-migrate-drop:
	@migrate -database "mysql://$(MYSQL_DB_USER):$(MYSQL_DB_PASSWORD)@tcp($(MYSQL_DB_ADDRESS))/$(MYSQL_DB_NAME)?multiStatements=true" -path=internal/mysql/migrations drop

# Mongo migrations
.PHONY: mongo-migrate-up
mongo-migrate-up:
	@go run cmd/mongo-migrate/*.go -action=up -db-address="mysql://$(MONGO_DB_ADDRESS)" -db-name=$(MONGO_DB_NAME)

.PHONY: mongo-migrate-down
mongo-migrate-down:
	@go run cmd/mongo-migrate/*.go -action=down -db-address="mysql://$(MONGO_DB_ADDRESS)" -db-name=$(MONGO_DB_NAME)

# Testing
.PHONY: test
test:
	@echo Run tests
	@go test $(TEST_OPTS) ./...

.PHONY: test-mysql
test-mysql:
	@echo Run mysql tests
	@go test -tags=integration $(TEST_OPTS) ./internal/mysql/... -db-user=$(MYSQL_DB_USER) -db-password=$(MYSQL_DB_PASSWORD) -db-address=$(MYSQL_DB_ADDRESS) -db-name=$(MYSQL_DB_NAME)

.PHONY: test-mongo
test-mongo:
	@echo Run mongo tests
	@go test -tags=integration $(TEST_OPTS) ./internal/mongo/... -db-address=$(MONGO_DB_ADDRESS) -db-name=$(MONGO_DB_NAME)

test-all:
	@echo Run all tests
	@go test -tags=integration $(TEST_OPTS) ./...

# Build binary
marketplace-up: cmd/$@ $(SOURCES)
	@echo "Building $@"
	@CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ cmd/$@/*.go

# Build docker image
.PHONY: docker-image
docker-image:
	@echo Build docker image
	@docker build -t marketplace .

# House keeping
.PHONY: clean
clean:
	@echo Cleaning...
	@if [ -f ./marketplace-up ]; then rm ./marketplace-up; fi
