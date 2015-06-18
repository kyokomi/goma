package goma

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Goma is sql.DB access wrapper.
type Goma struct {
	*sql.DB
	options Options
}

// Open is create goma client.
// - database open
func Open(configPath string) (*sql.DB, error) {
	opts, err := NewOptions(configPath)
	if err != nil {
		return nil, err
	}
	return OpenOptions(opts)
}

// OpenOptions is create goma client.
// - database open
func OpenOptions(options Options) (*sql.DB, error) {
	return sql.Open(options.Driver, options.Source())
}

// Close sql.DB close.
func (d *Goma) Close() error {
	d.debugPrintln("goma close")

	err := d.DB.Close()

	return err
}

func GenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	if len(argsMap) <= 0 {
		return queryString, nil, nil
	}
	return sqlx.Named(queryString, argsMap)
}

// MySQLGenerateQuery generate bind args query
func MySQLGenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	return GenerateQuery(queryString, argsMap)
}

// PostgresGenerateQuery generate bind args query
func PostgresGenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	query, args, err := GenerateQuery(queryString, argsMap)
	return sqlx.Rebind(sqlx.DOLLAR, query), args, err
}

func (d *Goma) debugPrintln(v ...interface{}) {
	if d.options.Debug {
		fmt.Println(v...)
	}
}
