package goma

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"
	"time"

	"io/ioutil"

	"fmt"
	"path/filepath"
	"strings"
)

// Goma is sql.DB access wrapper.
type Goma struct {
	*sql.DB
	options    Options
	queryCache map[tableName]map[queryName]string
}

// QueryArgs sql query args
type QueryArgs map[string]interface{}

type tableName string
type queryName string

// NewGoma is create goma client.
// - database open
// - query local memory cache
func NewGoma(options Options) (*Goma, error) {

	var d Goma
	d.options = options

	d.debugPrintln(options.Source())

	db, err := sql.Open(options.Driver, options.Source())
	if err != nil {
		return nil, err
	}

	d.DB = db
	d.cacheQuery()

	return &d, nil
}

func (d *Goma) cacheQuery() error {
	sqlRootDir := d.options.SQLRootDir

	// sql下のディレクトリをtableNameとする
	// 各ディレクトリのファイル名 - .sqlをqueryNameとする
	dirs, err := ioutil.ReadDir(sqlRootDir)
	if err != nil {
		return err
	}

	d.queryCache = make(map[tableName]map[queryName]string, len(dirs))
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		dirPath := filepath.Join(sqlRootDir, dir.Name())
		fileInfos, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Println(err)
			continue
		}

		tableName := tableName(dir.Name())

		d.queryCache[tableName] = make(map[queryName]string, len(fileInfos))

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				continue
			}

			filePath := filepath.Join(sqlRootDir, dir.Name(), fileInfo.Name())
			f, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Println(err)
				continue
			}

			// queryName.sqlのqueryNameを抽出
			queryName := queryName(strings.Replace(fileInfo.Name(), ".sql", "", -1))
			d.queryCache[tableName][queryName] = string(f)

			d.debugPrintln("cache ok: ", filePath)
			d.debugPrintln(d.queryCache[tableName][queryName])
		}
	}

	return nil
}

// Close sql.DB close.
func (d *Goma) Close() error {
	d.debugPrintln("goma close")

	err := d.DB.Close()

	return err
}

// QueryArgs create sql query in the local cache.
func (d *Goma) QueryArgs(tableName tableName, queryName queryName, args QueryArgs) string {
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

	for key, val := range args {
		re := regexp.MustCompile(`\/\* ` + key + ` \*\/.*`)

		replaceWord := ""
		switch val.(type) {
		case int:
			replaceWord = strconv.Itoa(val.(int))
		case string:
			replaceWord = "'" + val.(string) + "'"
		case time.Time:
			replaceWord = "'" + val.(time.Time).Format("2006-01-02 15:04:05") + "'"
		}
		queryString = re.ReplaceAllString(queryString, replaceWord)
	}

	return queryString
}

func (d *Goma) debugPrintln(v ...interface{}) {
	if d.options.Debug {
		fmt.Println(v...)
	}
}
