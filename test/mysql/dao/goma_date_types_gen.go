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

// GomaDateTypesDaoQueryer is interface
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

// TxGomaDateTypesDao is generated goma_date_types table transaction.
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
func (g GomaDateTypesDao) SelectAll() ([]entity.GomaDateTypesEntity, error) {
	return _GomaDateTypesSelectAll(g)
}

// SelectAll transaction select goma_date_types table all recode.
func (g TxGomaDateTypesDao) SelectAll() ([]entity.GomaDateTypesEntity, error) {
	return _GomaDateTypesSelectAll(g)
}

func _GomaDateTypesSelectAll(g GomaDateTypesDaoQueryer) ([]entity.GomaDateTypesEntity, error) {
	queryString := queryArgs("goma_date_types", "selectAll", nil)

	var es []entity.GomaDateTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	for rows.Next() {
		var e entity.GomaDateTypesEntity
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
func (g GomaDateTypesDao) SelectByID(id int64) (entity.GomaDateTypesEntity, error) {
	return _GomaDateTypesSelectByID(g, id)
}

// SelectByID transaction select goma_date_types table by primaryKey.
func (g TxGomaDateTypesDao) SelectByID(id int64) (entity.GomaDateTypesEntity, error) {
	return _GomaDateTypesSelectByID(g, id)
}

func _GomaDateTypesSelectByID(g GomaDateTypesDaoQueryer, id int64) (entity.GomaDateTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_date_types", "selectByID", args)

	rows, err := g.Query(queryString)
	if err != nil {
		return entity.GomaDateTypesEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaDateTypesEntity{}, sql.ErrNoRows
	}

	var e entity.GomaDateTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaDateTypesEntity{}, err
	}

	return e, nil
}

// Insert insert goma_date_types table.
func (g GomaDateTypesDao) Insert(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return _GomaDateTypesInsert(g, entity)
}

// Insert transaction insert goma_date_types table.
func (g TxGomaDateTypesDao) Insert(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return _GomaDateTypesInsert(g, entity)
}

func _GomaDateTypesInsert(g GomaDateTypesDaoQueryer, entity entity.GomaDateTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := queryArgs("goma_date_types", "insert", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_date_types table.
func (g GomaDateTypesDao) Update(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return _GomaDateTypesUpdate(g, entity)
}

// Update transaction update goma_date_types table.
func (g TxGomaDateTypesDao) Update(entity entity.GomaDateTypesEntity) (sql.Result, error) {
	return _GomaDateTypesUpdate(g, entity)
}

// Update update goma_date_types table.
func _GomaDateTypesUpdate(g GomaDateTypesDaoQueryer, entity entity.GomaDateTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := queryArgs("goma_date_types", "update", args)

	result, err := g.Exec(queryString)
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
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_date_types", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
