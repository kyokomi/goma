package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate goma --debug gen --driver=mysql --user=root --host=localhost --port=3306 --db=goma_test --config
// MEMO: go:generate go-bindata -o dao/asset_gen.go -pkg dao sql/...

func main() {
	fmt.Println("Hello goma!")
}
