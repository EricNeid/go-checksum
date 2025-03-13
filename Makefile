# SPDX-FileCopyrightText: 2021 Eric Neidhardt
# SPDX-License-Identifier: CC0-1.0
DIR := ${CURDIR}
DOCKER := podman
GO_IMAGE := golang:1.22-alpine
LINTER_IMAGE := golangci/golangci-lint:v1.54-alpine
CMD_PATH := ./cmd/checksum


.PHONY: build-windows
build-windows:
	${DOCKER} run -it --rm \
		-e GOOS=windows \
		-e GOARCH=amd64 \
		-w /app -v ${DIR}:/app \
		${GO_IMAGE} \
		go build -o ./out/ ${CMD_PATH}


.PHONY: build-linux
build-linux:
	${DOCKER} run -it --rm \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-w /app -v ${DIR}:/app \
		${GO_IMAGE} \
		go build -o ./out/ ${CMD_PATH}


.PHONY: test
test:
	${DOCKER} run -it --rm \
		-w /app -v ${DIR}:/app \
		${GO_IMAGE} \
		go test ./...


.PHONY: cover
cover:
	${DOCKER} run -it --rm \
		-w /app -v ${DIR}:/app \
		${GO_IMAGE} \
		mkdir -p out && go test -coverprofile=out/cover.out ./... && go tool cover -html=out/cover.out


.PHONY: format
format:
	${DOCKER} run -it --rm \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-w /app -v ${DIR}:/app \
		${GO_IMAGE} \
		go fmt ./...


.PHONY: lint
lint:
	${DOCKER} run -it --rm \
		-e CGO_ENABLED=0 \
		-w /app -v ${DIR}:/app \
		${LINTER_IMAGE} \
		golangci-lint --timeout=5m run ./...