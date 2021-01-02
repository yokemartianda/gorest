.PHONY: default help build run

SHELL         = /bin/bash
APP_NAME      = gorest

default: help

help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                 Compile the project.'
	@echo

build:
	@echo "Building ${APP_NAME}"
	go build -o bin/${APP_NAME}

run: build
	@echo "Running ${APP_NAME}"
	bin/${APP_NAME} ${ARGS}