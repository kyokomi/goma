package main

import "fmt"

//go:generate goma --debug gen-migu --user root --db test --models models/models.go --entity models --config false

func main() {
	fmt.Println("Hello!")
}
