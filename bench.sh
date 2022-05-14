#!/usr/bin/env bash
set -euxo pipefail
go test -run=^$ --bench=. -benchmem -benchtime=10000x . -memprofile memprofile.out -cpuprofile cpuprofile.out