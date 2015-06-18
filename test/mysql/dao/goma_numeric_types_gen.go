package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/mysql/entity"
)

var tableGomaNumericTypes = "goma_numeric_types"
var columnsGomaNumericTypes = []string{
	"id",
	"tinyint_columns",
	"bool_columns",
	"smallint_columns",
	"mediumint_columns",
	"int_columns",
	"integer_columns",
	"serial_columns",
	"decimal_columns",
	"numeric_columns",
	"float_columns",
	"double_columns",
}

// GomaNumericTypesTableName goma_numeric_types table name
func GomaNumericTypesTableName() string {
	return tableGomaNumericTypes
}

// GomaNumericTypesTableColumns goma_numeric_types table columns
func GomaNumericTypesTableColumns() []string {
	return columnsGomaNumericTypes
}

// GomaNumericTypesDaoQueryer is interface
type GomaNumericTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaNumericTypesDao is generated goma_numeric_types table.
type GomaNumericTypesDao struct {
	*sql.DB
	TableName string
	Columns   []string
}

// Query GomaNumericTypesDao query
func (g GomaNumericTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec GomaNumericTypesDao exec
func (g GomaNumericTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ GomaNumericTypesDaoQueryer = (*GomaNumericTypesDao)(nil)

// GomaNumericTypes is GomaNumericTypesDao.
func GomaNumericTypes(db *sql.DB) GomaNumericTypesDao {
	tblDao := GomaNumericTypesDao{}
	tblDao.DB = db
	tblDao.TableName = tableGomaNumericTypes
	tblDao.Columns = columnsGomaNumericTypes
	return tblDao
}

// TxGomaNumericTypesDao is generated goma_numeric_types table transaction.
type TxGomaNumericTypesDao struct {
	*sql.Tx
	TableName string
	Columns   []string
}

// Query TxGomaNumericTypesDao query
func (g TxGomaNumericTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxGomaNumericTypesDao exec
func (g TxGomaNumericTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ GomaNumericTypesDaoQueryer = (*TxGomaNumericTypesDao)(nil)

// TxGomaNumericTypes is GomaNumericTypesDao.
func TxGomaNumericTypes(tx *sql.Tx) TxGomaNumericTypesDao {
	tblDao := TxGomaNumericTypesDao{}
	tblDao.Tx = tx
	tblDao.TableName = tableGomaNumericTypes
	tblDao.Columns = columnsGomaNumericTypes
	return tblDao
}

// SelectAll select goma_numeric_types table all recode.
func (g GomaNumericTypesDao) SelectAll() ([]entity.GomaNumericTypes, error) {
	return _GomaNumericTypesSelectAll(g)
}

// SelectAll transaction select goma_numeric_types table all recode.
func (g TxGomaNumericTypesDao) SelectAll() ([]entity.GomaNumericTypes, error) {
	return _GomaNumericTypesSelectAll(g)
}

func _GomaNumericTypesSelectAll(g GomaNumericTypesDaoQueryer) ([]entity.GomaNumericTypes, error) {
	queryString, args, err := queryArgs("goma_numeric_types", "selectAll", nil)
	if err != nil {
		return nil, err
	}

	var es []entity.GomaNumericTypes
	rows, err := g.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.GomaNumericTypes
		if err := e.Scan(rows); err != nil {
			break
		}

		es = append(es, e)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return es, nil
}

// SelectByID select goma_numeric_types table by primaryKey.
func (g GomaNumericTypesDao) SelectByID(id int64) (entity.GomaNumericTypes, error) {
	return _GomaNumericTypesSelectByID(g, id)
}

// SelectByID transaction select goma_numeric_types table by primaryKey.
func (g TxGomaNumericTypesDao) SelectByID(id int64) (entity.GomaNumericTypes, error) {
	return _GomaNumericTypesSelectByID(g, id)
}

func _GomaNumericTypesSelectByID(g GomaNumericTypesDaoQueryer, id int64) (entity.GomaNumericTypes, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("goma_numeric_types", "selectByID", argsMap)
	if err != nil {
		return entity.GomaNumericTypes{}, err
	}

	rows, err := g.Query(queryString, args...)
	if err != nil {
		return entity.GomaNumericTypes{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaNumericTypes{}, sql.ErrNoRows
	}

	var e entity.GomaNumericTypes
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaNumericTypes{}, err
	}

	return e, nil
}

// Insert insert goma_numeric_types table.
func (g GomaNumericTypesDao) Insert(e entity.GomaNumericTypes) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, e)
}

// Insert transaction insert goma_numeric_types table.
func (g TxGomaNumericTypesDao) Insert(e entity.GomaNumericTypes) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, e)
}

func _GomaNumericTypesInsert(g GomaNumericTypesDaoQueryer, e entity.GomaNumericTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":                e.ID,
		"tinyint_columns":   e.TinyintColumns,
		"bool_columns":      e.BoolColumns,
		"smallint_columns":  e.SmallintColumns,
		"mediumint_columns": e.MediumintColumns,
		"int_columns":       e.IntColumns,
		"integer_columns":   e.IntegerColumns,

		"decimal_columns": e.DecimalColumns,
		"numeric_columns": e.NumericColumns,
		"float_columns":   e.FloatColumns,
		"double_columns":  e.DoubleColumns,
	}
	queryString, args, err := queryArgs("goma_numeric_types", "insert", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (g GomaNumericTypesDao) Update(e entity.GomaNumericTypes) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, e)
}

// Update transaction update goma_numeric_types table.
func (g TxGomaNumericTypesDao) Update(e entity.GomaNumericTypes) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, e)
}

// Update update goma_numeric_types table.
func _GomaNumericTypesUpdate(g GomaNumericTypesDaoQueryer, e entity.GomaNumericTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":                e.ID,
		"tinyint_columns":   e.TinyintColumns,
		"bool_columns":      e.BoolColumns,
		"smallint_columns":  e.SmallintColumns,
		"mediumint_columns": e.MediumintColumns,
		"int_columns":       e.IntColumns,
		"integer_columns":   e.IntegerColumns,
		"serial_columns":    e.SerialColumns,
		"decimal_columns":   e.DecimalColumns,
		"numeric_columns":   e.NumericColumns,
		"float_columns":     e.FloatColumns,
		"double_columns":    e.DoubleColumns,
	}
	queryString, args, err := queryArgs("goma_numeric_types", "update", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_numeric_types table.
func (g GomaNumericTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaNumericTypesDelete(g, id)
}

// Delete transaction delete goma_numeric_types table.
func (g TxGomaNumericTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaNumericTypesDelete(g, id)
}

// Delete delete goma_numeric_types table by primaryKey.
func _GomaNumericTypesDelete(g GomaNumericTypesDaoQueryer, id int64) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("goma_numeric_types", "delete", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
