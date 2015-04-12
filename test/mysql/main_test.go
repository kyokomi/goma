package main

import (
	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"

	"reflect"
	"testing"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/test/mysql/dao"
	"github.com/kyokomi/goma/test/mysql/entity"
)

func TestNumeric(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	id := int64(1234567890)

	// numeric
	d := dao.GomaNumericTypes(db)

	insertData := entity.GomaNumericTypesEntity{
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
	}

	if _, err := d.Insert(insertData); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if !reflect.DeepEqual(e, &insertData) {
		t.Errorf("ERROR: %+v != %+v", e, &insertData)
	}

	if _, err := d.Delete(id); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Deleteしたのにnilじゃない")
	}
}

func TestString(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	id := int64(1234567890)

	// string
	d := dao.GomaStringTypes(db)

	insertData := entity.GomaStringTypesEntity{
		ID:                id,
		TextColumns:       "あいうえおかきくけこ",
		TinytextColumns:   "abc",
		MediumtextColumns: "abcdefg",
		LongtextColumns:   "鉄1234567890abcdefghijkelmnopqrstuvwxyz1234567890abcdefghijkelmnopqrstuvwxyz柱",
		CharColumns:       "a",
		VarcharColumns:    "1234567890abcdefghijkelmnopqrstuvwxyz",
	}

	if _, err := d.Insert(insertData); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if !reflect.DeepEqual(e, &insertData) {
		t.Errorf("ERROR: %+v != %+v", e, &insertData)
	}

	if _, err := d.Delete(id); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Deleteしたのにnilじゃない")
	}
}

func TestDate(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	id := int64(1234567890)

	// date
	d := dao.GomaDateTypes(db)

	now := time.Now()
	insertData := entity.GomaDateTypesEntity{
		ID:               id,
		DateColumns:      time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
		DatetimeColumns:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
		TimestampColumns: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
	}

	if _, err := d.Insert(insertData); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if !reflect.DeepEqual(e, &insertData) {
		t.Errorf("ERROR: %+v != %+v", e, &insertData)
	}

	if _, err := d.Delete(id); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Deleteしたのにnilじゃない")
	}
}

func TestBinary(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	id := int64(1234567890)

	// date
	d := dao.GomaBinaryTypes(db)

	insertData := entity.GomaBinaryTypesEntity{
		ID:                id,
		BinaryColumns:     []uint8{49, 49, 49},
		TinyblobColumns:   []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		BlobColumns:       []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		MediumblobColumns: []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
		LongblobColumns:   []uint8{110, 111, 112, 113, 114, 115, 116, 117, 118},
		VarbinaryColumns:  []uint8{49, 49, 50, 51, 52, 53, 54, 55, 56},
	}

	if _, err := d.Insert(insertData); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if !reflect.DeepEqual(e, &insertData) {
		t.Errorf("ERROR: %+v != %+v", e, &insertData)
	}

	if _, err := d.Delete(id); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Deleteしたのにnilじゃない")
	}
}

func TestTx(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	log.SetFlags(log.Llongfile)

	id := int64(1234567890)

	tx, err := db.Begin()
	if err != nil {
		t.Errorf("ERROR: %s", err)
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
		t.Errorf("ERROR: %s", err)
	}

	// Rollback（insertを無効にする）
	if err := tx.Rollback(); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	// rollbackしたのでトランザクション生成しなおす
	tx, err = db.Begin()
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}

	// string
	dtx = dao.TxGomaStringTypes(tx)

	// Rollbackでnilのはず
	if e, err := dtx.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Rollbackしたのにnilじゃない")
	}

	// Insertする
	_, err = dtx.Insert(e)
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}

	// Commit
	if err := tx.Commit(); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	// commitしたのでtxじゃないdao
	d := dao.GomaStringTypes(db)

	// Commitしたのでnilじゃない
	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e == nil {
		t.Errorf("ERROR: %s", "Commitしたのにnil")
	}

	_, err = d.Delete(id)
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}

	if e, err := d.SelectByID(id); err != nil {
		t.Errorf("ERROR: %s", err)
	} else if e != nil {
		t.Errorf("ERROR: %s", "Deleteしたのにnilじゃない")
	}
}
