# clients Api

## Summary

* [Description](#description)
* [Dependencies](#dependencies)
* [Project Architecture](#project-architecture)
* [Environment](#environment)
* [Installing](#installing)
* [Building](#building)
* [Running](#running)
* [Testing](#testing)

## Description

API responsible for user managment with Golang

## Dependencies
* [GoLang (v1.17)](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Go Mod](https://github.com/golang/mod) - Go Modules is a way of dealing with dependencies in Go.
* [Gin Web Framework](https://github.com/gin-gonic/gin) - A powerful HTTP router and URL matcher for building Go web servers.
* [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
* [Aurora](https://github.com/logrusorgru/aurora) - Ultimate ANSI colors for Golang. The package supports Printf/Sprintf etc.
* [GoDotEnv](https://github.com/joho/godotenv) - A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).

## Installing
Cloning the Client API project:
``` bash
$ git clone git@github.com:Genarodaniel/ClientApi.git
```

## Building

Building project:

``` bash
$ go get
$ go mod tidy
$ go mod vendor
```

## Running

Running application:
``` bash
$ make run
```

``` bash
$ go run main.go
```

## Testing

Running project's tests:

```bash
$ go test -v cover ./...
```

Generating coverage files and metrics:

```bash
$ go test ./... -coverprofile fmtcoverage.html fmt
$ go test ./... -coverprofile cover.out
$ go tool cover -html=cover.out
$ go tool cover -html=cover.out -o cover.html
```
