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

// GomaBinaryTypesDao is generated goma_binary_types table.
type GomaBinaryTypesDao struct {
	*goma.Goma
	tx        *sql.Tx
	TableName string
}

// GomaBinaryTypes is GomaBinaryTypesDao.
func GomaBinaryTypes(g *goma.Goma) GomaBinaryTypesDao {
	tblDao := GomaBinaryTypesDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "GomaBinaryTypes"
	return tblDao
}

// IsTx started transaction?
func (d *GomaBinaryTypesDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *GomaBinaryTypesDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *GomaBinaryTypesDao) ResetTx() {
	d.tx = nil
}

func (d *GomaBinaryTypesDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d *GomaBinaryTypesDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select goma_binary_types table all recode.
func (d *GomaBinaryTypesDao) SelectAll() ([]*entity.GomaBinaryTypesEntity, error) {
	queryString := d.QueryArgs("goma_binary_types", "selectAll", nil)

	var entitys []*entity.GomaBinaryTypesEntity
	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.GomaBinaryTypesEntity
		err = rows.Scan(&entity.ID, &entity.BinaryColumns, &entity.TinyblobColumns, &entity.BlobColumns, &entity.MediumblobColumns, &entity.LongblobColumns, &entity.VarbinaryColumns)
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

// SelectByID select goma_binary_types table by primaryKey.
func (d *GomaBinaryTypesDao) SelectByID(id int64) (*entity.GomaBinaryTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_binary_types", "selectByID", args)

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.GomaBinaryTypesEntity
	if err := rows.Scan(&entity.ID, &entity.BinaryColumns, &entity.TinyblobColumns, &entity.BlobColumns, &entity.MediumblobColumns, &entity.LongblobColumns, &entity.VarbinaryColumns); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert goma_binary_types table.
func (d *GomaBinaryTypesDao) Insert(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"binary_columns":     entity.BinaryColumns,
		"tinyblob_columns":   entity.TinyblobColumns,
		"blob_columns":       entity.BlobColumns,
		"mediumblob_columns": entity.MediumblobColumns,
		"longblob_columns":   entity.LongblobColumns,
		"varbinary_columns":  entity.VarbinaryColumns,
	}
	queryString := d.QueryArgs("goma_binary_types", "insert", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (d *GomaBinaryTypesDao) Update(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"binary_columns":     entity.BinaryColumns,
		"tinyblob_columns":   entity.TinyblobColumns,
		"blob_columns":       entity.BlobColumns,
		"mediumblob_columns": entity.MediumblobColumns,
		"longblob_columns":   entity.LongblobColumns,
		"varbinary_columns":  entity.VarbinaryColumns,
	}
	queryString := d.QueryArgs("goma_binary_types", "update", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_binary_types table by primaryKey.
func (d *GomaBinaryTypesDao) Delete(id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("goma_binary_types", "delete", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
