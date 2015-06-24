#!/usr/bin/env bash

go generate github.com/kyokomi/goma/gen/goma
go install github.com/kyokomi/goma/gen/goma
go generate github.com/kyokomi/goma/test/mysql
go generate github.com/kyokomi/goma/test/postgres
go test .
go test ./gen/goma/...
go test ./test/...

golint ./test/... | grep -v ego

