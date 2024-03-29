
.PHONY: generate
generate:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate .

.PHONY: test
test:
	go test ./...

.PHONY: setup_db
setup_db:
	docker compose up -d database --wait

.PHONY: lint
lint:
	golangci-lint run

.PHONY: migration_test
migration_test:
	$(OSFLAG) docker-compose run --rm dbmigrator \
		ridgepole \
			-E test \
			-c /andpad/database/config/ridgepole.yaml \
			-f /andpad/database/schemas/Schemafile \
			--apply --verbose; \

.PHONY: migration_dev
migration_dev:
	$(OSFLAG) docker-compose run --rm dbmigrator \
		ridgepole \
			-E dev \
			-c /andpad/database/config/ridgepole.yaml \
			-f /andpad/database/schemas/Schemafile \
			--apply --verbose; \


.PHONY: up-db
up-db:
	docker-compose --profile db up -d

.PHONY: down-db
down-db:
	docker-compose --profile db down
