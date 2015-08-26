#!/usr/bin/env bash

mysql -u root < ./test/mysql/ddl/create_database.sql
mysql goma_test -u root < ./test/mysql/ddl/create_table_data.sql
mysql goma_test -u root < ./test/mysql/ddl/create_table_numeric.sql
mysql goma_test -u root < ./test/mysql/ddl/create_table_string.sql
mysql goma_test -u root < ./test/mysql/ddl/create_table_binary.sql
psql -U postgres -f ./test/postgres/ddl/create_database.sql
psql -d goma_test -U postgres -f ./test/postgres/ddl/create_table_numeric.sql
psql -d goma_test -U postgres -f ./test/postgres/ddl/create_table_string.sql
psql -d goma_test -U postgres -f ./test/postgres/ddl/create_table_date.sql

go generate github.com/kyokomi/goma/gen/goma
go install github.com/kyokomi/goma/gen/goma
go generate github.com/kyokomi/goma/test/mysql
go generate github.com/kyokomi/goma/test/postgres
go test .
go test ./gen/goma/...
go test ./test/...

golint ./test/... | grep -v ego

