package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strings"

	"go/format"

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

type ColumnTemplateData struct {
	Name         string
	TitleName    string
	TypeName     string
	IsPrimaryKey bool
}

type DaoTemplateData struct {
	Name       string
	MemberName string
	EntityName string
	Table      TableTemplateData
	Imports    []string
}

type TableTemplateData struct {
	Name      string
	TitleName string
	Columns   []ColumnTemplateData
}

type Set map[string]bool

func (s Set) Add(key string) {
	if key != "" {
		s[key] = true
	}
}

func (s Set) Slice() []string {
	var keys []string
	for key, val := range s {
		if val {
			keys = append(keys, key)
		}
	}
	return keys
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
		log.Println("sql dir exsist")
	}

	if err := os.Mkdir("dao", 0755); err != nil {
		log.Println("dao dir exsist")
	}

	for _, table := range tables {
		fmt.Println("========== ", table.Name, " ==========")

		importsMap := make(Set, 0)

		columns := make([]ColumnTemplateData, 0)
		for _, c := range table.Columns() {

			typ := core.SQLType2Type(c.SQLType)
			importsMap.Add(typ.PkgPath())

			typeName := typ.Name()
			if typ.PkgPath() != "" {
				typeName = typ.PkgPath() + "." + typ.Name()
			}

			if c.SQLType.IsTime() {
				typeName = "*" + typeName
			}

			column := ColumnTemplateData{
				Name:         c.Name,
				TitleName:    lint.String(strings.Title(c.Name)),
				TypeName:     typeName,
				IsPrimaryKey: c.IsPrimaryKey,
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
			Imports: importsMap.Slice(),
		}

		var buf bytes.Buffer
		if err := DaoTemplate(&buf, data); err != nil {
			log.Fatalln(err)
		} else {

			bts, err := format.Source(buf.Bytes())
			if err != nil {
				log.Fatalln(err)
			}
			if err := ioutil.WriteFile("dao/"+table.Name+"_gen.go", bts, 0644); err != nil {
				log.Fatalln(err)
			}
		}

		if err := os.Mkdir("sql/"+table.Name, 0755); err != nil {
			log.Println("sql/" + table.Name + " dir exsist")
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
	}
}
