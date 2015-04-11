package main

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/test/postgres/dao"
	"github.com/kyokomi/goma/test/postgres/entity"
)

//go:generate goma --debug gen --driver=postgres --user=admin --password=password --host=localhost --port=5432 --db=goma-test --ssl=disable

func main() {
	fmt.Println("Hello goma!")

	db, err := goma.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	numericTest(db)
	stringTest(db)
	dateTest(db)
	txTest(db)
}

func numericTest(db *sql.DB) {
	id := int64(1234567890)

	// numeric
	d := dao.GomaNumericTypes(db)

	_, err := d.Insert(entity.GomaNumericTypesEntity{
		ID:              id,
		BoolColumns:     true,
		SmallintColumns: int(123),
		IntColumns:      int(11111111),
		IntegerColumns:  int(22222222),
		SerialColumns:   1234567890,
		DecimalColumns:  "1234567890",
		NumericColumns:  "1234567890",
		FloatColumns:    float64(1.234),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}

	_, err = d.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}
}

func stringTest(db *sql.DB) {
	id := int64(1234567890)

	// string
	d := dao.GomaStringTypes(db)

	_, err := d.Insert(entity.GomaStringTypesEntity{
		ID:             id,
		TextColumns:    "あいうえおかきくけこ",
		CharColumns:    "a",
		VarcharColumns: "1234567890abcdefghijkelmnopqrstuvwxyz",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}

	_, err = d.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}
}

func dateTest(db *sql.DB) {
	id := int64(1234567890)

	// date
	d := dao.GomaDateTypes(db)

	_, err := d.Insert(entity.GomaDateTypesEntity{
		ID:               id,
		DateColumns:      time.Now(),
		TimestampColumns: time.Now(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}

	_, err = d.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", d.TableName, e)
	}
}

func txTest(db *sql.DB) {
	id := int64(1234567890)

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	// string
	dtx := dao.TxGomaStringTypes(tx)

	e := entity.GomaStringTypesEntity{
		ID:             id,
		TextColumns:    "あいうえおかきくけこ",
		CharColumns:    "a",
		VarcharColumns: "1234567890abcdefghijkelmnopqrstuvwxyz",
	}

	_, err = dtx.Insert(e)
	if err != nil {
		log.Fatalln(err)
	}

	// Rollback（insertを無効にする）
	if err := tx.Rollback(); err != nil {
		log.Fatalln(err)
	}

	// 再度トランザクションをはる
	tx, err = db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	dtx = dao.TxGomaStringTypes(tx)

	// Rollbackでnilのはず
	if e, err := dtx.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Rollbackでnilのはず => %s: %+v\n", dtx.TableName, e)
	}

	// Insertする
	_, err = dtx.Insert(e)
	if err != nil {
		log.Fatalln(err)
	}

	// Commit
	if err := tx.Commit(); err != nil {
		log.Fatalln(err)
	}

	// commitしたのでtxじゃないdao
	d := dao.GomaStringTypes(db)

	// Commitしたのでnilじゃない
	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Commitしたのでnilじゃない => %s: %+v\n", d.TableName, e)
	}

	_, err = d.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("普通にDeleteしてnilのはず => %s: %+v\n", d.TableName, e)
	}
}
