/*
Copyright (c) 2014 Naoya Inada <naoina@kuune.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package dialect

import (
	"fmt"
	"strings"
)

type MySQL struct {
}

func (d *MySQL) ColumnType(name string, size uint64, autoIncrement bool) (typ string, null bool) {
	switch name {
	case "string":
		return d.varchar(size), false
	case "sql.NullString", "*string":
		return d.varchar(size), true
	case "int", "int32":
		return "INT", false
	case "*int", "*int32":
		return "INT", true
	case "int8":
		return "TINYINT", false
	case "*int8":
		return "TINYINT", true
	case "bool":
		return "TINYINT", false
	case "*bool", "sql.NullBool":
		return "TINYINT", true
	case "int16":
		return "SMALLINT", false
	case "*int16":
		return "SMALLINT", true
	case "int64":
		return "BIGINT", false
	case "sql.NullInt64", "*int64":
		return "BIGINT", true
	case "uint", "uint32":
		return "INT UNSIGNED", false
	case "*uint", "*uint32":
		return "INT UNSIGNED", true
	case "uint8":
		return "TINYINT UNSIGNED", false
	case "*uint8":
		return "TINYINT UNSIGNED", true
	case "uint16":
		return "SMALLINT UNSIGNED", false
	case "*uint16":
		return "SMALLINT UNSIGNED", true
	case "uint64":
		return "BIGINT UNSIGNED", false
	case "*uint64":
		return "BIGINT UNSIGNED", true
	case "float32", "float64":
		return "DOUBLE", false
	case "sql.NullFloat64", "*float32", "*float64":
		return "DOUBLE", true
	case "time.Time":
		return "DATETIME", false
	case "*time.Time":
		return "DATETIME", true
	default:
		return "VARCHAR(255)", true
	}
}

func (d *MySQL) Quote(s string) string {
	return fmt.Sprintf("`%s`", strings.Replace(s, "`", "``", -1))
}

func (d *MySQL) QuoteString(s string) string {
	return fmt.Sprintf("'%s'", strings.Replace(s, "'", "''", -1))
}

func (d *MySQL) AutoIncrement() string {
	return "AUTO_INCREMENT"
}

func (d *MySQL) varchar(size uint64) string {
	if size == 0 {
		size = 255 // default.
	}
	switch {
	case size < 21846:
		return fmt.Sprintf("VARCHAR(%d)", size)
	case size < (1<<16)-1-2: // approximate 64KB.
		// 65533 ((2^16) - 1) - (length of prefix)
		// See http://dev.mysql.com/doc/refman/5.5/en/string-type-overview.html#idm140418628949072
		return "TEXT"
	case size < 1<<24: // 16MB.
		return "MEDIUMTEXT"
	}
	return "LONGTEXT"
}
