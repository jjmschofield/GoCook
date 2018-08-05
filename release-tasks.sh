#!/usr/bin/env bash

## Provides release phase instructions for Heroku

## Setup
set -e
echo "Preparing for release..."

## Database migrations
echo "Getting github.com/golang-migrate/migrate binary..."
curl -L https://github.com/golang-migrate/migrate/releases/download/v3.4.0/migrate.linux-amd64.tar.gz | tar xvz > migrate

echo "Running migrations..."
./migrate -database $DB_CONNECTION --path sql/migrations up