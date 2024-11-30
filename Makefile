cf = -f deploy/docker/compose.yaml
af = -f deploy/docker/compose-api-test.yaml
amf = -f deploy/docker/compose-api-test-override.yaml

help: ## Print this help
	@grep -E '^[a-zA-Z0-9_-]+:.*## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Run in docker

build: ## Build docker containers
	docker compose $(cf) build
up: ## Start docker containers
	docker compose $(cf) up -d --remove-orphans
down: ## Stop docker containers
	docker compose $(cf) down
rebuild: ## Rebuild and start docker containers
	@make down
	@make build
	@make up
restart: ## Restart docker containers
	docker compose $(cf) restart

# Hurl API testing in docker

apitestbuild: ## Build containers for API testing
	docker compose $(cf) $(amf) build
apitestup: ## Start containers for API testing
	docker compose $(cf) $(amf) up -d --remove-orphans
apitestdown: ## Stop containers for API testing
	docker compose $(cf) $(amf) down
apitestrun: ## Run Hurl testing scripts in docker container and in mutual network
	docker run --rm -v ./test/:/test --net frrapit-news-public ghcr.io/orange-opensource/hurl:latest --test --color --variables-file=/test/api/docker-vars /test/api/news.hurl
apitest: ## Build and start docker services and run API testing on them
	@make apitestbuild
	@make apitestup
	@make apitestrun
	@make apitestdown

# Local development

hurl: ## Run hurl API testing on localhost installation
	hurl --variables-file=./test/api/local-vars ./test/api/news.hurl
gen: ## Generate code for reform logic
	go generate ./...
swag: ## Generate Swagger documentation for REST API
	swag init --dir . -g ./internal/api/rest/router/router.go
# test: ## Run unit tests
# 	go test -count=1 ./...
fmt: ## Format code
	gofumpt -w .
lint: ## Check linter inspections
	golangci-lint run
install-dev-tools: ## Install developer tools (linter, formatter)
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

dbdockerinit:
	docker network create frr-local
dbup: dbdockerinit
	docker run --name frr-news-storage -p3307:3306 --network frr-local \
	-e MYSQL_ROOT_PASSWORD=pw -e MYSQL_DATABASE=frr -e MYSQL_USER=admin -e MYSQL_PASSWORD=123 \
	-v /home/coding/repos/go/go_news_rest_api_tt/deploy/docker/storage/initdb:/docker-entrypoint-initdb.d:ro \
	-d mysql:latest
dbdown:
	docker stop frr-news-storage
	docker rm frr-news-storage
	docker network rm frr-local
dbrestart:
	@make dbdown
	@make dbup
dbclient:
	docker run -it --rm --network frr-local --name frr-mysql-client mysql mysql -hfrr-news-storage -uadmin -p --database frr
dbclientdown:
	docker stop frr-mysql-client
localdbclient:
	mysql --port 3307 -uadmin -p --database frr
appup:
	go run ./cmd/

.PHONY: \
		apitest \
		apitestbuild \
		apitestdown \
		apitestrun \
		apitestup \
		build \
		dbclient \
		dbclientdown \
		dbdown \
		dbrestart \
		dbup \
		dockerinit \
		down \
		fmt \
		gen \
		help \
		hurl \
		install-dev-tools \
		lint \
		localdbclient \
		rebuild \
		restart \
		up \
