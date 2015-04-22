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

// GomaBinaryTypesDaoQueryer is interface
type GomaBinaryTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaBinaryTypesDao is generated goma_binary_types table.
type GomaBinaryTypesDao struct {
	*sql.DB
	TableName string
}

// Query GomaBinaryTypesDao query
func (g GomaBinaryTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec GomaBinaryTypesDao exec
func (g GomaBinaryTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ GomaBinaryTypesDaoQueryer = (*GomaBinaryTypesDao)(nil)

// GomaBinaryTypes is GomaBinaryTypesDao.
func GomaBinaryTypes(db *sql.DB) GomaBinaryTypesDao {
	tblDao := GomaBinaryTypesDao{}
	tblDao.DB = db
	tblDao.TableName = "GomaBinaryTypes"
	return tblDao
}

// TxGomaBinaryTypesDao is generated goma_binary_types table transaction.
type TxGomaBinaryTypesDao struct {
	*sql.Tx
	TableName string
}

// Query TxGomaBinaryTypesDao query
func (g TxGomaBinaryTypesDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxGomaBinaryTypesDao exec
func (g TxGomaBinaryTypesDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ GomaBinaryTypesDaoQueryer = (*TxGomaBinaryTypesDao)(nil)

// TxGomaBinaryTypes is GomaBinaryTypesDao.
func TxGomaBinaryTypes(tx *sql.Tx) TxGomaBinaryTypesDao {
	tblDao := TxGomaBinaryTypesDao{}
	tblDao.Tx = tx
	tblDao.TableName = "GomaBinaryTypes"
	return tblDao
}

// SelectAll select goma_binary_types table all recode.
func (g GomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectAll(g)
}

// SelectAll transaction select goma_binary_types table all recode.
func (g TxGomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectAll(g)
}

func _GomaBinaryTypesSelectAll(g GomaBinaryTypesDaoQueryer) ([]entity.GomaBinaryTypesEntity, error) {
	queryString := queryArgs("goma_binary_types", "selectAll", nil)

	var es []entity.GomaBinaryTypesEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.GomaBinaryTypesEntity
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

// SelectByID select goma_binary_types table by primaryKey.
func (g GomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

// SelectByID transaction select goma_binary_types table by primaryKey.
func (g TxGomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypesEntity, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

func _GomaBinaryTypesSelectByID(g GomaBinaryTypesDaoQueryer, binaryID int64) (entity.GomaBinaryTypesEntity, error) {
	args := goma.QueryArgs{
		"binary_id": binaryID,
	}
	queryString := queryArgs("goma_binary_types", "selectByID", args)

	rows, err := g.Query(queryString)
	if err != nil {
		return entity.GomaBinaryTypesEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaBinaryTypesEntity{}, sql.ErrNoRows
	}

	var e entity.GomaBinaryTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaBinaryTypesEntity{}, err
	}

	return e, nil
}

// Insert insert goma_binary_types table.
func (g GomaBinaryTypesDao) Insert(e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

// Insert transaction insert goma_binary_types table.
func (g TxGomaBinaryTypesDao) Insert(e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

func _GomaBinaryTypesInsert(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"binary_id":          e.BinaryID,
		"binary_columns":     e.BinaryColumns,
		"tinyblob_columns":   e.TinyblobColumns,
		"blob_columns":       e.BlobColumns,
		"mediumblob_columns": e.MediumblobColumns,
		"longblob_columns":   e.LongblobColumns,
		"varbinary_columns":  e.VarbinaryColumns,
	}
	queryString := queryArgs("goma_binary_types", "insert", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (g GomaBinaryTypesDao) Update(e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update transaction update goma_binary_types table.
func (g TxGomaBinaryTypesDao) Update(e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update update goma_binary_types table.
func _GomaBinaryTypesUpdate(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"binary_id":          e.BinaryID,
		"binary_columns":     e.BinaryColumns,
		"tinyblob_columns":   e.TinyblobColumns,
		"blob_columns":       e.BlobColumns,
		"mediumblob_columns": e.MediumblobColumns,
		"longblob_columns":   e.LongblobColumns,
		"varbinary_columns":  e.VarbinaryColumns,
	}
	queryString := queryArgs("goma_binary_types", "update", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_binary_types table.
func (g GomaBinaryTypesDao) Delete(binaryID int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, binaryID)
}

// Delete transaction delete goma_binary_types table.
func (g TxGomaBinaryTypesDao) Delete(binaryID int64) (sql.Result, error) {
	return _GomaBinaryTypesDelete(g, binaryID)
}

// Delete delete goma_binary_types table by primaryKey.
func _GomaBinaryTypesDelete(g GomaBinaryTypesDaoQueryer, binaryID int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"binary_id": binaryID,
	}
	queryString := queryArgs("goma_binary_types", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
