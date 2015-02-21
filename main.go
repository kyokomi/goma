package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"strings"

	"go/format"

	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kyokomi/goma/lint"
)

var (
	// command line flags
	driver     string
	dataSource string
	out        string // output file
	file       string // input file (or directory)
	pkg        string // output package name
)

func init() {
	flag.StringVar(&driver, "driver", "", "sql driver")
	flag.StringVar(&dataSource, "source", "", "sql data source")
	flag.StringVar(&out, "o", "", "output file")
	flag.StringVar(&file, "file", os.Getenv("GOFILE"), "input file")
	flag.StringVar(&pkg, "pkg", os.Getenv("GOPACKAGE"), "output package")
	flag.Parse()
}

var sampleDataMap = map[reflect.Type]string{
	reflect.TypeOf(int(1)):         "1",
	reflect.TypeOf(string("test")): "'test'",
	reflect.TypeOf(time.Now()):     "'2006/01/02 13:40:00'",
}

//go:generate ego -package main templates

func main() {
	log.SetFlags(log.Llongfile)

	// xorm reverse mysql root:@/test?charset=utf8 templates/goxorm
	orm, err := xorm.NewEngine(driver, dataSource)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	tables, err := orm.DBMetas()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	if err := os.Mkdir("sql", 0755); err != nil {
		debugPrintln("sql dir exist")
	}

	if err := os.Mkdir("dao", 0755); err != nil {
		debugPrintln("dao dir exist")
	}

	for _, table := range tables {
		debugPrintln("start ", table.Name)

		importsMap := make(set, 0)

		var columns []ColumnTemplateData
		for _, c := range table.Columns() {

			typ := core.SQLType2Type(c.SQLType)
			importsMap.add(typ.PkgPath())

			typeName := typ.Name()
			if typ.PkgPath() != "" {
				typeName = typ.PkgPath() + "." + typ.Name()
			}

			sampleData := sampleDataMap[typ]

			column := ColumnTemplateData{
				Name:         c.Name,
				TitleName:    lint.String(strings.Title(c.Name)),
				TypeName:     typeName,
				IsPrimaryKey: c.IsPrimaryKey,
				Sample:       sampleData,
			}
			columns = append(columns, column)
		}

		data := DaoTemplateData{
			Name:       lint.String(strings.Title(table.Name) + "Dao"),
			MemberName: lint.String(table.Name),
			EntityName: lint.String(strings.Title(table.Name) + "Entity"),
			Table: TableTemplateData{
				Name:      table.Name,
				TitleName: lint.String(strings.Title(table.Name)),
				Columns:   columns,
			},
			Imports: importsMap.slice(),
		}

		var buf bytes.Buffer
		if err := DaoTemplate(&buf, data); err != nil {
			log.Fatalln(err)
		} else {

			bts, err := format.Source(buf.Bytes())
			if err != nil {
				log.Fatalln(err, buf.String())
			}
			if err := ioutil.WriteFile("dao/"+table.Name+"_gen.go", bts, 0644); err != nil {
				log.Fatalln(err)
			}
		}

		if err := os.Mkdir("sql/"+table.Name, 0755); err != nil {
			debugPrintln("sql/" + table.Name + " dir exsist")
		}

		buf.Reset()

		if err := SelectAllTemplate(&buf, data.Table); err != nil {
			log.Fatalln(err)
		} else {
			if err := ioutil.WriteFile("sql/"+table.Name+"/selectAll.sql", buf.Bytes(), 0644); err != nil {
				log.Fatalln(err)
			}
		}

		buf.Reset()

		if err := SelectByIDTemplate(&buf, data.Table); err != nil {
			log.Fatalln(err)
		} else {
			if err := ioutil.WriteFile("sql/"+table.Name+"/selectByID.sql", buf.Bytes(), 0644); err != nil {
				log.Fatalln(err)
			}
		}

		buf.Reset()

		if err := InsertTemplate(&buf, data.Table); err != nil {
			log.Fatalln(err)
		} else {
			if err := ioutil.WriteFile("sql/"+table.Name+"/insert.sql", buf.Bytes(), 0644); err != nil {
				log.Fatalln(err)
			}
		}

		buf.Reset()

		if err := UpdateTemplate(&buf, data.Table); err != nil {
			log.Fatalln(err)
		} else {
			if err := ioutil.WriteFile("sql/"+table.Name+"/update.sql", buf.Bytes(), 0644); err != nil {
				log.Fatalln(err)
			}
		}

		buf.Reset()

		if err := DeleteTemplate(&buf, data.Table); err != nil {
			log.Fatalln(err)
		} else {
			if err := ioutil.WriteFile("sql/"+table.Name+"/delete.sql", buf.Bytes(), 0644); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func debugPrintln(v ...interface{}) {
	// SliceInsert (https://code.google.com/p/go-wiki/wiki/SliceTricks)
	v = append(v[:0], append([]interface{}{"[goma]", "[debug]"}, v[0:]...)...)

	fmt.Println(v...)
}
