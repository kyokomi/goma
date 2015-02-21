package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kyokomi/goma/example/dao"
)

//go:generate goma -driver=mysql -user=admin -password=password -host=localhost -port=3306 -db=test -debug=true

func main() {
	fmt.Println("Hello goma!")

	g, err := Goma()
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	// TODO: 念のため削除
	_, err = dao.Quest(g).Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()

	_, err = dao.Quest(g).Insert(dao.QuestEntity{
		ID:       99,
		Name:     "test",
		Detail:   "test detail",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("insert after: %+v\n", q)
	}

	_, err = dao.Quest(g).Update(dao.QuestEntity{
		ID:       99,
		Name:     "test 2",
		Detail:   "test detail 2",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("update after: %+v\n", q)
	}

	_, err = dao.Quest(g).Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("delete after: %+v\n", q)
	}
}
