package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/postgres/entity"

	"github.com/kyokomi/goma"
)

// GomaNumericTypesDaoQueryer is interface
type GomaNumericTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaNumericTypesDao is generated goma_numeric_types table.
type GomaNumericTypesDao struct {
	*sql.DB
	TableName string
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
	tblDao.TableName = "GomaNumericTypes"
	return tblDao
}

// TxGomaNumericTypesDao is generated goma_numeric_types table transaction.
type TxGomaNumericTypesDao struct {
	*sql.Tx
	TableName string
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
	tblDao.TableName = "GomaNumericTypes"
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
	queryString := queryArgs("goma_numeric_types", "selectAll", nil)

	var es []entity.GomaNumericTypes
	rows, err := g.Query(queryString)
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
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_numeric_types", "selectByID", args)

	rows, err := g.Query(queryString)
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
	args := goma.QueryArgs{
		"id":               e.ID,
		"bool_columns":     e.BoolColumns,
		"smallint_columns": e.SmallintColumns,
		"int_columns":      e.IntColumns,
		"integer_columns":  e.IntegerColumns,

		"decimal_columns": e.DecimalColumns,
		"numeric_columns": e.NumericColumns,
		"float_columns":   e.FloatColumns,
	}
	queryString := queryArgs("goma_numeric_types", "insert", args)

	result, err := g.Exec(queryString)
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
	args := goma.QueryArgs{
		"id":               e.ID,
		"bool_columns":     e.BoolColumns,
		"smallint_columns": e.SmallintColumns,
		"int_columns":      e.IntColumns,
		"integer_columns":  e.IntegerColumns,
		"serial_columns":   e.SerialColumns,
		"decimal_columns":  e.DecimalColumns,
		"numeric_columns":  e.NumericColumns,
		"float_columns":    e.FloatColumns,
	}
	queryString := queryArgs("goma_numeric_types", "update", args)

	result, err := g.Exec(queryString)
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
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_numeric_types", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
