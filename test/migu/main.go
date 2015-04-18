package main

import "fmt"

//go:generate goma --debug gen-migu --models entity/quest_gen.sql

func main() {
	fmt.Println("Hello!")
}
