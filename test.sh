#!/usr/bin/env bash
set -euxo pipefail
go test --race ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
