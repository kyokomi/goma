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
func (g GomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectAll(g)
}

// SelectAll transaction select goma_binary_types table all recode.
func (g TxGomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectAll(g)
}

func _GomaBinaryTypesSelectAll(g GomaBinaryTypesDaoQueryer) ([]entity.GomaBinaryTypes, error) {
	queryString := `
select
  binary_id
, binary_columns
, tinyblob_columns
, blob_columns
, mediumblob_columns
, longblob_columns
, varbinary_columns
FROM
  goma_binary_types`

	var es []entity.GomaBinaryTypes
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaBinaryTypes
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
func (g GomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

// SelectByID transaction select goma_binary_types table by primaryKey.
func (g TxGomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

func _GomaBinaryTypesSelectByID(g GomaBinaryTypesDaoQueryer, binaryID int64) (entity.GomaBinaryTypes, error) {
	queryString := `
select
  binary_id
, binary_columns
, tinyblob_columns
, blob_columns
, mediumblob_columns
, longblob_columns
, varbinary_columns
FROM
  goma_binary_types
WHERE
  binary_id = ?
`
	rows, err := g.Query(queryString,
		binaryID,
	)
	if err != nil {
		return entity.GomaBinaryTypes{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaBinaryTypes{}, sql.ErrNoRows
	}

	var e entity.GomaBinaryTypes
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaBinaryTypes{}, err
	}

	return e, nil
}

// Insert insert goma_binary_types table.
func (g GomaBinaryTypesDao) Insert(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

// Insert transaction insert goma_binary_types table.
func (g TxGomaBinaryTypesDao) Insert(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

func _GomaBinaryTypesInsert(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypes) (sql.Result, error) {
	queryString := `
insert into goma_binary_types(
  binary_id
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
)`
	result, err := g.Exec(queryString,
		e.BinaryID,
		e.BinaryColumns,
		e.TinyblobColumns,
		e.BlobColumns,
		e.MediumblobColumns,
		e.LongblobColumns,
		e.VarbinaryColumns,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (g GomaBinaryTypesDao) Update(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update transaction update goma_binary_types table.
func (g TxGomaBinaryTypesDao) Update(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update update goma_binary_types table.
func _GomaBinaryTypesUpdate(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypes) (sql.Result, error) {
	queryString := `
update goma_binary_types set
    binary_id = ?
,   binary_columns = ?
,   tinyblob_columns = ?
,   blob_columns = ?
,   mediumblob_columns = ?
,   longblob_columns = ?
,   varbinary_columns = ?
 where
    binary_id = ?

`
	result, err := g.Exec(queryString,
		e.BinaryID,
		e.BinaryColumns,
		e.TinyblobColumns,
		e.BlobColumns,
		e.MediumblobColumns,
		e.LongblobColumns,
		e.VarbinaryColumns,

		e.BinaryID,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_binary_types table.
func (g GomaBinaryTypesDao) Delete(binaryID int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, binaryID)
}

// Delete transaction delete goma_binary_types table.
func (g TxGomaBinaryTypesDao) Delete(binaryID int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, binaryID)
}

// Delete delete goma_binary_types table by primaryKey.
func _GomaBinaryTypesDelete(g GomaBinaryTypesDaoQueryer, binaryID int64) (sql.Result, error) {
	queryString := `
delete
from
  goma_binary_types
where
  binary_id = ?

`
	result, err := g.Exec(queryString,
		binaryID,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
