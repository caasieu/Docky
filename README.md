# Docky

A lightweight Docker container and image manager written in Go.

## Features

- Container lifecycle management
- Image management
- Container logs
- Container inspection
- CLI interface
- Docker Engine API integration
- Unit and integration tests

## Architecture

CLI
 │
 ▼
Application Services
 │
 ▼
Domain Interfaces
 │
 ▼
Docker Infrastructure
 │
 ▼
Docker Engine

## Tech Stack

- Go
- Docker Engine API
- Cobra
- Docker
- Go testing

## Running

go mod download
go run ./cmd/dockyard
