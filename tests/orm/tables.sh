#!/usr/bin/env bash

export DATABASE_URL='postgres://blink:888asdf%21%23%25@192.168.50.88:5432/orm-demo?sslmode=disable'
#DATABASE_URL='postgres://blink:888asdf!#%@192.168.50.88:5432/orm-demo?sslmode=disable'
DB="$DATABASE_URL"

go install -tags=fts5 github.com/bokwoon95/sqddl@latest

sqddl tables -db "$DB"