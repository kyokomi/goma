package main

import (
	"fmt"
	"log"

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
}

func numericTest(goma Goma) {

	id := int64(1234567890)

	// numeric

	_, err := goma.GomaNumericTypes.Insert(entity.GomaNumericTypesEntity{
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

	if numericTbl, err := goma.GomaNumericTypes.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", goma.GomaNumericTypes.TableName, numericTbl)
	}

	_, err = goma.GomaNumericTypes.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if numericTbl, err := goma.GomaNumericTypes.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", goma.GomaNumericTypes.TableName, numericTbl)
	}
}

func stringTest(goma Goma) {

	id := int64(1234567890)

	// string

	_, err := goma.GomaStringTypes.Insert(entity.GomaStringTypesEntity{
		ID:             id,
		Text:           "あいうえおかきくけこ",
		CharColumns:    "a",
		VarcharColumns: "1234567890abcdefghijkelmnopqrstuvwxyz",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if numericTbl, err := goma.GomaStringTypes.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", goma.GomaStringTypes.TableName, numericTbl)
	}

	_, err = goma.GomaStringTypes.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	if numericTbl, err := goma.GomaStringTypes.SelectByID(id); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s: %+v\n", goma.GomaStringTypes.TableName, numericTbl)
	}

}
