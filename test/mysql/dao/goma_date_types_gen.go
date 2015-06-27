package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/mysql/entity"
)

var tableGomaDateTypes = "goma_date_types"
var columnsGomaDateTypes = []string{
	"id",
	"datetime_columns",
	"timestamp_columns",
	"nil_datetime_columns",
}

// GomaDateTypesTableName goma_date_types table name
func GomaDateTypesTableName() string {
	return tableGomaDateTypes
}

// GomaDateTypesTableColumns goma_date_types table columns
func GomaDateTypesTableColumns() []string {
	return columnsGomaDateTypes
}

// GomaDateTypesDaoQueryer is interface
type GomaDateTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaDateTypesDao is generated goma_date_types table.
type GomaDateTypesDao struct {
	*sql.DB
	TableName string
	Columns   []string
}

// Query GomaDateTypesDao query
func (g GomaDateTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec GomaDateTypesDao exec
func (g GomaDateTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ GomaDateTypesDaoQueryer = (*GomaDateTypesDao)(nil)

// GomaDateTypes is GomaDateTypesDao.
func GomaDateTypes(db *sql.DB) GomaDateTypesDao {
	tblDao := GomaDateTypesDao{}
	tblDao.DB = db
	tblDao.TableName = tableGomaDateTypes
	tblDao.Columns = columnsGomaDateTypes
	return tblDao
}

// TxGomaDateTypesDao is generated goma_date_types table transaction.
type TxGomaDateTypesDao struct {
	*sql.Tx
	TableName string
	Columns   []string
}

// Query TxGomaDateTypesDao query
func (g TxGomaDateTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxGomaDateTypesDao exec
func (g TxGomaDateTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ GomaDateTypesDaoQueryer = (*TxGomaDateTypesDao)(nil)

// TxGomaDateTypes is GomaDateTypesDao.
func TxGomaDateTypes(tx *sql.Tx) TxGomaDateTypesDao {
	tblDao := TxGomaDateTypesDao{}
	tblDao.Tx = tx
	tblDao.TableName = tableGomaDateTypes
	tblDao.Columns = columnsGomaDateTypes
	return tblDao
}

// SelectAll select goma_date_types table all recode.
func (g GomaDateTypesDao) SelectAll() ([]entity.GomaDateTypes, error) {
	return _GomaDateTypesSelectAll(g)
}

// SelectAll transaction select goma_date_types table all recode.
func (g TxGomaDateTypesDao) SelectAll() ([]entity.GomaDateTypes, error) {
	return _GomaDateTypesSelectAll(g)
}

func _GomaDateTypesSelectAll(g GomaDateTypesDaoQueryer) ([]entity.GomaDateTypes, error) {
	queryString, args, err := queryArgs("goma_date_types", "selectAll", nil)
	if err != nil {
		return nil, err
	}

	var es []entity.GomaDateTypes
	rows, err := g.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.GomaDateTypes
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

// SelectByID select goma_date_types table by primaryKey.
func (g GomaDateTypesDao) SelectByID(id int64) (entity.GomaDateTypes, error) {
	return _GomaDateTypesSelectByID(g, id)
}

// SelectByID transaction select goma_date_types table by primaryKey.
func (g TxGomaDateTypesDao) SelectByID(id int64) (entity.GomaDateTypes, error) {
	return _GomaDateTypesSelectByID(g, id)
}

func _GomaDateTypesSelectByID(g GomaDateTypesDaoQueryer, id int64) (entity.GomaDateTypes, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("goma_date_types", "selectByID", argsMap)
	if err != nil {
		return entity.GomaDateTypes{}, err
	}

	rows, err := g.Query(queryString, args...)
	if err != nil {
		return entity.GomaDateTypes{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaDateTypes{}, sql.ErrNoRows
	}

	var e entity.GomaDateTypes
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaDateTypes{}, err
	}

	return e, nil
}

// Insert insert goma_date_types table.
func (g GomaDateTypesDao) Insert(e entity.GomaDateTypes) (sql.Result, error) {
	return _GomaDateTypesInsert(g, e)
}

// Insert transaction insert goma_date_types table.
func (g TxGomaDateTypesDao) Insert(e entity.GomaDateTypes) (sql.Result, error) {
	return _GomaDateTypesInsert(g, e)
}

func _GomaDateTypesInsert(g GomaDateTypesDaoQueryer, e entity.GomaDateTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":                   e.ID,
		"datetime_columns":     e.DatetimeColumns,
		"timestamp_columns":    e.TimestampColumns,
		"nil_datetime_columns": e.NilDatetimeColumns,
	}
	queryString, args, err := queryArgs("goma_date_types", "insert", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_date_types table.
func (g GomaDateTypesDao) Update(e entity.GomaDateTypes) (sql.Result, error) {
	return _GomaDateTypesUpdate(g, e)
}

// Update transaction update goma_date_types table.
func (g TxGomaDateTypesDao) Update(e entity.GomaDateTypes) (sql.Result, error) {
	return _GomaDateTypesUpdate(g, e)
}

// Update update goma_date_types table.
func _GomaDateTypesUpdate(g GomaDateTypesDaoQueryer, e entity.GomaDateTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":                   e.ID,
		"datetime_columns":     e.DatetimeColumns,
		"timestamp_columns":    e.TimestampColumns,
		"nil_datetime_columns": e.NilDatetimeColumns,
	}
	queryString, args, err := queryArgs("goma_date_types", "update", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_date_types table.
func (g GomaDateTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaDateTypesDelete(g, id)
}

// Delete transaction delete goma_date_types table.
func (g TxGomaDateTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaDateTypesDelete(g, id)
}

// Delete delete goma_date_types table by primaryKey.
func _GomaDateTypesDelete(g GomaDateTypesDaoQueryer, id int64) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("goma_date_types", "delete", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
