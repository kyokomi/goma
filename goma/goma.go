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
    options Options
}

type QueryArgs map[string]interface{}

type Options struct {
    Debug bool
}

func NewGoma(options Options) (*Goma, error) {

	var d Goma
	d.options = options

	// TODO: 引数でもらうようにする
	db, err := sql.Open("mysql", "admin:password@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}

	d.DB = db

	return &d, nil
}

func (d *Goma) Close() error {
	d.Println("goma close")

	err := d.DB.Close()

	return err
}

func (d *Goma) QueryArgs(queryString string, args QueryArgs) string {

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
