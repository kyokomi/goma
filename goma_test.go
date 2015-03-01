package goma

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var numericTypeCacheSQL = `
update goma_numeric_types set
    id = /* id */64
,   tinyint_columns = /* tinyint_columns */1
,   bool_columns = /* bool_columns */1
,   smallint_columns = /* smallint_columns */1
,   mediumint_columns = /* mediumint_columns */1
,   int_columns = /* int_columns */1
,   integer_columns = /* integer_columns */1
,   serial_columns = /* serial_columns */64
,   decimal_columns = /* decimal_columns */'1111'
,   numeric_columns = /* numeric_columns */'1111'
,   float_columns = /* float_columns */32.1
,   double_columns = /* double_columns */64.1
 where
    id = /* id */1

`

var numericTypeQueryArgs = `
update goma_numeric_types set
    id = 123456789012345678
,   tinyint_columns = 8
,   bool_columns = 1
,   smallint_columns = 123
,   mediumint_columns = 256
,   int_columns = 11111111
,   integer_columns = 22222222
,   serial_columns = 1234567890
,   decimal_columns = '1234567890'
,   numeric_columns = '1234567890'
,   float_columns = 100000.234
,   double_columns = 1000000000.234
 where
    id = 123456789012345678

`

func TestQueryArgs(t *testing.T) {
	opt := Options{}
	opt.Driver = "mysql"
	opt.Debug = true
	opt.DBName = "goma_test"
	g, err := NewGoma(opt)
	if err != nil {
		t.Errorf("[ERROR] new goma: %s\nopts=%+v", err, opt)
	}

	// test data
	g.queryCache = make(map[tableName]map[queryName]string, 1)
	g.queryCache["goma_numeric_types"] = make(map[queryName]string, 1)
	g.queryCache["goma_numeric_types"]["update"] = numericTypeCacheSQL

	// test
	args := QueryArgs{
		"id":                int64(123456789012345678),
		"tinyint_columns":   int(8),
		"bool_columns":      int(1),
		"smallint_columns":  int(123),
		"mediumint_columns": int(256),
		"int_columns":       int(11111111),
		"integer_columns":   int(22222222),
		"serial_columns":    int64(1234567890),
		"decimal_columns":   "1234567890",
		"numeric_columns":   "1234567890",
		"float_columns":     float32(100000.234),
		"double_columns":    float64(1000000000.234),
	}

	res := g.QueryArgs("goma_numeric_types", "update", args)
	if res != numericTypeQueryArgs {
		t.Errorf("[ERROR] QueryArgs: %s\n != \n%s", res, numericTypeQueryArgs)
	}
}
