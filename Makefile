PORT_DB := 8932
POSTGRES_DB := pet_one
POSTGRES_USER := pet_one
POSTGRES_PASSWORD := pass123


PWD := $(shell pwd)
USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
DOCKER := DOCKER_BUILDKIT=1 $(shell which docker)

.PHONY: build
.DEFAULT_GOAL := build

up:
	docker-compose up -d

start:
	docker-compose up -d --build

down:
	docker-compose down

build:
	go build -v ./cmd/api/

migrate_up:
	migrate -path migrations -database "postgres://localhost:${PORT_DB}/${POSTGRES_DB}?sslmode=disable&user=${POSTGRES_USER}&password=${POSTGRES_PASSWORD}" up

migrate_down:
	migrate -path migrations -database "postgres://localhost:${PORT_DB}/${POSTGRES_DB}?sslmode=disable&user=${POSTGRES_USER}&password=${POSTGRES_PASSWORD}" down
