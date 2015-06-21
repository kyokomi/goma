package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/example/entity"
)

var tableSample = "sample"
var columnsSample = []string{
	"id",
	"name",
	"create_at",
}

// SampleTableName sample table name
func SampleTableName() string {
	return tableSample
}

// SampleTableColumns sample table columns
func SampleTableColumns() []string {
	return columnsSample
}

// SampleDaoQueryer is interface
type SampleDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// SampleDao is generated sample table.
type SampleDao struct {
	*sql.DB
	TableName string
	Columns   []string
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
	tblDao.TableName = tableSample
	tblDao.Columns = columnsSample
	return tblDao
}

// TxSampleDao is generated sample table transaction.
type TxSampleDao struct {
	*sql.Tx
	TableName string
	Columns   []string
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
	tblDao.TableName = tableSample
	tblDao.Columns = columnsSample
	return tblDao
}

// SelectAll select sample table all recode.
func (g SampleDao) SelectAll() ([]entity.Sample, error) {
	return _SampleSelectAll(g)
}

// SelectAll transaction select sample table all recode.
func (g TxSampleDao) SelectAll() ([]entity.Sample, error) {
	return _SampleSelectAll(g)
}

func _SampleSelectAll(g SampleDaoQueryer) ([]entity.Sample, error) {
	queryString, args, err := queryArgs("sample", "selectAll", nil)
	if err != nil {
		return nil, err
	}

	var es []entity.Sample
	rows, err := g.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.Sample
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
func (g SampleDao) SelectByID(id int) (entity.Sample, error) {
	return _SampleSelectByID(g, id)
}

// SelectByID transaction select sample table by primaryKey.
func (g TxSampleDao) SelectByID(id int) (entity.Sample, error) {
	return _SampleSelectByID(g, id)
}

func _SampleSelectByID(g SampleDaoQueryer, id int) (entity.Sample, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("sample", "selectByID", argsMap)
	if err != nil {
		return entity.Sample{}, err
	}

	rows, err := g.Query(queryString, args...)
	if err != nil {
		return entity.Sample{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Sample{}, sql.ErrNoRows
	}

	var e entity.Sample
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.Sample{}, err
	}

	return e, nil
}

// Insert insert sample table.
func (g SampleDao) Insert(e entity.Sample) (sql.Result, error) {
	return _SampleInsert(g, e)
}

// Insert transaction insert sample table.
func (g TxSampleDao) Insert(e entity.Sample) (sql.Result, error) {
	return _SampleInsert(g, e)
}

func _SampleInsert(g SampleDaoQueryer, e entity.Sample) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":        e.ID,
		"name":      e.Name,
		"create_at": e.CreateAt,
	}
	queryString, args, err := queryArgs("sample", "insert", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update sample table.
func (g SampleDao) Update(e entity.Sample) (sql.Result, error) {
	return _SampleUpdate(g, e)
}

// Update transaction update sample table.
func (g TxSampleDao) Update(e entity.Sample) (sql.Result, error) {
	return _SampleUpdate(g, e)
}

// Update update sample table.
func _SampleUpdate(g SampleDaoQueryer, e entity.Sample) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":        e.ID,
		"name":      e.Name,
		"create_at": e.CreateAt,
	}
	queryString, args, err := queryArgs("sample", "update", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
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
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("sample", "delete", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
