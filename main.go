package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/kyokomi/goma/goma"
	"github.com/kyokomi/goma/lint"
)

var (
	// command line flags
	driver   string
	dbName   string
	user     string
	password string
	host     string
	port     int

	debug      bool
	sqlRootDir string
	daoRootDir string

	// go generate default flags
	file string // input file (or directory)
	pkg  string // output package name
)

func init() {
	flag.StringVar(&driver, "driver", "mysql", "sql driver")
	flag.StringVar(&user, "user", "admin", "database access user's name")
	flag.StringVar(&password, "password", "password", "database access user's password")
	flag.StringVar(&host, "host", "localhost", "database host")
	flag.IntVar(&port, "port", 3306, "database port")
	flag.StringVar(&dbName, "db", "test", "database name")
	flag.BoolVar(&debug, "debug", false, "goma debug mode")
	flag.StringVar(&sqlRootDir, "sql", "sql", "generate sql root dir")
	flag.StringVar(&daoRootDir, "dao", "dao", "generate dao root dir")

	flag.StringVar(&file, "file", os.Getenv("GOFILE"), "input file")
	flag.StringVar(&pkg, "pkg", os.Getenv("GOPACKAGE"), "output package")
	flag.Parse()
}

var sampleDataMap = map[reflect.Type]string{
	reflect.TypeOf(int(1)):         "1",
	reflect.TypeOf(string("test")): "'test'",
	reflect.TypeOf(time.Now()):     "'2006/01/02 13:40:00'",
}

var driverImports = map[string]string{
	"mysql":    `_ "github.com/go-sql-driver/mysql"`,
	"postgres": `_ "github.com/lib/pq"`,
}

//go:generate ego -package main templates

func main() {
	log.SetFlags(log.Llongfile)

	fmt.Println(pkg, file)

	opt := goma.Options{}
	opt.Driver = driver
	opt.UserName = user
	opt.PassWord = password
	opt.Host = host
	opt.Port = port
	opt.DBName = dbName

	opt.Debug = debug
	opt.SQLRootDir = sqlRootDir
	opt.DaoRootDir = daoRootDir

	// xorm reverse mysql root:@/test?charset=utf8 templates/goxorm
	orm, err := xorm.NewEngine(driver, opt.Source())
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	tables, err := orm.DBMetas()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	helperData := HelperTemplateData{}
	helperData.PkgName = pkg
	helperData.DriverImport = driverImports[opt.Driver]
	helperData.Options = opt.Tuples()

	daoPkgName := strings.Replace(opt.DaoRootDir, `\./`, "", -1)
	daoPkgName = strings.Replace(opt.DaoRootDir, `\/`, "", -1)
	srcIdx := strings.Index(currentDir, "src/")
	helperData.DaoImport = filepath.Join(currentDir[srcIdx+len("src/"):], daoPkgName)
	helperData.DaoPkgName = daoPkgName

	var daoList []DaoTemplateData

	// sql, dao generate

	for _, table := range tables {
		// create templateData
		data := newTemplateData(table)

		// dao template
		daoRootPath := filepath.Join(currentDir, opt.DaoRootDir)
		if err := data.execDaoTemplate(daoRootPath); err != nil {
			log.Fatalln(err)
		}

		// sql template
		sqlRootPath := filepath.Join(currentDir, opt.SQLRootDir)
		if err := data.Table.execTableTemplate(sqlRootPath); err != nil {
			log.Fatalln(err)
		}

		daoList = append(daoList, data)
	}

	// helper generate

	helperData.DaoList = daoList

	if err := helperData.execHelperTemplate(currentDir); err != nil {
		log.Fatalln(err)
	}
}

func newTemplateData(table *core.Table) DaoTemplateData {
	imports := newImports(table.Columns())
	columns := newColumns(table.Columns())

	data := DaoTemplateData{}
	data.Name = lint.String(strings.Title(table.Name) + "Dao")
	data.MemberName = lint.String(table.Name)
	data.EntityName = lint.String(strings.Title(table.Name) + "Entity")
	data.Table = TableTemplateData{
		Name:      table.Name,
		TitleName: lint.String(strings.Title(table.Name)),
		Columns:   columns,
	}
	data.Imports = imports.slice()
	return data
}

func newImports(columns []*core.Column) set {
	importsMap := make(set, 0)
	for _, c := range columns {
		typ := core.SQLType2Type(c.SQLType)
		importsMap.add(typ.PkgPath())
	}
	return importsMap
}

func newColumns(columns []*core.Column) []ColumnTemplateData {
	var results []ColumnTemplateData
	for _, c := range columns {

		typ := core.SQLType2Type(c.SQLType)
		typeName := typ.Name()
		if typ.PkgPath() != "" {
			typeName = typ.PkgPath() + "." + typ.Name()
		}

		column := ColumnTemplateData{
			Name:         c.Name,
			TitleName:    lint.String(strings.Title(c.Name)),
			TypeName:     typeName,
			IsPrimaryKey: c.IsPrimaryKey,
			Sample:       sampleDataMap[typ],
		}
		results = append(results, column)
	}
	return results
}
