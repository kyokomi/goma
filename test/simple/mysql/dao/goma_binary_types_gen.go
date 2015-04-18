package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/simple/mysql/entity"
)

// GomaBinaryTypesDaoQueryer is interface
type GomaBinaryTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaBinaryTypesDao is generated goma_binary_types table.
type GomaBinaryTypesDao struct {
	*sql.DB
	TableName string
}

// Query GomaBinaryTypesDao query
func (g GomaBinaryTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec GomaBinaryTypesDao exec
func (g GomaBinaryTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ GomaBinaryTypesDaoQueryer = (*GomaBinaryTypesDao)(nil)

// GomaBinaryTypes is GomaBinaryTypesDao.
func GomaBinaryTypes(db *sql.DB) GomaBinaryTypesDao {
	tblDao := GomaBinaryTypesDao{}
	tblDao.DB = db
	tblDao.TableName = "GomaBinaryTypes"
	return tblDao
}

// TxGomaBinaryTypesDao is generated goma_binary_types table transaction.
type TxGomaBinaryTypesDao struct {
	*sql.Tx
	TableName string
}

// Query TxGomaBinaryTypesDao query
func (g TxGomaBinaryTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxGomaBinaryTypesDao exec
func (g TxGomaBinaryTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ GomaBinaryTypesDaoQueryer = (*TxGomaBinaryTypesDao)(nil)

// TxGomaBinaryTypes is GomaBinaryTypesDao.
func TxGomaBinaryTypes(tx *sql.Tx) TxGomaBinaryTypesDao {
	tblDao := TxGomaBinaryTypesDao{}
	tblDao.Tx = tx
	tblDao.TableName = "GomaBinaryTypes"
	return tblDao
}

// SelectAll select goma_binary_types table all recode.
func (g GomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectAll(g)
}

// SelectAll transaction select goma_binary_types table all recode.
func (g TxGomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectAll(g)
}

func _GomaBinaryTypesSelectAll(g GomaBinaryTypesDaoQueryer) ([]entity.GomaBinaryTypesEntity, error) {
	queryString := `
select
  id
, binary_columns
, tinyblob_columns
, blob_columns
, mediumblob_columns
, longblob_columns
, varbinary_columns
FROM
  goma_binary_types`

	var es []entity.GomaBinaryTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	for rows.Next() {
		var e entity.GomaBinaryTypesEntity
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

// SelectByID select goma_binary_types table by primaryKey.
func (g GomaBinaryTypesDao) SelectByID(id int64) (entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectByID(g, id)
}

// SelectByID transaction select goma_binary_types table by primaryKey.
func (g TxGomaBinaryTypesDao) SelectByID(id int64) (entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectByID(g, id)
}

func _GomaBinaryTypesSelectByID(g GomaBinaryTypesDaoQueryer, id int64) (entity.GomaBinaryTypesEntity, error) {
	queryString := `
select
  id
, binary_columns
, tinyblob_columns
, blob_columns
, mediumblob_columns
, longblob_columns
, varbinary_columns
FROM
  goma_binary_types
WHERE
  id = ?
`
	rows, err := g.Query(queryString,
		id,
	)
	if err != nil {
		return entity.GomaBinaryTypesEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaBinaryTypesEntity{}, sql.ErrNoRows
	}

	var e entity.GomaBinaryTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaBinaryTypesEntity{}, err
	}

	return e, nil
}

// Insert insert goma_binary_types table.
func (g GomaBinaryTypesDao) Insert(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, entity)
}

// Insert transaction insert goma_binary_types table.
func (g TxGomaBinaryTypesDao) Insert(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, entity)
}

func _GomaBinaryTypesInsert(g GomaBinaryTypesDaoQueryer, entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	queryString := `
insert into goma_binary_types(
  id
, binary_columns
, tinyblob_columns
, blob_columns
, mediumblob_columns
, longblob_columns
, varbinary_columns
) values(
  ?
, ?
, ?
, ?
, ?
, ?
, ?
);`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.BinaryColumns,
		entity.TinyblobColumns,
		entity.BlobColumns,
		entity.MediumblobColumns,
		entity.LongblobColumns,
		entity.VarbinaryColumns,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (g GomaBinaryTypesDao) Update(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, entity)
}

// Update transaction update goma_binary_types table.
func (g TxGomaBinaryTypesDao) Update(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, entity)
}

// Update update goma_binary_types table.
func _GomaBinaryTypesUpdate(g GomaBinaryTypesDaoQueryer, entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	queryString := `
update goma_binary_types set
    id = ?
,   binary_columns = ?
,   tinyblob_columns = ?
,   blob_columns = ?
,   mediumblob_columns = ?
,   longblob_columns = ?
,   varbinary_columns = ?
 where
    id = ?

`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.BinaryColumns,
		entity.TinyblobColumns,
		entity.BlobColumns,
		entity.MediumblobColumns,
		entity.LongblobColumns,
		entity.VarbinaryColumns,

		entity.ID,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_binary_types table.
func (g GomaBinaryTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, id)
}

// Delete transaction delete goma_binary_types table.
func (g TxGomaBinaryTypesDao) Delete(id int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, id)
}

// Delete delete goma_binary_types table by primaryKey.
func _GomaBinaryTypesDelete(g GomaBinaryTypesDaoQueryer, id int64) (sql.Result, error) {
	queryString := `
delete
from
  goma_binary_types
where
  id = ?

`
	result, err := g.Exec(queryString,
		id,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
