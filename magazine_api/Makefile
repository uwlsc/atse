include .env
export

MIGRATE=docker-compose exec web sql-migrate

ifeq ($(p),host)
 	MIGRATE=sql-migrate
endif

migrate-status:
	$(MIGRATE) status

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 

check_install:
	which swag || go install github.com/swaggo/swag/cmd/swag@latest

swagger_format:
	swag fmt -g http/application.go

swagger: check_install swagger_format
	swag init -g application.go 

redo:
	@read -p  "Are you sure to reapply the last migration? [y/n]" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		$(MIGRATE) redo; \
	fi

create:
	@read -p  "What is the name of migration?" NAME; \
	${MIGRATE} new $$NAME

TEST_RUNNER=docker-compose exec web go

ifeq ($(p),host)
	TEST_RUNNER=go
endif

test:
	@echo "running tests 🧪 ..."
	$(TEST_RUNNER) test -v ./...

.PHONY: migrate-status migrate-up migrate-down redo create