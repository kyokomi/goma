package main

import (
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
	
	"github.com/kyokomi/goma/example/dao"
	"github.com/kyokomi/goma/goma"
)

//go:generate goma -driver=mysql -source=admin:password@tcp(localhost:3306)/test

func main() {
	fmt.Println("Hello goma!")

	opts := goma.Options{
		Driver: "mysql",
		Source: "admin:password@tcp(localhost:3306)/test",
		Debug:  false,
	}
	g, err := goma.NewGoma(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	q, err := dao.Quest(g).SelectByID(1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", q)
}
