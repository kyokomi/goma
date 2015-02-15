package main

import (
	"fmt"

	"log"

	"github.com/k0kubun/pp"
	"github.com/kyokomi/goma/example/dao"
	"github.com/kyokomi/goma/goma"
)

func main() {
	log.SetFlags(log.Llongfile)
	fmt.Println("Hello doma!")

	opts := goma.Options{
		Driver: "mysql",
		Source: "admin:password@tcp(localhost:3306)/test",
		Debug:  true,
	}
	g, err := goma.NewGoma(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	q, err := dao.Quest(g).SelectByID(goma.QueryArgs{
		"id":   1,
		"name": "quest1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(q)

	qs, err := dao.Quest(g).SelectAll()
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(qs)
}
