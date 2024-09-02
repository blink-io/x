#!/usr/bin/env bash

DATABASE_URL='postgres://blink:888asdf%21%23%25@192.168.50.88:5432/orm-demo?sslmode=disable'
#DATABASE_URL='postgres://blink:888asdf!#%@192.168.50.88:5432/orm-demo?sslmode=disable'

go install -tags=fts5 github.com/bokwoon95/sqddl@latest

sqddl generate -db "$DATABASE_URL" \
  -dest ./tables.go \
  -output-dir ./