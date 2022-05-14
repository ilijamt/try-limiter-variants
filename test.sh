#!/usr/bin/env bash
set -euxo pipefail
go test --race -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
