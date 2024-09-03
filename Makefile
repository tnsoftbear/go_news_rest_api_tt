cf = -f deploy/docker/compose.yaml
af = -f deploy/docker/compose-api-test.yaml

# Deploy to local docker 

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
restart: ## Restart
	docker compose $(cf) restart

# E2E API testing
apitestbuild:
	docker compose $(af) build
apitestup:
	docker compose $(af) up -d
apitestdown:
	docker compose $(af) down
apitestrun:
	docker run --rm -v .\test\:/test --net frr-news-public ghcr.io/orange-opensource/hurl:latest --test --color --variables-file=/test/api/docker-vars /test/api/news.hurl
apitest: ## Build and start docker services and run API testing on them
	@make apitestbuild
	@make apitestup
	@make apitestrun
	@make apitestdown

# Local development

gen: ## Generate code for reform logic
	go generate ./...
hurl: ## Run hurl API testing on localhost installation
	hurl --variables-file=.\test\api\local-vars .\test\api\news.hurl

dockerinit:
	docker network create frr
dbup:	# https://hub.docker.com/_/mysql
	docker run --name frr-news-storage -p3307:3306 --network frr \
	-e MYSQL_ROOT_PASSWORD=pw -e MYSQL_DATABASE=frr -e MYSQL_USER=admin -e MYSQL_PASSWORD=123 \
	-v E:/coding/my/go/own/fiber-reform-rest/deploy/docker/storage/initdb:/docker-entrypoint-initdb.d:ro \
	-d mysql:latest
dbdown:
	docker stop frr-news-storage
	docker rm frr-news-storage
dbrestart:
	@make dbdown
	@make dbup
dbclient:
	docker run -it --rm --network frr --name frr-mysql-client mysql mysql -hfrr-news-storage -uadmin -p --database frr
dbclientdown:
	docker stop frr-mysql-client
localdbclient:
	mysql.exe --port 3307 -uadmin -p --database frr

.PHONY: \
		build \
		down \
		rebuild \
		up \