# Benchmark ORM
[![Go Report Card](https://goreportcard.com/badge/github.com/kostozyb/orm-bench)](https://goreportcard.com/report/github.com/kostozyb/orm-bench)
1) docker-compose up
2) go run cmd/migrations/main.go
3) go test -bench=. ./... -benchmem -benchtime=10s
