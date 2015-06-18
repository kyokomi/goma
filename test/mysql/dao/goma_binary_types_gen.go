package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/mysql/entity"
)

var tableGomaBinaryTypes = "goma_binary_types"
var columnsGomaBinaryTypes = []string{
	"binary_id",
	"binary_columns",
	"tinyblob_columns",
	"blob_columns",
	"mediumblob_columns",
	"longblob_columns",
	"varbinary_columns",
}

// GomaBinaryTypesTableName goma_binary_types table name
func GomaBinaryTypesTableName() string {
	return tableGomaBinaryTypes
}

// GomaBinaryTypesTableColumns goma_binary_types table columns
func GomaBinaryTypesTableColumns() []string {
	return columnsGomaBinaryTypes
}

// GomaBinaryTypesDaoQueryer is interface
type GomaBinaryTypesDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// GomaBinaryTypesDao is generated goma_binary_types table.
type GomaBinaryTypesDao struct {
	*sql.DB
	TableName string
	Columns   []string
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
	tblDao.TableName = tableGomaBinaryTypes
	tblDao.Columns = columnsGomaBinaryTypes
	return tblDao
}

// TxGomaBinaryTypesDao is generated goma_binary_types table transaction.
type TxGomaBinaryTypesDao struct {
	*sql.Tx
	TableName string
	Columns   []string
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
	tblDao.TableName = tableGomaBinaryTypes
	tblDao.Columns = columnsGomaBinaryTypes
	return tblDao
}

// SelectAll select goma_binary_types table all recode.
func (g GomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectAll(g)
}

// SelectAll transaction select goma_binary_types table all recode.
func (g TxGomaBinaryTypesDao) SelectAll() ([]entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectAll(g)
}

func _GomaBinaryTypesSelectAll(g GomaBinaryTypesDaoQueryer) ([]entity.GomaBinaryTypes, error) {
	queryString, args, err := queryArgs("goma_binary_types", "selectAll", nil)
	if err != nil {
		return nil, err
	}

	var es []entity.GomaBinaryTypes
	rows, err := g.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.GomaBinaryTypes
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
func (g GomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

// SelectByID transaction select goma_binary_types table by primaryKey.
func (g TxGomaBinaryTypesDao) SelectByID(binaryID int64) (entity.GomaBinaryTypes, error) {
	return _GomaBinaryTypesSelectByID(g, binaryID)
}

func _GomaBinaryTypesSelectByID(g GomaBinaryTypesDaoQueryer, binaryID int64) (entity.GomaBinaryTypes, error) {
	argsMap := map[string]interface{}{
		"binary_id": binaryID,
	}
	queryString, args, err := queryArgs("goma_binary_types", "selectByID", argsMap)
	if err != nil {
		return entity.GomaBinaryTypes{}, err
	}

	rows, err := g.Query(queryString, args...)
	if err != nil {
		return entity.GomaBinaryTypes{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.GomaBinaryTypes{}, sql.ErrNoRows
	}

	var e entity.GomaBinaryTypes
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.GomaBinaryTypes{}, err
	}

	return e, nil
}

// Insert insert goma_binary_types table.
func (g GomaBinaryTypesDao) Insert(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

// Insert transaction insert goma_binary_types table.
func (g TxGomaBinaryTypesDao) Insert(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesInsert(g, e)
}

func _GomaBinaryTypesInsert(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"binary_id":          e.BinaryID,
		"binary_columns":     e.BinaryColumns,
		"tinyblob_columns":   e.TinyblobColumns,
		"blob_columns":       e.BlobColumns,
		"mediumblob_columns": e.MediumblobColumns,
		"longblob_columns":   e.LongblobColumns,
		"varbinary_columns":  e.VarbinaryColumns,
	}
	queryString, args, err := queryArgs("goma_binary_types", "insert", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (g GomaBinaryTypesDao) Update(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update transaction update goma_binary_types table.
func (g TxGomaBinaryTypesDao) Update(e entity.GomaBinaryTypes) (sql.Result, error) {
	return _GomaBinaryTypesUpdate(g, e)
}

// Update update goma_binary_types table.
func _GomaBinaryTypesUpdate(g GomaBinaryTypesDaoQueryer, e entity.GomaBinaryTypes) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"binary_id":          e.BinaryID,
		"binary_columns":     e.BinaryColumns,
		"tinyblob_columns":   e.TinyblobColumns,
		"blob_columns":       e.BlobColumns,
		"mediumblob_columns": e.MediumblobColumns,
		"longblob_columns":   e.LongblobColumns,
		"varbinary_columns":  e.VarbinaryColumns,
	}
	queryString, args, err := queryArgs("goma_binary_types", "update", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
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
	argsMap := map[string]interface{}{
		"binary_id": binaryID,
	}
	queryString, args, err := queryArgs("goma_binary_types", "delete", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
