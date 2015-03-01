package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kyokomi/goma/test/postgres/entity"
)

//go:generate goma -driver=postgres -user=admin -password=password -host=localhost -port=5432 -db=goma-test -debug=true -ssl=disable

func main() {
	fmt.Println("Hello goma!")

	goma, err := NewGoma()
	if err != nil {
		log.Fatalln(err)
	}
	defer goma.Close()

	numericTest(goma)
	stringTest(goma)
	dateTest(goma)
	txTest(goma)
}

func numericTest(g Goma) {

	id := int64(1234567890)

	// numeric
	dao := g.GomaNumericTypes()

	_, err := dao.Insert(entity.GomaNumericTypesEntity{
		ID:               id,
		BoolColumns:      true,
		SmallintColumns:  int(123),
		IntColumns:       int(11111111),
		IntegerColumns:   int(22222222),
		SerialColumns:    1234567890,
		DecimalColumns:   "1234567890",
		NumericColumns:   "1234567890",
		FloatColumns:     float64(1.234),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}

	_, err = dao.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}
}

func stringTest(g Goma) {

	id := int64(1234567890)

	// string
	dao := g.GomaStringTypes()

	_, err := dao.Insert(entity.GomaStringTypesEntity{
		ID:                id,
		TextColumns:       "あいうえおかきくけこ",
		CharColumns:       "a",
		VarcharColumns:    "1234567890abcdefghijkelmnopqrstuvwxyz",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}

	_, err = dao.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}
}

func dateTest(g Goma) {

	id := int64(1234567890)

	// date
	dao := g.GomaDateTypes()

	_, err := dao.Insert(entity.GomaDateTypesEntity{
		ID:               id,
		DateColumns:      time.Now(),
		TimestampColumns: time.Now(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}

	_, err = dao.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", dao.TableName, e)
	}
}

func txTest(g Goma) {

	id := int64(1234567890)

	tx, err := g.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	// string
	dao := g.GomaStringTypes()
	dao.SetTx(tx)

	e := entity.GomaStringTypesEntity{
		ID:                id,
		TextColumns:       "あいうえおかきくけこ",
		CharColumns:       "a",
		VarcharColumns:    "1234567890abcdefghijkelmnopqrstuvwxyz",
	}

	_, err = dao.Insert(e)
	if err != nil {
		log.Fatalln(err)
	}

	// Rollback（insertを無効にする）
	if err := tx.Rollback(); err != nil {
		log.Fatalln(err)
	}

	// Rollback後に使うならResetもしくは新しいtxを設定する
	dao.ResetTx()

	// 再度トランザクションをはる
	tx, err = g.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	dao.SetTx(tx)

	// Rollbackでnilのはず
	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Rollbackでnilのはず => %s: %+v\n", dao.TableName, e)
	}

	// Insertする
	_, err = dao.Insert(e)
	if err != nil {
		log.Fatalln(err)
	}

	// Commit
	if err := tx.Commit(); err != nil {
		log.Fatalln(err)
	}
	// Commit後後に使うならResetもしくは新しいtxを設定する
	dao.ResetTx()

	// Commitしたのでnilじゃない
	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Commitしたのでnilじゃない => %s: %+v\n", dao.TableName, e)
	}

	_, err = dao.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := dao.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("普通にDeleteしてnilのはず => %s: %+v\n", dao.TableName, e)
	}
}
