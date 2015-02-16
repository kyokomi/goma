package goma

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Goma struct {
	*sql.DB
	options    Options
	queryCache map[TableName]map[QueryName]string
}

type QueryArgs map[string]interface{}

type TableName string
type QueryName string

type Options struct {
	Driver string // DriverName
	Source string // DataSource
	DBName string // DataBaseName
	Debug  bool
}

// TODO: あとで消す
const questSelectAll = `
select
	*
FROM
	quest
`

// TODO: あとで消す
const questSelectByID = `
select
  *
FROM
  quest
WHERE
  id = /* id */1
and
  name = /* name */"hoge"
`

func NewGoma(options Options) (*Goma, error) {

	var d Goma
	d.options = options

	db, err := sql.Open(options.Driver, options.Source)
	if err != nil {
		return nil, err
	}

	d.DB = db

	// sql下のディレクトリをtableNameとする
	// 各ディレクトリのファイル名 - .sqlをqueryNameとする

	// TODO: queryCacheする（一旦ハードコーディング）
	d.queryCache = make(map[TableName]map[QueryName]string, 1)
	d.queryCache["quest"] = make(map[QueryName]string, 2)
	d.queryCache["quest"]["selectAll"] = questSelectAll
	d.queryCache["quest"]["selectByID"] = questSelectByID

	return &d, nil
}

func (d *Goma) Close() error {
	d.Println("goma close")

	err := d.DB.Close()

	return err
}

func (d *Goma) QueryArgs(tableName TableName, queryName QueryName, args QueryArgs) string {
	cacheQuery := d.queryCache[tableName][queryName]
	if cacheQuery == "" {
		// TODO: ないときどうする? 再度読み込み?
	}
	return d.queryArgs(cacheQuery, args)
}

func (d *Goma) queryArgs(queryString string, args QueryArgs) string {
	if len(args) <= 0 {
		return queryString
	}

	d.Println("old: ", queryString)

	for key, val := range args {
		re := regexp.MustCompile(`\/\* ` + key + ` \*\/.*`)

		replaceWord := ""
		switch val.(type) {
		case int:
			replaceWord = strconv.Itoa(val.(int))
		case string:
			replaceWord = "\"" + val.(string) + "\""
		}
		queryString = re.ReplaceAllString(queryString, replaceWord)
	}

	d.Println("new: ", queryString)

	return queryString
}

func (d *Goma) Println(v ...interface{}) {
	if d.options.Debug {
		log.Println(v)
	}
}
