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

// GomaNumericTypesDao is generated goma_numeric_types table.
type GomaNumericTypesDao struct {
	*goma.Goma
	tx        *sql.Tx
	TableName string
}

// GomaNumericTypes is GomaNumericTypesDao.
func GomaNumericTypes(g *goma.Goma) GomaNumericTypesDao {
	tblDao := GomaNumericTypesDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "GomaNumericTypes"
	return tblDao
}

// IsTx started transaction?
func (d GomaNumericTypesDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *GomaNumericTypesDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *GomaNumericTypesDao) ResetTx() {
	d.tx = nil
}

func (d GomaNumericTypesDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d GomaNumericTypesDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select goma_numeric_types table all recode.
func (d GomaNumericTypesDao) SelectAll() ([]*entity.GomaNumericTypesEntity, error) {
	queryString := d.QueryArgs("goma_numeric_types", "selectAll", nil)

	var entitys []*entity.GomaNumericTypesEntity
	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaNumericTypesEntity
		err = rows.Scan(&entity.ID, &entity.BoolColumns, &entity.SmallintColumns, &entity.IntColumns, &entity.IntegerColumns, &entity.SerialColumns, &entity.DecimalColumns, &entity.NumericColumns, &entity.FloatColumns)
		if err != nil {
			break
		}

		entitys = append(entitys, &entity)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return entitys, nil
}

// SelectByID select goma_numeric_types table by primaryKey.
func (d GomaNumericTypesDao) SelectByID(id int64) (*entity.GomaNumericTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_numeric_types", "selectByID", args)

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaNumericTypesEntity
	if err := rows.Scan(&entity.ID, &entity.BoolColumns, &entity.SmallintColumns, &entity.IntColumns, &entity.IntegerColumns, &entity.SerialColumns, &entity.DecimalColumns, &entity.NumericColumns, &entity.FloatColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_numeric_types table.
func (d GomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
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
	queryString := d.QueryArgs("goma_numeric_types", "insert", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (d GomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {
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
	queryString := d.QueryArgs("goma_numeric_types", "update", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_numeric_types table by primaryKey.
func (d GomaNumericTypesDao) Delete(id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_numeric_types", "delete", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
