#!/usr/bin/env bash
set -euxo pipefail
go test --bench=. -benchmem ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
