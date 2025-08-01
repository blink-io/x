#!/usr/bin/env bash

export DATABASE_URL='postgres://test:test@192.168.50.88:5432/test?sslmode=disable'
export DB="$DATABASE_URL"

go install -tags=fts5 github.com/blink-io/sqddl@latest

sqddl views \
    -db "$DB"
