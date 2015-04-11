package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/mysql/entity"

	"github.com/kyokomi/goma"
)

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

// GomaStringTypesDaoTx is generated goma_string_types table transaction.
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
func (d GomaStringTypesDao) SelectAll() ([]*entity.GomaStringTypesEntity, error) {
	return goma_string_typesSelectAll(d)
}

// SelectAll transaction select goma_string_types table all recode.
func (d TxGomaStringTypesDao) SelectAll() ([]*entity.GomaStringTypesEntity, error) {
	return goma_string_typesSelectAll(d)
}

func goma_string_typesSelectAll(d GomaStringTypesDaoQueryer) ([]*entity.GomaStringTypesEntity, error) {
	queryString := queryArgs("goma_string_types", "selectAll", nil)

	var es []*entity.GomaStringTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaStringTypesEntity
		if err := e.Scan(rows); err != nil {
			break
		}

		es = append(es, &e)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return es, nil
}

// SelectByID select goma_string_types table by primaryKey.
func (d GomaStringTypesDao) SelectByID(id int64) (*entity.GomaStringTypesEntity, error) {
	return goma_string_typesSelectByID(d, id)
}

// SelectByID transaction select goma_string_types table by primaryKey.
func (d TxGomaStringTypesDao) SelectByID(id int64) (*entity.GomaStringTypesEntity, error) {
	return goma_string_typesSelectByID(d, id)
}

func goma_string_typesSelectByID(d GomaStringTypesDaoQueryer, id int64) (*entity.GomaStringTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_string_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.GomaStringTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert goma_string_types table.
func (d GomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return goma_string_typesInsert(d, entity)
}

// Insert transaction insert goma_string_types table.
func (d TxGomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return goma_string_typesInsert(d, entity)
}

func goma_string_typesInsert(d GomaStringTypesDaoQueryer, entity entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"text_columns":       entity.TextColumns,
		"tinytext_columns":   entity.TinytextColumns,
		"mediumtext_columns": entity.MediumtextColumns,
		"longtext_columns":   entity.LongtextColumns,
		"char_columns":       entity.CharColumns,
		"varchar_columns":    entity.VarcharColumns,
	}
	queryString := queryArgs("goma_string_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_string_types table.
func (d GomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return goma_string_typesUpdate(d, entity)
}

// Update transaction update goma_string_types table.
func (d TxGomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return goma_string_typesUpdate(d, entity)
}

// Update update goma_string_types table.
func goma_string_typesUpdate(d GomaStringTypesDaoQueryer, entity entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"text_columns":       entity.TextColumns,
		"tinytext_columns":   entity.TinytextColumns,
		"mediumtext_columns": entity.MediumtextColumns,
		"longtext_columns":   entity.LongtextColumns,
		"char_columns":       entity.CharColumns,
		"varchar_columns":    entity.VarcharColumns,
	}
	queryString := queryArgs("goma_string_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_string_types table.
func (d GomaStringTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_string_typesDelete(d, id)
}

// Delete transaction delete goma_string_types table.
func (d TxGomaStringTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_string_typesDelete(d, id)
}

// Delete delete goma_string_types table by primaryKey.
func goma_string_typesDelete(d GomaStringTypesDaoQueryer, id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_string_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
