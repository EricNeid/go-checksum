# SPDX-FileCopyrightText: 2021 Eric Neidhardt
# SPDX-License-Identifier: CC0-1.0

all: test build

build:
	go build ./cmd/checksum/

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build ./cmd/checksum/

build-linux-arm:
	GOOS=linux GOARCH=arm go build ./cmd/checksum/

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build ./cmd/checksum/

test:
	go test ./...