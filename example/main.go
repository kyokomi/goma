package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/example/dao"
	"github.com/kyokomi/goma/example/entity"
	"database/sql"
)

//go:generate goma --debug gen --driver=mysql --user=admin --password=password --host=localhost --port=3306 --db=test --config
// MEMO: go:generate go-bindata -o dao/asset_gen.go -pkg dao sql/...

func main() {
	log.SetFlags(log.Llongfile)
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

	name := "test"
	detail := "test detail"
	_, err = dao.Quest(db).Insert(entity.Quest{
		ID:       99,
		Name:     &name,
		Detail:   &detail,
		CreateAt: &now,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(db).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("insert after: %+v\n", q)
	}

	name = "test 2"
	detail = "test detail 2"
	_, err = dao.Quest(db).Update(entity.Quest{
		ID:       99,
		Name:     &name,
		Detail:   &detail,
		CreateAt: &now,
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

	if q, err := dao.Quest(db).SelectByID(99); err != sql.ErrNoRows {
		log.Fatalln(err)
	} else {
		fmt.Printf("delete after: %+v\n", q)
	}

	if q, err := dao.Quest(db).SelectLimit(100); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("SelectLimit: %+v\n", q)
	}
}
