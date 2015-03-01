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

// GomaStringTypesDao is generated goma_string_types table.
type GomaStringTypesDao struct {
	*goma.Goma
	tx        *sql.Tx
	TableName string
}

// GomaStringTypes is GomaStringTypesDao.
func GomaStringTypes(g *goma.Goma) GomaStringTypesDao {
	tblDao := GomaStringTypesDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "GomaStringTypes"
	return tblDao
}

// IsTx started transaction?
func (d *GomaStringTypesDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *GomaStringTypesDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *GomaStringTypesDao) ResetTx() {
	d.tx = nil
}

func (d *GomaStringTypesDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d *GomaStringTypesDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select goma_string_types table all recode.
func (d *GomaStringTypesDao) SelectAll() ([]*entity.GomaStringTypesEntity, error) {
	queryString := d.QueryArgs("goma_string_types", "selectAll", nil)

	var entitys []*entity.GomaStringTypesEntity
	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaStringTypesEntity
		err = rows.Scan(&entity.ID, &entity.TextColumns, &entity.CharColumns, &entity.VarcharColumns)
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

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaStringTypesEntity
	if err := rows.Scan(&entity.ID, &entity.TextColumns, &entity.CharColumns, &entity.VarcharColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_string_types table.
func (d *GomaStringTypesDao) Insert(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":              entity.ID,
		"text_columns":    entity.TextColumns,
		"char_columns":    entity.CharColumns,
		"varchar_columns": entity.VarcharColumns,
	}
	queryString := d.QueryArgs("goma_string_types", "insert", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_string_types table.
func (d *GomaStringTypesDao) Update(entity entity.GomaStringTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":              entity.ID,
		"text_columns":    entity.TextColumns,
		"char_columns":    entity.CharColumns,
		"varchar_columns": entity.VarcharColumns,
	}
	queryString := d.QueryArgs("goma_string_types", "update", args)

	result, err := d.daoExec(queryString)
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

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
