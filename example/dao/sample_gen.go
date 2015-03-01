package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/example/entity"

	"github.com/kyokomi/goma"
)

// SampleDao is generated sample table.
type SampleDao struct {
	*goma.Goma
	tx        *sql.Tx
	TableName string
}

// Sample is SampleDao.
func Sample(g *goma.Goma) SampleDao {
	tblDao := SampleDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "Sample"
	return tblDao
}

// IsTx started transaction?
func (d SampleDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *SampleDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *SampleDao) ResetTx() {
	d.tx = nil
}

func (d SampleDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d SampleDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select sample table all recode.
func (d SampleDao) SelectAll() ([]*entity.SampleEntity, error) {
	queryString := d.QueryArgs("sample", "selectAll", nil)

	var entitys []*entity.SampleEntity
	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.SampleEntity
		err = rows.Scan(&entity.ID, &entity.Name, &entity.CreateAt)
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

// SelectByID select sample table by primaryKey.
func (d SampleDao) SelectByID(id int) (*entity.SampleEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("sample", "selectByID", args)

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.SampleEntity
	if err := rows.Scan(&entity.ID, &entity.Name, &entity.CreateAt); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert sample table.
func (d SampleDao) Insert(entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("sample", "insert", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update sample table.
func (d SampleDao) Update(entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("sample", "update", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete sample table by primaryKey.
func (d SampleDao) Delete(id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("sample", "delete", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
