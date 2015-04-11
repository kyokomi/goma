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

// GomaNumericTypesDaoTx is generated goma_numeric_types table transaction.
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
func (d GomaNumericTypesDao) SelectAll() ([]*entity.GomaNumericTypesEntity, error) {
	return goma_numeric_typesSelectAll(d)
}

// SelectAll transaction select goma_numeric_types table all recode.
func (d TxGomaNumericTypesDao) SelectAll() ([]*entity.GomaNumericTypesEntity, error) {
	return goma_numeric_typesSelectAll(d)
}

func goma_numeric_typesSelectAll(d GomaNumericTypesDaoQueryer) ([]*entity.GomaNumericTypesEntity, error) {
	queryString := queryArgs("goma_numeric_types", "selectAll", nil)

	var es []*entity.GomaNumericTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaNumericTypesEntity
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

// SelectByID select goma_numeric_types table by primaryKey.
func (d GomaNumericTypesDao) SelectByID(id int64) (*entity.GomaNumericTypesEntity, error) {
	return goma_numeric_typesSelectByID(d, id)
}

// SelectByID transaction select goma_numeric_types table by primaryKey.
func (d TxGomaNumericTypesDao) SelectByID(id int64) (*entity.GomaNumericTypesEntity, error) {
	return goma_numeric_typesSelectByID(d, id)
}

func goma_numeric_typesSelectByID(d GomaNumericTypesDaoQueryer, id int64) (*entity.GomaNumericTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_numeric_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.GomaNumericTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert goma_numeric_types table.
func (d GomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return goma_numeric_typesInsert(d, entity)
}

// Insert transaction insert goma_numeric_types table.
func (d TxGomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return goma_numeric_typesInsert(d, entity)
}

func goma_numeric_typesInsert(d GomaNumericTypesDaoQueryer, entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":               entity.ID,
		"bool_columns":     entity.BoolColumns,
		"smallint_columns": entity.SmallintColumns,
		"int_columns":      entity.IntColumns,
		"integer_columns":  entity.IntegerColumns,
		"serial_columns":   entity.SerialColumns,
		"decimal_columns":  entity.DecimalColumns,
		"numeric_columns":  entity.NumericColumns,
		"float_columns":    entity.FloatColumns,
	}
	queryString := queryArgs("goma_numeric_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (d GomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return goma_numeric_typesUpdate(d, entity)
}

// Update transaction update goma_numeric_types table.
func (d TxGomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return goma_numeric_typesUpdate(d, entity)
}

// Update update goma_numeric_types table.
func goma_numeric_typesUpdate(d GomaNumericTypesDaoQueryer, entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":               entity.ID,
		"bool_columns":     entity.BoolColumns,
		"smallint_columns": entity.SmallintColumns,
		"int_columns":      entity.IntColumns,
		"integer_columns":  entity.IntegerColumns,
		"serial_columns":   entity.SerialColumns,
		"decimal_columns":  entity.DecimalColumns,
		"numeric_columns":  entity.NumericColumns,
		"float_columns":    entity.FloatColumns,
	}
	queryString := queryArgs("goma_numeric_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_numeric_types table.
func (d GomaNumericTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_numeric_typesDelete(d, id)
}

// Delete transaction delete goma_numeric_types table.
func (d TxGomaNumericTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_numeric_typesDelete(d, id)
}

// Delete delete goma_numeric_types table by primaryKey.
func goma_numeric_typesDelete(d GomaNumericTypesDaoQueryer, id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_numeric_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
