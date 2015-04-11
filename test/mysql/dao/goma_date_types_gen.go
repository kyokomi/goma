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

type GomaDateTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaDateTypesDao is generated goma_date_types table.
type GomaDateTypesDao struct {
	*sql.DB
	TableName string
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
	tblDao.TableName = "GomaDateTypes"
	return tblDao
}

// GomaDateTypesDaoTx is generated goma_date_types table transaction.
type TxGomaDateTypesDao struct {
	*sql.Tx
	TableName string
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
	tblDao.TableName = "GomaDateTypes"
	return tblDao
}

// SelectAll select goma_date_types table all recode.
func (d GomaDateTypesDao) SelectAll() ([]*entity.GomaDateTypesEntity, error) {
	return goma_date_typesSelectAll(d)
}

// SelectAll transaction select goma_date_types table all recode.
func (d TxGomaDateTypesDao) SelectAll() ([]*entity.GomaDateTypesEntity, error) {
	return goma_date_typesSelectAll(d)
}

func goma_date_typesSelectAll(d GomaDateTypesDaoQueryer) ([]*entity.GomaDateTypesEntity, error) {
	queryString := queryArgs("goma_date_types", "selectAll", nil)

	var es []*entity.GomaDateTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaDateTypesEntity
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

// SelectByID select goma_date_types table by primaryKey.
func (d GomaDateTypesDao) SelectByID(id int64) (*entity.GomaDateTypesEntity, error) {
	return goma_date_typesSelectByID(d, id)
}

// SelectByID transaction select goma_date_types table by primaryKey.
func (d TxGomaDateTypesDao) SelectByID(id int64) (*entity.GomaDateTypesEntity, error) {
	return goma_date_typesSelectByID(d, id)
}

func goma_date_typesSelectByID(d GomaDateTypesDaoQueryer, id int64) (*entity.GomaDateTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_date_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.GomaDateTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert goma_date_types table.
func (d GomaDateTypesDao) Insert(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return goma_date_typesInsert(d, entity)
}

// Insert transaction insert goma_date_types table.
func (d TxGomaDateTypesDao) Insert(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return goma_date_typesInsert(d, entity)
}

func goma_date_typesInsert(d GomaDateTypesDaoQueryer, entity entity.GomaDateTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := queryArgs("goma_date_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_date_types table.
func (d GomaDateTypesDao) Update(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return goma_date_typesUpdate(d, entity)
}

// Update transaction update goma_date_types table.
func (d TxGomaDateTypesDao) Update(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return goma_date_typesUpdate(d, entity)
}

// Update update goma_date_types table.
func goma_date_typesUpdate(d GomaDateTypesDaoQueryer, entity entity.GomaDateTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := queryArgs("goma_date_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_date_types table.
func (d GomaDateTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_date_typesDelete(d, id)
}

// Delete transaction delete goma_date_types table.
func (d TxGomaDateTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_date_typesDelete(d, id)
}

// Delete delete goma_date_types table by primaryKey.
func goma_date_typesDelete(d GomaDateTypesDaoQueryer, id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_date_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
