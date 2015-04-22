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

// SampleDaoQueryer is interface
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

// TxSampleDao is generated sample table transaction.
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
func (g SampleDao) SelectAll() ([]entity.SampleEntity, error) {
	return _SampleSelectAll(g)
}

// SelectAll transaction select sample table all recode.
func (g TxSampleDao) SelectAll() ([]entity.SampleEntity, error) {
	return _SampleSelectAll(g)
}

func _SampleSelectAll(g SampleDaoQueryer) ([]entity.SampleEntity, error) {
	queryString := queryArgs("sample", "selectAll", nil)

	var es []entity.SampleEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	for rows.Next() {
		var e entity.SampleEntity
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

// SelectByID select sample table by primaryKey.
func (g SampleDao) SelectByID(id int) (entity.SampleEntity, error) {
	return _SampleSelectByID(g, id)
}

// SelectByID transaction select sample table by primaryKey.
func (g TxSampleDao) SelectByID(id int) (entity.SampleEntity, error) {
	return _SampleSelectByID(g, id)
}

func _SampleSelectByID(g SampleDaoQueryer, id int) (entity.SampleEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("sample", "selectByID", args)

	rows, err := g.Query(queryString)
	if err != nil {
		return entity.SampleEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.SampleEntity{}, sql.ErrNoRows
	}

	var e entity.SampleEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.SampleEntity{}, err
	}

	return e, nil
}

// Insert insert sample table.
func (g SampleDao) Insert(entity entity.SampleEntity) (sql.Result, error) {
	return _SampleInsert(g, entity)
}

// Insert transaction insert sample table.
func (g TxSampleDao) Insert(entity entity.SampleEntity) (sql.Result, error) {
	return _SampleInsert(g, entity)
}

func _SampleInsert(g SampleDaoQueryer, entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("sample", "insert", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update sample table.
func (g SampleDao) Update(entity entity.SampleEntity) (sql.Result, error) {
	return _SampleUpdate(g, entity)
}

// Update transaction update sample table.
func (g TxSampleDao) Update(entity entity.SampleEntity) (sql.Result, error) {
	return _SampleUpdate(g, entity)
}

// Update update sample table.
func _SampleUpdate(g SampleDaoQueryer, entity entity.SampleEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("sample", "update", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete sample table.
func (g SampleDao) Delete(id int) (sql.Result, error) {
	return _SampleDelete(g, id)
}

// Delete transaction delete sample table.
func (g TxSampleDao) Delete(id int) (sql.Result, error) {
	return _SampleDelete(g, id)
}

// Delete delete sample table by primaryKey.
func _SampleDelete(g SampleDaoQueryer, id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("sample", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
