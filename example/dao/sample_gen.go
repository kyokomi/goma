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

type SampleDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// SampleDao is generated sample table.
type SampleDao struct {
	*sql.DB
	TableName string
}

// Query SampleDao query
func (g SampleDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec SampleDao exec
func (g SampleDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ SampleDaoQueryer = (*SampleDao)(nil)

// Sample is SampleDao.
func Sample(db *sql.DB) SampleDao {
	tblDao := SampleDao{}
	tblDao.DB = db
	tblDao.TableName = "Sample"
	return tblDao
}

// SampleDaoTx is generated sample table transaction.
type TxSampleDao struct {
	*sql.Tx
	TableName string
}

// Query TxSampleDao query
func (g TxSampleDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxSampleDao exec
func (g TxSampleDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ SampleDaoQueryer = (*TxSampleDao)(nil)

// TxSample is SampleDao.
func TxSample(tx *sql.Tx) TxSampleDao {
	tblDao := TxSampleDao{}
	tblDao.Tx = tx
	tblDao.TableName = "Sample"
	return tblDao
}

// SelectAll select sample table all recode.
func (d SampleDao) SelectAll() ([]*entity.SampleEntity, error) {
	return sampleSelectAll(d)
}

// SelectAll transaction select sample table all recode.
func (d TxSampleDao) SelectAll() ([]*entity.SampleEntity, error) {
	return sampleSelectAll(d)
}

func sampleSelectAll(d SampleDaoQueryer) ([]*entity.SampleEntity, error) {
	queryString := queryArgs("sample", "selectAll", nil)

	var es []*entity.SampleEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.SampleEntity
		if err := e.Scan(rows); err != nil {
			break
		}

		es = append(es, &e)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return es, nil
}

// SelectByID select sample table by primaryKey.
func (d SampleDao) SelectByID(id int) (*entity.SampleEntity, error) {
	return sampleSelectByID(d, id)
}

// SelectByID transaction select sample table by primaryKey.
func (d TxSampleDao) SelectByID(id int) (*entity.SampleEntity, error) {
	return sampleSelectByID(d, id)
}

func sampleSelectByID(d SampleDaoQueryer, id int) (*entity.SampleEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("sample", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.SampleEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert sample table.
func (d SampleDao) Insert(entity entity.SampleEntity) (sql.Result, error) {
	return sampleInsert(d, entity)
}

// Insert transaction insert sample table.
func (d TxSampleDao) Insert(entity entity.SampleEntity) (sql.Result, error) {
	return sampleInsert(d, entity)
}

func sampleInsert(d SampleDaoQueryer, entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("sample", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update sample table.
func (d SampleDao) Update(entity entity.SampleEntity) (sql.Result, error) {
	return sampleUpdate(d, entity)
}

// Update transaction update sample table.
func (d TxSampleDao) Update(entity entity.SampleEntity) (sql.Result, error) {
	return sampleUpdate(d, entity)
}

// Update update sample table.
func sampleUpdate(d SampleDaoQueryer, entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("sample", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete sample table.
func (d SampleDao) Delete(id int) (sql.Result, error) {
	return sampleDelete(d, id)
}

// Delete transaction delete sample table.
func (d TxSampleDao) Delete(id int) (sql.Result, error) {
	return sampleDelete(d, id)
}

// Delete delete sample table by primaryKey.
func sampleDelete(d SampleDaoQueryer, id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("sample", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
