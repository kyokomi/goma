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

// GomaDateTypesDao is generated goma_date_types table.
type GomaDateTypesDao struct {
	*goma.Goma
	TableName string
}

var sGomaDateTypes *GomaDateTypesDao

// GomaDateTypes is GomaDateTypesDao singleton.
func GomaDateTypes(g *goma.Goma) *GomaDateTypesDao {
	if sGomaDateTypes == nil {
		sGomaDateTypes = &GomaDateTypesDao{
			Goma:      g,
			TableName: "GomaDateTypes",
		}
	}
	return sGomaDateTypes
}

// SelectAll select goma_date_types table all recode.
func (d *GomaDateTypesDao) SelectAll() ([]*entity.GomaDateTypesEntity, error) {

	queryString := d.QueryArgs("goma_date_types", "selectAll", nil)

	var entitys []*entity.GomaDateTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaDateTypesEntity
		err = rows.Scan(&entity.ID, &entity.DateColumns, &entity.DatetimeColumns, &entity.TimestampColumns)
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

// SelectByID select goma_date_types table by primaryKey.
func (d *GomaDateTypesDao) SelectByID(id int64) (*entity.GomaDateTypesEntity, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_date_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaDateTypesEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.DateColumns, &entity.DatetimeColumns, &entity.TimestampColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_date_types table.
func (d *GomaDateTypesDao) Insert(entity entity.GomaDateTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := d.QueryArgs("goma_date_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_date_types table.
func (d *GomaDateTypesDao) Update(entity entity.GomaDateTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                entity.ID,
		"date_columns":      entity.DateColumns,
		"datetime_columns":  entity.DatetimeColumns,
		"timestamp_columns": entity.TimestampColumns,
	}
	queryString := d.QueryArgs("goma_date_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_date_types table by primaryKey.
func (d *GomaDateTypesDao) Delete(id int64) (sql.Result, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_date_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
