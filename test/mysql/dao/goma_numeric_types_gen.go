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

// GomaNumericTypesDao is generated goma_numeric_types table.
type GomaNumericTypesDao struct {
	*goma.Goma
	TableName string
}

var sGomaNumericTypes *GomaNumericTypesDao

// GomaNumericTypes is GomaNumericTypesDao singleton.
func GomaNumericTypes(g *goma.Goma) *GomaNumericTypesDao {
	if sGomaNumericTypes == nil {
		sGomaNumericTypes = &GomaNumericTypesDao{
			Goma:      g,
			TableName: "GomaNumericTypes",
		}
	}
	return sGomaNumericTypes
}

// SelectAll select goma_numeric_types table all recode.
func (d *GomaNumericTypesDao) SelectAll() ([]*entity.GomaNumericTypesEntity, error) {

	queryString := d.QueryArgs("goma_numeric_types", "selectAll", nil)

	var entitys []*entity.GomaNumericTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaNumericTypesEntity
		err = rows.Scan(&entity.ID, &entity.TinyintColumns, &entity.BoolColumns, &entity.SmallintColumns, &entity.MediumintColumns, &entity.IntColumns, &entity.IntegerColumns, &entity.SerialColumns, &entity.DecimalColumns, &entity.NumericColumns, &entity.FloatColumns, &entity.DoubleColumns)
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
func (d *GomaNumericTypesDao) SelectByID(id int64) (*entity.GomaNumericTypesEntity, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_numeric_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaNumericTypesEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.TinyintColumns, &entity.BoolColumns, &entity.SmallintColumns, &entity.MediumintColumns, &entity.IntColumns, &entity.IntegerColumns, &entity.SerialColumns, &entity.DecimalColumns, &entity.NumericColumns, &entity.FloatColumns, &entity.DoubleColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_numeric_types table.
func (d *GomaNumericTypesDao) Insert(entity entity.GomaNumericTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                entity.ID,
		"tinyint_columns":   entity.TinyintColumns,
		"bool_columns":      entity.BoolColumns,
		"smallint_columns":  entity.SmallintColumns,
		"mediumint_columns": entity.MediumintColumns,
		"int_columns":       entity.IntColumns,
		"integer_columns":   entity.IntegerColumns,
		"serial_columns":    entity.SerialColumns,
		"decimal_columns":   entity.DecimalColumns,
		"numeric_columns":   entity.NumericColumns,
		"float_columns":     entity.FloatColumns,
		"double_columns":    entity.DoubleColumns,
	}
	queryString := d.QueryArgs("goma_numeric_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_numeric_types table.
func (d *GomaNumericTypesDao) Update(entity entity.GomaNumericTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                entity.ID,
		"tinyint_columns":   entity.TinyintColumns,
		"bool_columns":      entity.BoolColumns,
		"smallint_columns":  entity.SmallintColumns,
		"mediumint_columns": entity.MediumintColumns,
		"int_columns":       entity.IntColumns,
		"integer_columns":   entity.IntegerColumns,
		"serial_columns":    entity.SerialColumns,
		"decimal_columns":   entity.DecimalColumns,
		"numeric_columns":   entity.NumericColumns,
		"float_columns":     entity.FloatColumns,
		"double_columns":    entity.DoubleColumns,
	}
	queryString := d.QueryArgs("goma_numeric_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_numeric_types table by primaryKey.
func (d *GomaNumericTypesDao) Delete(id int64) (sql.Result, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_numeric_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
