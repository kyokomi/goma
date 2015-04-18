package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

//go:generate goma --debug gen --driver=postgres --user=postgres --host=localhost --port=5432 --db=goma_test --ssl=disable
// MEMO: go:generate go-bindata -o dao/asset_gen.go -pkg dao sql/...

func main() {
	fmt.Println("Hello goma!")
}
