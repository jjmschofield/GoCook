#!/usr/bin/env bash

## Provides release phase instructions for Heroku

## Setup
set -e
echo "Preparing for release..."

## Database migrations
echo "Getting github.com/golang-migrate/migrate binary..."
curl -L https://github.com/golang-migrate/migrate/releases/download/v3.4.0/migrate.linux-amd64.tar.gz | tar xvz

echo "Running migrations..."
./migrate.linux-amd64 -database $DB_CONNECTION --path sql/migrations up

## Update Swagger
echo "Getting swagger generator..."
go get -u github.com/swaggo/swag/cmd/swag

echo "Generating swagger..."
swag init
cp -p docs/swagger.json api/public/swagger.json

## Complete
echo "Release preparation complete!"