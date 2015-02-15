package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kyokomi/goma/goma"

	_ "github.com/go-sql-driver/mysql"
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

type DaoTemplateData struct {
	Name       string
	EntityName string
	Table      TableTemplateData
}

type TableTemplateData struct {
	Name       string
	MemberName string
}

//go:generate ego -package main templates

func main() {
	log.SetFlags(log.Llongfile)

	fmt.Println(driver, dataSource)

	// TODO: DBを読み込む いまはquestテーブル固定
	opts := goma.Options{
		Driver: "mysql",
		Source: "admin:password@tcp(localhost:3306)/test",
		DBName: "test",
		Debug:  true,
	}

	// TODO: xormのtool使ったほうが早そう
	/*
		// xorm reverse mysql root:@/test?charset=utf8 templates/goxorm
		Orm, err := xorm.NewEngine(args[0], args[1])
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		tables, err := Orm.DBMetas()
		if err != nil {
			log.Errorf("%v", err)
			return
		}
	*/

	g, err := goma.NewGoma(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	data := DaoTemplateData{
		Name:       "QuestDao",
		EntityName: "QuestEntity",
		Table: TableTemplateData{
			Name:       "Quest",
			MemberName: "quest",
		},
	}

	if err := os.Mkdir("sql", 0755); err != nil {
		log.Println("sql dir exsist")
	}

	if err := os.Mkdir("dao", 0755); err != nil {
		log.Println("dao dir exsist")
	}

	var buf bytes.Buffer
	DaoTemplate(&buf, data)

	if err := ioutil.WriteFile("dao/quest_gen.go", buf.Bytes(), 0644); err != nil {
		log.Fatalln(err)
	}
}
