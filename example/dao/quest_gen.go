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

// QuestDaoQueryer is interface
type QuestDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// QuestDao is generated quest table.
type QuestDao struct {
	*sql.DB
	TableName string
}

// Query QuestDao query
func (g QuestDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec QuestDao exec
func (g QuestDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ QuestDaoQueryer = (*QuestDao)(nil)

// Quest is QuestDao.
func Quest(db *sql.DB) QuestDao {
	tblDao := QuestDao{}
	tblDao.DB = db
	tblDao.TableName = "Quest"
	return tblDao
}

// TxQuestDao is generated quest table transaction.
type TxQuestDao struct {
	*sql.Tx
	TableName string
}

// Query TxQuestDao query
func (g TxQuestDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxQuestDao exec
func (g TxQuestDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ QuestDaoQueryer = (*TxQuestDao)(nil)

// TxQuest is QuestDao.
func TxQuest(tx *sql.Tx) TxQuestDao {
	tblDao := TxQuestDao{}
	tblDao.Tx = tx
	tblDao.TableName = "Quest"
	return tblDao
}

// SelectAll select quest table all recode.
func (g QuestDao) SelectAll() ([]entity.QuestEntity, error) {
	return _QuestSelectAll(g)
}

// SelectAll transaction select quest table all recode.
func (g TxQuestDao) SelectAll() ([]entity.QuestEntity, error) {
	return _QuestSelectAll(g)
}

func _QuestSelectAll(g QuestDaoQueryer) ([]entity.QuestEntity, error) {
	queryString := queryArgs("quest", "selectAll", nil)

	var es []entity.QuestEntity
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	for rows.Next() {
		var e entity.QuestEntity
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

// SelectByID select quest table by primaryKey.
func (g QuestDao) SelectByID(id int) (entity.QuestEntity, error) {
	return _QuestSelectByID(g, id)
}

// SelectByID transaction select quest table by primaryKey.
func (g TxQuestDao) SelectByID(id int) (entity.QuestEntity, error) {
	return _QuestSelectByID(g, id)
}

func _QuestSelectByID(g QuestDaoQueryer, id int) (entity.QuestEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("quest", "selectByID", args)

	rows, err := g.Query(queryString)
	if err != nil {
		return entity.QuestEntity{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.QuestEntity{}, sql.ErrNoRows
	}

	var e entity.QuestEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.QuestEntity{}, err
	}

	return e, nil
}

// Insert insert quest table.
func (g QuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {
	return _QuestInsert(g, entity)
}

// Insert transaction insert quest table.
func (g TxQuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {
	return _QuestInsert(g, entity)
}

func _QuestInsert(g QuestDaoQueryer, entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("quest", "insert", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (g QuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {
	return _QuestUpdate(g, entity)
}

// Update transaction update quest table.
func (g TxQuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {
	return _QuestUpdate(g, entity)
}

// Update update quest table.
func _QuestUpdate(g QuestDaoQueryer, entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("quest", "update", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete quest table.
func (g QuestDao) Delete(id int) (sql.Result, error) {
	return _QuestDelete(g, id)
}

// Delete transaction delete quest table.
func (g TxQuestDao) Delete(id int) (sql.Result, error) {
	return _QuestDelete(g, id)
}

// Delete delete quest table by primaryKey.
func _QuestDelete(g QuestDaoQueryer, id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("quest", "delete", args)

	result, err := g.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
