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

// GomaStringTypesDao is generated goma_string_types table.
type GomaStringTypesDao struct {
	*goma.Goma
	TableName string
}

var sGomaStringTypes *GomaStringTypesDao

// GomaStringTypes is GomaStringTypesDao singleton.
func GomaStringTypes(g *goma.Goma) *GomaStringTypesDao {
	if sGomaStringTypes == nil {
		sGomaStringTypes = &GomaStringTypesDao{
			Goma:      g,
			TableName: "GomaStringTypes",
		}
	}
	return sGomaStringTypes
}

// SelectAll select goma_string_types table all recode.
func (d *GomaStringTypesDao) SelectAll() ([]*entity.GomaStringTypesEntity, error) {

	queryString := d.QueryArgs("goma_string_types", "selectAll", nil)

	var entitys []*entity.GomaStringTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaStringTypesEntity
		err = rows.Scan(&entity.ID, &entity.TextColumns, &entity.TinytextColumns, &entity.MediumtextColumns, &entity.LongtextColumns, &entity.CharColumns, &entity.VarcharColumns)
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

// SelectByID select goma_string_types table by primaryKey.
func (d *GomaStringTypesDao) SelectByID(id int64) (*entity.GomaStringTypesEntity, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_string_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaStringTypesEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.TextColumns, &entity.TinytextColumns, &entity.MediumtextColumns, &entity.LongtextColumns, &entity.CharColumns, &entity.VarcharColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_string_types table.
func (d *GomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                 entity.ID,
		"text_columns":       entity.TextColumns,
		"tinytext_columns":   entity.TinytextColumns,
		"mediumtext_columns": entity.MediumtextColumns,
		"longtext_columns":   entity.LongtextColumns,
		"char_columns":       entity.CharColumns,
		"varchar_columns":    entity.VarcharColumns,
	}
	queryString := d.QueryArgs("goma_string_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_string_types table.
func (d *GomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":                 entity.ID,
		"text_columns":       entity.TextColumns,
		"tinytext_columns":   entity.TinytextColumns,
		"mediumtext_columns": entity.MediumtextColumns,
		"longtext_columns":   entity.LongtextColumns,
		"char_columns":       entity.CharColumns,
		"varchar_columns":    entity.VarcharColumns,
	}
	queryString := d.QueryArgs("goma_string_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_string_types table by primaryKey.
func (d *GomaStringTypesDao) Delete(id int64) (sql.Result, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_string_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
