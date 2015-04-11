package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/example/dao"
	"github.com/kyokomi/goma/example/entity"
)

//go:generate goma --debug gen --driver=mysql --user=admin --password=password --host=localhost --port=3306 --db=test
// MEMO: go:generate go-bindata -o dao/asset_gen.go -pkg dao sql/...

func main() {
	fmt.Println("Hello goma!")

	db, err := goma.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// TODO: 念のため削除
	_, err = dao.Quest(db).Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()

	_, err = dao.Quest(db).Insert(entity.QuestEntity{
		ID:       99,
		Name:     "test",
		Detail:   "test detail",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(db).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("insert after: %+v\n", q)
	}

	_, err = dao.Quest(db).Update(entity.QuestEntity{
		ID:       99,
		Name:     "test 2",
		Detail:   "test detail 2",
		CreateAt: now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(db).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("update after: %+v\n", q)
	}

	_, err = dao.Quest(db).Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(db).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("delete after: %+v\n", q)
	}
}
