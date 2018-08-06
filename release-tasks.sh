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

## Documentation
echo "Gettiing go-swagger binary..."
curl -L https://github.com/go-swagger/go-swagger/releases/download/0.15.0/swagger_linux_amd64 > swagger
ls
echo "Regenerating swagger documents..."
./swagger generate spec -o api/public/swagger.json

## Complete
echo "Release preparation complete!"