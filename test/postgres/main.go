package main

import (
	"fmt"
)

//go:generate goma -driver=postgres -user=admin -password=password -host=localhost -port=3306 -db=goma_test -debug=true

func main() {
	fmt.Println("Hello goma!")

	// TODO: こんどやる
//	goma, err := NewGoma()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer goma.Close()
}
