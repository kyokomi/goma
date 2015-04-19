package main

import "fmt"

//go:generate goma --debug gen-migu --models entity/

func main() {
	fmt.Println("Hello!")
}
