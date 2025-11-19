APP_COMPOSE_FILE := -f docker-compose.yml
APP_COMPOSE_FILE_DEV := -f docker-compose-dev.yml
APP_SERVICE := eform-v3
MYSQL_URL := mysql://localhost:49155/riskmanagement2
include .env
export

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

kill:
	kill -9 $(lsof -t -i:5000)
	
migrate-create:  ### create new migration
	@echo "Input Table Name>"
	@read name; migrate create -ext sql -dir config/database $$name
		# @read name; migrate create -ext sql -dir config/database -seq $$name
.PHONY: migrate-create

migrate-up: ### migration up
	# @if ! command -v migrate &> /dev/null; then go get -d github.com/golang-migrate/migrate/v4 ; fi
	# @migrate -path config/database -database '$(MYSQL_URL)?multiStatements=true' -verbose up
	# @migrate -path config/database -database "mysql://root:P@ssw0rd@127.0.0.1:49155/riskmanagement2?sslmode=disable" -verbose up
	@migrate -source "file://config/database" -database "mysql://root:P@ssw0rd@tcp(localhost:3306)/riskmanagement2" up
.PHONY: migrate-up

migrate-down: ### migration up
	migrate -path migrations -database '$(MYSQL_URL)?multiStatements=true' -verbose down
.PHONY: migrate-down


test: 
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

docker-start:
	@echo "${NOW} STARTING CONTAINER..."
	@docker-compose up -d --build

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@docker-compose stop
 
docker-down:
	@echo "${NOW} STOPPING & REMOVING CONTAINER..."
	@docker-compose down

docker-rebuilddb:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	docker exec -it basedb  mysql -uroot -ppassword -e'drop database if exists ${DB}'
	@echo "${NOW} CREATE DB..."
	docker exec -it basedb  mysql -uroot -ppassword -e'create database ${DB}'
	@echo "${NOW} RUN SQL SCRIPTS..." 
	docker exec -it basedb setup.sh /config/database

swag:
	@echo "> Generate Swagger Docs"
	@if ! command swag -v &> /dev/null; then go install github.com/swaggo/swag ; fi
	@swag init --parseVendor

build:
	@echo "> Building Project"
	@go build ${MAIN}
	@echo "> Copying files to stagging folder"
	@cp -r main .env .env.development .env.production .env.staging database staging

dump:
	@echo "> Dump Database"
	@mysqldump -u ${DB_USER} -p ${DB_NAME} > staging/database/${DB_NAME}.sql

restore:
	@echo "> Restore Database"
	@mysqldump -u ${DB_USER} -p ${DB_NAME} > staging/database/${DB_NAME}.sql

restart:
	@echo "> Restarting service app ${APP_NAME}"
	@sudo service ${APP_NAME} restart

zip:
	@echo "> Input Name of file:"
	@read name; zip -r $$name.zip staging;
postman:
	@ps aux | grep -ie Postman | awk '{print $2}' | xargs kill -9

app:
	sudo docker-compose ${APP_COMPOSE_FILE} up -d --build
app-down:
	sudo docker-compose ${APP_COMPOSE_FILE} down -v
push:
	@echo "Push To dev-dik and development"
	@git push origin dev-dik;
	@git checkout development; 
	@git merge dev-dik;
	@git push origin development; 
	@git checkout master;
	@git merge dev-dik;
	@git push origin master;
	@git checkout dev-dik;
lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	@golangci-lint run