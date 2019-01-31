# Benchmark ORM
[![Go Report Card](https://goreportcard.com/badge/github.com/kostozyb/orm-bench)](https://goreportcard.com/report/github.com/kostozyb/orm-bench) 
[![Build Status](https://travis-ci.org/KosToZyB/orm-bench.svg?branch=master)](https://travis-ci.org/KosToZyB/orm-bench)
1) docker-compose up
2) OB_POSTGRES_USER=docker OB_POSTGRES_PASSWORD=dockerpass OB_POSTGRES_DB=test OB_POSTGRES_HOST=localhost OB_POSTGRES_SSL_MODE=disable OB_DB_DRIVER=postgres OB_DB_NAME=test go run cmd/migrations/main.go
3) OB_POSTGRES_USER=docker OB_POSTGRES_PASSWORD=dockerpass OB_POSTGRES_DB=test OB_POSTGRES_HOST=localhost OB_POSTGRES_SSL_MODE=disable OB_DB_DRIVER=postgres OB_DB_NAME=test go test -bench=. ./... -benchmem -benchtime=10s
