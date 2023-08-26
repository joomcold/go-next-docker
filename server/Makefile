include .env

GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BUILD_FILE = release/go-next-docker
MAIN_FILE = cmd/main.go

all: clean build run

build:
	$(GOBUILD) -o $(BUILD_FILE) $(MAIN_FILE)

test:
	$(GOTEST) -v ./...

clean:
	rm -f $(BUILD_FILE)

run:
	$(GOCMD) run $(MAIN_FILE)

db_create:
	docker exec -it postgresql psql -U $(dbUser) -c "CREATE DATABASE $(dbName);"

db_drop:
	docker exec -it postgresql psql -U $(dbUser) -c "DROP DATABASE $(dbName);"

lint:
	golangci-lint run

.PHONY: all build test clean run db_create db_drop lint
