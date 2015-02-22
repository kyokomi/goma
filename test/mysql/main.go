package main

import (
	"fmt"
	"log"

	"time"

	"github.com/kyokomi/goma/test/mysql/entity"
)

//go:generate goma -driver=mysql -user=admin -password=password -host=localhost -port=3306 -db=goma_test -debug=true

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
}

func numericTest(g Goma) {

	id := int64(1234567890)

	// numeric
	dao := g.GomaNumericTypes

	_, err := dao.Insert(entity.GomaNumericTypesEntity{
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
	dao := g.GomaStringTypes

	_, err := dao.Insert(entity.GomaStringTypesEntity{
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
	dao := g.GomaDateTypes

	_, err := dao.Insert(entity.GomaDateTypesEntity{
		ID:               id,
		DateColumns:      time.Now(),
		DatetimeColumns:  time.Now(),
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
