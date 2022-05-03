# screecher-rest-api

## Overview

An RESTFul API for Screecher. Screecher is a mock social media platform where user can posts screeches.

## Prerequisites

### Environment Variables

* `APP_ENV`
* `PORT`

### Programs, Databases, or Runtime Environments

* Docker
* Docker Compose
* Golang

## How to Run

### With Docker Compose

1. `docker-compose up`

### With Docker

1. `docker build -t screecher_rest_api .`
2. `docker run -p 8080:5000 screecher_rest_api`

### Without Docker or Docker Compose

#### Dev Mode

1. `make install`
2. `make dev`

#### Production Mode

1. `make install`
2. `make build`
3. `make run`
