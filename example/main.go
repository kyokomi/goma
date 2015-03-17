package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kyokomi/goma/example/entity"
)

//go:generate goma --debug gen --driver=mysql --user=admin --password=password --host=localhost --port=3306 --db=test

func main() {
	fmt.Println("Hello goma!")

	goma, err := NewGoma("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer goma.Close()

	// TODO: 念のため削除
	_, err = goma.Quest().Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()

	_, err = goma.Quest().Insert(entity.QuestEntity{
		ID:       99,
		Name:     "test",
		Detail:   "test detail",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := goma.Quest().SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("insert after: %+v\n", q)
	}

	_, err = goma.Quest().Update(entity.QuestEntity{
		ID:       99,
		Name:     "test 2",
		Detail:   "test detail 2",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := goma.Quest().SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("update after: %+v\n", q)
	}

	_, err = goma.Quest().Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := goma.Quest().SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("delete after: %+v\n", q)
	}
}
