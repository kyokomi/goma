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

// GomaStringTypesDaoQueryer is interface
type GomaStringTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaStringTypesDao is generated goma_string_types table.
type GomaStringTypesDao struct {
	*sql.DB
	TableName string
}

// Query GomaStringTypesDao query
func (g GomaStringTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec GomaStringTypesDao exec
func (g GomaStringTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ GomaStringTypesDaoQueryer = (*GomaStringTypesDao)(nil)

// GomaStringTypes is GomaStringTypesDao.
func GomaStringTypes(db *sql.DB) GomaStringTypesDao {
	tblDao := GomaStringTypesDao{}
	tblDao.DB = db
	tblDao.TableName = "GomaStringTypes"
	return tblDao
}

// TxGomaStringTypesDao is generated goma_string_types table transaction.
type TxGomaStringTypesDao struct {
	*sql.Tx
	TableName string
}

// Query TxGomaStringTypesDao query
func (g TxGomaStringTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxGomaStringTypesDao exec
func (g TxGomaStringTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ GomaStringTypesDaoQueryer = (*TxGomaStringTypesDao)(nil)

// TxGomaStringTypes is GomaStringTypesDao.
func TxGomaStringTypes(tx *sql.Tx) TxGomaStringTypesDao {
	tblDao := TxGomaStringTypesDao{}
	tblDao.Tx = tx
	tblDao.TableName = "GomaStringTypes"
	return tblDao
}

// SelectAll select goma_string_types table all recode.
func (g GomaStringTypesDao) SelectAll() ([]entity.GomaStringTypesEntity, error) {
	return _GomaStringTypesSelectAll(g)
}

// SelectAll transaction select goma_string_types table all recode.
func (g TxGomaStringTypesDao) SelectAll() ([]entity.GomaStringTypesEntity, error) {
	return _GomaStringTypesSelectAll(g)
}

func _GomaStringTypesSelectAll(g GomaStringTypesDaoQueryer) ([]entity.GomaStringTypesEntity, error) {
	queryString := queryArgs("goma_string_types", "selectAll", nil)

	var es []entity.GomaStringTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.GomaStringTypesEntity
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

// SelectByID select goma_string_types table by primaryKey.
func (g GomaStringTypesDao) SelectByID(id int64) (entity.GomaStringTypesEntity, error) {
	return _GomaStringTypesSelectByID(g, id)
}

// SelectByID transaction select goma_string_types table by primaryKey.
func (g TxGomaStringTypesDao) SelectByID(id int64) (entity.GomaStringTypesEntity, error) {
	return _GomaStringTypesSelectByID(g, id)
}

func _GomaStringTypesSelectByID(g GomaStringTypesDaoQueryer, id int64) (entity.GomaStringTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_string_types", "selectByID", args)

	rows, err := g.Query(queryString)
	if err != nil {
		return entity.GomaStringTypesEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaStringTypesEntity{}, sql.ErrNoRows
	}

	var e entity.GomaStringTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaStringTypesEntity{}, err
	}

	return e, nil
}

// Insert insert goma_string_types table.
func (g GomaStringTypesDao) Insert(e entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesInsert(g, e)
}

// Insert transaction insert goma_string_types table.
func (g TxGomaStringTypesDao) Insert(e entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesInsert(g, e)
}

func _GomaStringTypesInsert(g GomaStringTypesDaoQueryer, e entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":              e.ID,
		"text_columns":    e.TextColumns,
		"char_columns":    e.CharColumns,
		"varchar_columns": e.VarcharColumns,
	}
	queryString := queryArgs("goma_string_types", "insert", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_string_types table.
func (g GomaStringTypesDao) Update(e entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesUpdate(g, e)
}

// Update transaction update goma_string_types table.
func (g TxGomaStringTypesDao) Update(e entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesUpdate(g, e)
}

// Update update goma_string_types table.
func _GomaStringTypesUpdate(g GomaStringTypesDaoQueryer, e entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":              e.ID,
		"text_columns":    e.TextColumns,
		"char_columns":    e.CharColumns,
		"varchar_columns": e.VarcharColumns,
	}
	queryString := queryArgs("goma_string_types", "update", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_string_types table.
func (g GomaStringTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaStringTypesDelete(g, id)
}

// Delete transaction delete goma_string_types table.
func (g TxGomaStringTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaStringTypesDelete(g, id)
}

// Delete delete goma_string_types table by primaryKey.
func _GomaStringTypesDelete(g GomaStringTypesDaoQueryer, id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_string_types", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
