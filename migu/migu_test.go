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

package migu

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	// TODO: dbName
	db, err = sql.Open("mysql", "travis@/migu_test")
	if err != nil {
		panic(err)
	}
}

func before(t *testing.T) {
	if _, err := db.Exec(`DROP TABLE IF EXISTS user`); err != nil {
		t.Fatal(err)
	}
}

func TestDiffWithSrc(t *testing.T) {
	before(t)
	types := map[string]string{
		"int":             "INT NOT NULL",
		"int8":            "TINYINT NOT NULL",
		"int16":           "SMALLINT NOT NULL",
		"int32":           "INT NOT NULL",
		"int64":           "BIGINT NOT NULL",
		"*int":            "INT",
		"*int8":           "TINYINT",
		"*int16":          "SMALLINT",
		"*int32":          "INT",
		"*int64":          "BIGINT",
		"uint":            "INT UNSIGNED NOT NULL",
		"uint8":           "TINYINT UNSIGNED NOT NULL",
		"uint16":          "SMALLINT UNSIGNED NOT NULL",
		"uint32":          "INT UNSIGNED NOT NULL",
		"uint64":          "BIGINT UNSIGNED NOT NULL",
		"*uint":           "INT UNSIGNED",
		"*uint8":          "TINYINT UNSIGNED",
		"*uint16":         "SMALLINT UNSIGNED",
		"*uint32":         "INT UNSIGNED",
		"*uint64":         "BIGINT UNSIGNED",
		"sql.NullInt64":   "BIGINT",
		"string":          "VARCHAR(255) NOT NULL",
		"*string":         "VARCHAR(255)",
		"sql.NullString":  "VARCHAR(255)",
		"bool":            "TINYINT NOT NULL",
		"*bool":           "TINYINT",
		"sql.NullBool":    "TINYINT",
		"float32":         "DOUBLE NOT NULL",
		"float64":         "DOUBLE NOT NULL",
		"*float32":        "DOUBLE",
		"*float64":        "DOUBLE",
		"sql.NullFloat64": "DOUBLE",
		"time.Time":       "DATETIME NOT NULL",
		"*time.Time":      "DATETIME",
	}
	for t1, s1 := range types {
		for t2, s2 := range types {
			if err := testDiffWithSrc(t, t1, s1, t2, s2); err != nil {
				t.Error(err)
				continue
			}
		}
	}
}

func testDiffWithSrc(t *testing.T, t1, s1, t2, s2 string) error {
	src := fmt.Sprintf("package migu_test\n"+
		"type User struct {\n"+
		"	A %s\n"+
		"}", t1)
	results, err := Diff(db, "", src)
	if err != nil {
		return err
	}
	actual := results
	expect := []string{
		fmt.Sprintf("CREATE TABLE `user` (\n"+
			"  `a` %s\n"+
			")", s1),
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf(`create: migu.Diff(db, "", %q) => %#v, nil; want %#v, nil`, src, actual, expect)
	}
	for _, s := range actual {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	defer db.Exec("DROP TABLE IF EXISTS `user`")

	results, err = Diff(db, "", src)
	if err != nil {
		return err
	}
	actual = results
	expect = []string(nil)
	if !reflect.DeepEqual(actual, expect) {
		return fmt.Errorf(`before: migu.Diff(db, "", %q) => %#v, nil; want %#v, nil`, src, actual, expect)
	}

	src = fmt.Sprintf("package migu_test\n"+
		"type User struct {\n"+
		"	A %s\n"+
		"}", t2)
	results, err = Diff(db, "", src)
	if err != nil {
		return err
	}
	actual = results
	if s1 == s2 {
		expect = []string(nil)
	} else {
		expect = []string{"ALTER TABLE `user` MODIFY `a` " + s2}
	}
	sort.Strings(actual)
	sort.Strings(expect)
	if !reflect.DeepEqual(actual, expect) {
		return fmt.Errorf(`diff: %s to %s; migu.Diff(db, "", %q) => %#v, nil; want %#v, nil`, t1, t2, src, actual, expect)
	}
	for _, s := range actual {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}

	src = "package migu_test"
	results, err = Diff(db, "", src)
	if err != nil {
		return err
	}
	actual = results
	expect = []string{"DROP TABLE `user`"}
	sort.Strings(actual)
	sort.Strings(expect)
	if !reflect.DeepEqual(actual, expect) {
		return fmt.Errorf(`drop: migu.Diff(db, "", %q) => %#v, nil; want %#v, nil`, src, actual, expect)
	}
	for _, s := range actual {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	return nil
}

func TestFprint(t *testing.T) {
	for _, v := range []struct {
		sqls   []string
		expect string
	}{
		{[]string{`
CREATE TABLE user (
  name VARCHAR(255)
)`}, `type User struct {
	Name *string
}

`},
		{[]string{`
CREATE TABLE user (
  name VARCHAR(255),
  age  INT
)`}, `type User struct {
	Name *string
	Age  *int
}

`},
		{[]string{`
CREATE TABLE user (
  name VARCHAR(255)
)`, `

CREATE TABLE post (
  title   VARCHAR(255),
  content VARCHAR(65533)
)`}, `type Post struct {
	Title   *string
	Content *string
}

type User struct {
	Name *string
}

`},
	} {
		v := v
		func() {
			for _, sql := range v.sqls {
				if _, err := db.Exec(sql); err != nil {
					t.Error(err)
					return
				}
			}
			defer func() {
				for _, v := range []string{
					`DROP TABLE IF EXISTS user`,
					`DROP TABLE IF EXISTS post`,
				} {
					if _, err := db.Exec(v); err != nil {
						t.Fatal(err)
					}
				}
			}()
			var buf bytes.Buffer
			if err := Fprint(&buf, db); err != nil {
				t.Error(err)
				return
			}
			actual := buf.String()
			expect := v.expect
			if !reflect.DeepEqual(actual, expect) {
				t.Errorf(`migu.Fprint(buf, db); buf => %#v; want %#v`, actual, expect)
			}
		}()
	}
}
