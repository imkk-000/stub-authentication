#!/bin/sh

go clean -testcache
go test ./service ./api
CGO_ENABLED=0 go build -o bin/app cmd/login/main.go
./bin/app&
go test ./atdd && kill -9 $$
