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

// QuestDaoTx is generated quest table transaction.
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
func (d QuestDao) SelectAll() ([]*entity.QuestEntity, error) {
	return questSelectAll(d)
}

// SelectAll transaction select quest table all recode.
func (d TxQuestDao) SelectAll() ([]*entity.QuestEntity, error) {
	return questSelectAll(d)
}

func questSelectAll(d QuestDaoQueryer) ([]*entity.QuestEntity, error) {
	queryString := queryArgs("quest", "selectAll", nil)

	var es []*entity.QuestEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.QuestEntity
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

// SelectByID select quest table by primaryKey.
func (d QuestDao) SelectByID(id int) (*entity.QuestEntity, error) {
	return questSelectByID(d, id)
}

// SelectByID transaction select quest table by primaryKey.
func (d TxQuestDao) SelectByID(id int) (*entity.QuestEntity, error) {
	return questSelectByID(d, id)
}

func questSelectByID(d QuestDaoQueryer, id int) (*entity.QuestEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("quest", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var e entity.QuestEntity
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &e, nil
}

// Insert insert quest table.
func (d QuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {
	return questInsert(d, entity)
}

// Insert transaction insert quest table.
func (d TxQuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {
	return questInsert(d, entity)
}

func questInsert(d QuestDaoQueryer, entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("quest", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (d QuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {
	return questUpdate(d, entity)
}

// Update transaction update quest table.
func (d TxQuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {
	return questUpdate(d, entity)
}

// Update update quest table.
func questUpdate(d QuestDaoQueryer, entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("quest", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete quest table.
func (d QuestDao) Delete(id int) (sql.Result, error) {
	return questDelete(d, id)
}

// Delete transaction delete quest table.
func (d TxQuestDao) Delete(id int) (sql.Result, error) {
	return questDelete(d, id)
}

// Delete delete quest table by primaryKey.
func questDelete(d QuestDaoQueryer, id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := queryArgs("quest", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
