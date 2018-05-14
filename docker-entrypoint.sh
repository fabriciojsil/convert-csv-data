#!/bin/bash
set -e

go run cmd/main.go -file="${FILE}" -format="${FORMAT}" -order="$ORDER" -field="${FIELD}"

exec "$@"
