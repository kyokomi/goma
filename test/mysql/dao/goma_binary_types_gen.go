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

// GomaBinaryTypesDaoTx is generated goma_binary_types table transaction.
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
func (d GomaBinaryTypesDao) SelectAll() ([]*entity.GomaBinaryTypesEntity, error) {
	return goma_binary_typesSelectAll(d)
}

// SelectAll transaction select goma_binary_types table all recode.
func (d TxGomaBinaryTypesDao) SelectAll() ([]*entity.GomaBinaryTypesEntity, error) {
	return goma_binary_typesSelectAll(d)
}

func goma_binary_typesSelectAll(d GomaBinaryTypesDaoQueryer) ([]*entity.GomaBinaryTypesEntity, error) {
	queryString := queryArgs("goma_binary_types", "selectAll", nil)

	var es []*entity.GomaBinaryTypesEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.GomaBinaryTypesEntity
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

// SelectByID select goma_binary_types table by primaryKey.
func (d GomaBinaryTypesDao) SelectByID(id int64) (*entity.GomaBinaryTypesEntity, error) {
	return goma_binary_typesSelectByID(d, id)
}

// SelectByID transaction select goma_binary_types table by primaryKey.
func (d TxGomaBinaryTypesDao) SelectByID(id int64) (*entity.GomaBinaryTypesEntity, error) {
	return goma_binary_typesSelectByID(d, id)
}

func goma_binary_typesSelectByID(d GomaBinaryTypesDaoQueryer, id int64) (*entity.GomaBinaryTypesEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_binary_types", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.GomaBinaryTypesEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert goma_binary_types table.
func (d GomaBinaryTypesDao) Insert(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return goma_binary_typesInsert(d, entity)
}

// Insert transaction insert goma_binary_types table.
func (d TxGomaBinaryTypesDao) Insert(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return goma_binary_typesInsert(d, entity)
}

func goma_binary_typesInsert(d GomaBinaryTypesDaoQueryer, entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"binary_columns":     entity.BinaryColumns,
		"tinyblob_columns":   entity.TinyblobColumns,
		"blob_columns":       entity.BlobColumns,
		"mediumblob_columns": entity.MediumblobColumns,
		"longblob_columns":   entity.LongblobColumns,
		"varbinary_columns":  entity.VarbinaryColumns,
	}
	queryString := queryArgs("goma_binary_types", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update goma_binary_types table.
func (d GomaBinaryTypesDao) Update(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return goma_binary_typesUpdate(d, entity)
}

// Update transaction update goma_binary_types table.
func (d TxGomaBinaryTypesDao) Update(entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	return goma_binary_typesUpdate(d, entity)
}

// Update update goma_binary_types table.
func goma_binary_typesUpdate(d GomaBinaryTypesDaoQueryer, entity entity.GomaBinaryTypesEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":                 entity.ID,
		"binary_columns":     entity.BinaryColumns,
		"tinyblob_columns":   entity.TinyblobColumns,
		"blob_columns":       entity.BlobColumns,
		"mediumblob_columns": entity.MediumblobColumns,
		"longblob_columns":   entity.LongblobColumns,
		"varbinary_columns":  entity.VarbinaryColumns,
	}
	queryString := queryArgs("goma_binary_types", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete goma_binary_types table.
func (d GomaBinaryTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_binary_typesDelete(d, id)
}

// Delete transaction delete goma_binary_types table.
func (d TxGomaBinaryTypesDao) Delete(id int64) (sql.Result, error) {
	return goma_binary_typesDelete(d, id)
}

// Delete delete goma_binary_types table by primaryKey.
func goma_binary_typesDelete(d GomaBinaryTypesDaoQueryer, id int64) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("goma_binary_types", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
