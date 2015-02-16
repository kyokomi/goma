package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kyokomi/goma/goma"

	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
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
	EntityName string
	Table      TableTemplateData
}

type TableTemplateData struct {
	Name      string
	TitleName string
	Columns   []ColumnTemplateData
}

//go:generate ego -package main templates

func main() {
	log.SetFlags(log.Llongfile)

	fmt.Println(driver, dataSource)

	opts := goma.Options{
		Driver: "mysql",
		Source: "admin:password@tcp(localhost:3306)/test",
		DBName: "test",
		Debug:  true,
	}

	// TODO: xorm reverse mysql root:@/test?charset=utf8 templates/goxorm
	orm, err := xorm.NewEngine(opts.Driver, opts.Source)
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

		columns := make([]ColumnTemplateData, 0)
		for _, c := range table.Columns() {
			column := ColumnTemplateData{
				Name:         c.Name,
				TitleName:    strings.Title(c.Name), // TODO: golintする
				TypeName:     core.SQLType2Type(c.SQLType).Name(),
				IsPrimaryKey: c.IsPrimaryKey,
			}
			columns = append(columns, column)
		}

		data := DaoTemplateData{
			Name:       strings.Title(table.Name) + "Dao",
			EntityName: strings.Title(table.Name) + "Entity",
			Table: TableTemplateData{
				Name:      table.Name,
				TitleName: strings.Title(table.Name),
				Columns:   columns,
			},
		}

		var buf bytes.Buffer
		if err := DaoTemplate(&buf, data); err != nil {
			log.Fatalln(err)
		} else {
			// TODO: gofmtする
			if err := ioutil.WriteFile("dao/"+table.Name+"_gen.go", buf.Bytes(), 0644); err != nil {
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
