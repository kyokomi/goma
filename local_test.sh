#!/usr/bin/env bash

go generate github.com/kyokomi/goma/gen/goma
go install github.com/kyokomi/goma/gen/goma
go generate github.com/kyokomi/goma/test/mysql
go generate github.com/kyokomi/goma/test/postgres
go generate github.com/kyokomi/goma/test/simple/mysql
go generate github.com/kyokomi/goma/test/simple/postgres
go test ./test/...

golint ./test/... | grep -v ego

