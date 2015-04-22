package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/simple/mysql/entity"
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
func (g GomaNumericTypesDao) SelectAll() ([]entity.GomaNumericTypesEntity, error) {
	return _GomaNumericTypesSelectAll(g)
}

// SelectAll transaction select goma_numeric_types table all recode.
func (g TxGomaNumericTypesDao) SelectAll() ([]entity.GomaNumericTypesEntity, error) {
	return _GomaNumericTypesSelectAll(g)
}

func _GomaNumericTypesSelectAll(g GomaNumericTypesDaoQueryer) ([]entity.GomaNumericTypesEntity, error) {
	queryString := `
select
  id
, tinyint_columns
, bool_columns
, smallint_columns
, mediumint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
, double_columns
FROM
  goma_numeric_types`

	var es []entity.GomaNumericTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaNumericTypesEntity
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
func (g GomaNumericTypesDao) SelectByID(id int64) (entity.GomaNumericTypesEntity, error) {
	return _GomaNumericTypesSelectByID(g, id)
}

// SelectByID transaction select goma_numeric_types table by primaryKey.
func (g TxGomaNumericTypesDao) SelectByID(id int64) (entity.GomaNumericTypesEntity, error) {
	return _GomaNumericTypesSelectByID(g, id)
}

func _GomaNumericTypesSelectByID(g GomaNumericTypesDaoQueryer, id int64) (entity.GomaNumericTypesEntity, error) {
	queryString := `
select
  id
, tinyint_columns
, bool_columns
, smallint_columns
, mediumint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
, double_columns
FROM
  goma_numeric_types
WHERE
  id = ?
`
	rows, err := g.Query(queryString,
		id,
	)
	if err != nil {
		return entity.GomaNumericTypesEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaNumericTypesEntity{}, sql.ErrNoRows
	}

	var e entity.GomaNumericTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaNumericTypesEntity{}, err
	}

	return e, nil
}

// Insert insert goma_numeric_types table.
func (g GomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, entity)
}

// Insert transaction insert goma_numeric_types table.
func (g TxGomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, entity)
}

func _GomaNumericTypesInsert(g GomaNumericTypesDaoQueryer, entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	queryString := `
insert into goma_numeric_types(
  id
, tinyint_columns
, bool_columns
, smallint_columns
, mediumint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
, double_columns
) values(
  ?
, ?
, ?
, ?
, ?
, ?
, ?
, ?
, ?
, ?
, ?
, ?
)`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.TinyintColumns,
		entity.BoolColumns,
		entity.SmallintColumns,
		entity.MediumintColumns,
		entity.IntColumns,
		entity.IntegerColumns,
		entity.SerialColumns,
		entity.DecimalColumns,
		entity.NumericColumns,
		entity.FloatColumns,
		entity.DoubleColumns,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (g GomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, entity)
}

// Update transaction update goma_numeric_types table.
func (g TxGomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, entity)
}

// Update update goma_numeric_types table.
func _GomaNumericTypesUpdate(g GomaNumericTypesDaoQueryer, entity entity.GomaNumericTypesEntity) (sql.Result, error) {
	queryString := `
update goma_numeric_types set
    id = ?
,   tinyint_columns = ?
,   bool_columns = ?
,   smallint_columns = ?
,   mediumint_columns = ?
,   int_columns = ?
,   integer_columns = ?
,   serial_columns = ?
,   decimal_columns = ?
,   numeric_columns = ?
,   float_columns = ?
,   double_columns = ?
 where
    id = ?

`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.TinyintColumns,
		entity.BoolColumns,
		entity.SmallintColumns,
		entity.MediumintColumns,
		entity.IntColumns,
		entity.IntegerColumns,
		entity.SerialColumns,
		entity.DecimalColumns,
		entity.NumericColumns,
		entity.FloatColumns,
		entity.DoubleColumns,

		entity.ID,
	)
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
	queryString := `
delete
from
  goma_numeric_types
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
