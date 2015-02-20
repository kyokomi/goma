package goma

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"

	"io/ioutil"

	"strings"
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

func NewGoma(options Options) (*Goma, error) {

	var d Goma
	d.options = options

	db, err := sql.Open(options.Driver, options.Source)
	if err != nil {
		return nil, err
	}

	d.DB = db

	// TODO: startディレクトリOptionsでもらう？
	// sql下のディレクトリをtableNameとする
	// 各ディレクトリのファイル名 - .sqlをqueryNameとする
	dirs, err := ioutil.ReadDir("./sql")
	if err != nil {
		return nil, err
	}

	d.queryCache = make(map[TableName]map[QueryName]string, len(dirs))
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		fileInfos, err := ioutil.ReadDir("./sql/" + dir.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		tableName := TableName(dir.Name())

		d.queryCache[tableName] = make(map[QueryName]string, len(fileInfos))

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				continue
			}

			f, err := ioutil.ReadFile("./sql/" + dir.Name() + "/" + fileInfo.Name())
			if err != nil {
				log.Println(err)
				continue
			}

			queryName := QueryName(strings.Replace(fileInfo.Name(), ".sql", "", -1))
			d.queryCache[tableName][queryName] = string(f)
		}
	}

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
