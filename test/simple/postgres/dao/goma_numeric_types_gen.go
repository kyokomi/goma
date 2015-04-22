package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/simple/postgres/entity"
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
, bool_columns
, smallint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
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
, bool_columns
, smallint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
FROM
  goma_numeric_types
WHERE
  id = $1
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
func (g GomaNumericTypesDao) Insert(e entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, e)
}

// Insert transaction insert goma_numeric_types table.
func (g TxGomaNumericTypesDao) Insert(e entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesInsert(g, e)
}

func _GomaNumericTypesInsert(g GomaNumericTypesDaoQueryer, e entity.GomaNumericTypesEntity) (sql.Result, error) {
	queryString := `
insert into goma_numeric_types(
  id
, bool_columns
, smallint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
) values(
  $1
, $2
, $3
, $4
, $5
, $6
, $7
, $8
, $9
)`
	result, err := g.Exec(queryString,
		e.ID,
		e.BoolColumns,
		e.SmallintColumns,
		e.IntColumns,
		e.IntegerColumns,
		e.SerialColumns,
		e.DecimalColumns,
		e.NumericColumns,
		e.FloatColumns,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (g GomaNumericTypesDao) Update(e entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, e)
}

// Update transaction update goma_numeric_types table.
func (g TxGomaNumericTypesDao) Update(e entity.GomaNumericTypesEntity) (sql.Result, error) {
	return _GomaNumericTypesUpdate(g, e)
}

// Update update goma_numeric_types table.
func _GomaNumericTypesUpdate(g GomaNumericTypesDaoQueryer, e entity.GomaNumericTypesEntity) (sql.Result, error) {
	queryString := `
update goma_numeric_types set
    id = $1
,   bool_columns = $2
,   smallint_columns = $3
,   int_columns = $4
,   integer_columns = $5
,   serial_columns = $6
,   decimal_columns = $7
,   numeric_columns = $8
,   float_columns = $9
 where
    id = $1

`
	result, err := g.Exec(queryString,
		e.ID,
		e.BoolColumns,
		e.SmallintColumns,
		e.IntColumns,
		e.IntegerColumns,
		e.SerialColumns,
		e.DecimalColumns,
		e.NumericColumns,
		e.FloatColumns,

		e.ID,
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
