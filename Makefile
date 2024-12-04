POSTGRES_DB := pet_one
POSTGRES_USER := user
POSTGRES_PASSWORD := 123


PWD := $(shell pwd)
USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
DOCKER := DOCKER_BUILDKIT=1 $(shell which docker)

.PHONY: build

up-build:
	docker-compose up -d --build

build:
	go build -v ./cmd/api/

.DEFAULT_GOAL := build