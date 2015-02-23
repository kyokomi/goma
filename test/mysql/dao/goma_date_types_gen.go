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
	tx        *sql.Tx
	TableName string
}

// GomaDateTypes is GomaDateTypesDao.
func GomaDateTypes(g *goma.Goma) *GomaDateTypesDao {
	tblDao := &GomaDateTypesDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "GomaDateTypes"
	return tblDao
}

// IsTx started transaction?
func (d *GomaDateTypesDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *GomaDateTypesDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *GomaDateTypesDao) ResetTx() {
	d.tx = nil
}

func (d *GomaDateTypesDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d *GomaDateTypesDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select goma_date_types table all recode.
func (d *GomaDateTypesDao) SelectAll() ([]*entity.GomaDateTypesEntity, error) {

	queryString := d.QueryArgs("goma_date_types", "selectAll", nil)

	var entitys []*entity.GomaDateTypesEntity
	rows, err := d.daoQuery(queryString)
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

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaDateTypesEntity
	if err := rows.Scan(&entity.ID, &entity.DateColumns, &entity.DatetimeColumns, &entity.TimestampColumns); err != nil {
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

	result, err := d.daoExec(queryString)
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

	result, err := d.daoExec(queryString)
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

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
