package main

import (
	"fmt"
	"log"

	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/test/mysql/dao"
	"github.com/kyokomi/goma/test/mysql/entity"
)

//go:generate goma --debug gen --driver=mysql --user=admin --password=password --host=localhost --port=3306 --db=goma_test

// MEMO: go-bindata -o dao/asset_gen.go -pkg dao sql/...

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
	binaryTest(db)
	txTest(db)
}

func numericTest(db *sql.DB) {
	id := int64(1234567890)

	// numeric
	d := dao.GomaNumericTypes(db)

	_, err := d.Insert(entity.GomaNumericTypesEntity{
		ID:               id,
		TinyintColumns:   int(8),
		BoolColumns:      int(1),
		SmallintColumns:  int(123),
		MediumintColumns: int(256),
		IntColumns:       int(11111111),
		IntegerColumns:   int(22222222),
		SerialColumns:    int64(1234567890),
		DecimalColumns:   "1234567890",
		NumericColumns:   "1234567890",
		FloatColumns:     float32(1.234),
		DoubleColumns:    float64(1000.234),
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
		ID:                id,
		TextColumns:       "あいうえおかきくけこ",
		TinytextColumns:   "abc",
		MediumtextColumns: "abcdefg",
		LongtextColumns:   "鉄1234567890abcdefghijkelmnopqrstuvwxyz1234567890abcdefghijkelmnopqrstuvwxyz柱",
		CharColumns:       "a",
		VarcharColumns:    "1234567890abcdefghijkelmnopqrstuvwxyz",
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
		DatetimeColumns:  time.Now(),
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

func binaryTest(db *sql.DB) {
	id := int64(1234567890)

	var err error

	// date
	d := dao.GomaBinaryTypes(db)

	_, err = d.Insert(entity.GomaBinaryTypesEntity{
		ID:                id,
		BinaryColumns:     []uint8{49},
		TinyblobColumns:   []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		BlobColumns:       []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		MediumblobColumns: []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		LongblobColumns:   []uint8{110, 111, 112, 113, 114, 115, 116, 117, 118},
		VarbinaryColumns:  []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
	})
	if err != nil {
		log.Fatalln(err)
	}

	if e, err := d.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %s\n", d.TableName, string(e.BinaryColumns))
		fmt.Printf("%s: %s\n", d.TableName, string(e.TinyblobColumns))
		fmt.Printf("%s: %s\n", d.TableName, string(e.BlobColumns))
		fmt.Printf("%s: %s\n", d.TableName, string(e.MediumblobColumns))
		fmt.Printf("%s: %s\n", d.TableName, string(e.LongblobColumns))
		fmt.Printf("%s: %s\n", d.TableName, string(e.VarbinaryColumns))
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
	log.SetFlags(log.Llongfile)

	id := int64(1234567890)

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	// string
	dtx := dao.TxGomaStringTypes(tx)

	e := entity.GomaStringTypesEntity{
		ID:                id,
		TextColumns:       "あいうえおかきくけこ",
		TinytextColumns:   "abc",
		MediumtextColumns: "abcdefg",
		LongtextColumns:   "鉄1234567890abcdefghijkelmnopqrstuvwxyz1234567890abcdefghijkelmnopqrstuvwxyz柱",
		CharColumns:       "a",
		VarcharColumns:    "1234567890abcdefghijkelmnopqrstuvwxyz",
	}

	_, err = dtx.Insert(e)
	if err != nil {
		log.Fatalln(err)
	}

	// Rollback（insertを無効にする）
	if err := tx.Rollback(); err != nil {
		log.Fatalln(err)
	}

	// rollbackしたのでトランザクション生成しなおす
	tx, err = db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	// string
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
