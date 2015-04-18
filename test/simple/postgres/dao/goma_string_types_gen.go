package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/simple/postgres/entity"
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
	queryString := `
select
  id
, text_columns
, char_columns
, varchar_columns
FROM
  goma_string_types`

	var es []entity.GomaStringTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

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
	queryString := `
select
  id
, text_columns
, char_columns
, varchar_columns
FROM
  goma_string_types
WHERE
  id = $1
`
	rows, err := g.Query(queryString,
		id,
	)
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
func (g GomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesInsert(g, entity)
}

// Insert transaction insert goma_string_types table.
func (g TxGomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesInsert(g, entity)
}

func _GomaStringTypesInsert(g GomaStringTypesDaoQueryer, entity entity.GomaStringTypesEntity) (sql.Result, error) {
	queryString := `
insert into goma_string_types(
  id
, text_columns
, char_columns
, varchar_columns
) values(
  $1
, $2
, $3
, $4
);`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.TextColumns,
		entity.CharColumns,
		entity.VarcharColumns,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_string_types table.
func (g GomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesUpdate(g, entity)
}

// Update transaction update goma_string_types table.
func (g TxGomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	return _GomaStringTypesUpdate(g, entity)
}

// Update update goma_string_types table.
func _GomaStringTypesUpdate(g GomaStringTypesDaoQueryer, entity entity.GomaStringTypesEntity) (sql.Result, error) {
	queryString := `
update goma_string_types set
    id = $1
,   text_columns = $2
,   char_columns = $3
,   varchar_columns = $4
 where
    id = $1

`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.TextColumns,
		entity.CharColumns,
		entity.VarcharColumns,

		entity.ID,
	)
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
	queryString := `
delete
from
  goma_string_types
where
  id = $1

`
	result, err := g.Exec(queryString,
		id,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
