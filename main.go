package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	flag.StringVar(&driver, "d", "", "sql driver")
	flag.StringVar(&dataSource, "s", "", "sql data source")
	flag.StringVar(&out, "o", "", "output file")
	flag.StringVar(&file, "file", os.Getenv("GOFILE"), "input file")
	flag.StringVar(&pkg, "pkg", os.Getenv("GOPACKAGE"), "output package")
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

	// TODO: DBを読み込む いまはquestテーブル固定

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

	fmt.Println(buf.String())
}
