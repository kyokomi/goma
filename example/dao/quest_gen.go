package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/example/entity"
)

var tableQuest = "quest"
var columnsQuest = []string{
	"id",
	"name",
	"detail",
	"create_at",
}

// QuestTableName quest table name
func QuestTableName() string {
	return tableQuest
}

// QuestTableColumns quest table columns
func QuestTableColumns() []string {
	return columnsQuest
}

// QuestDaoQueryer is interface
type QuestDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// QuestDao is generated quest table.
type QuestDao struct {
	*sql.DB
	TableName string
	Columns   []string
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
	tblDao.TableName = tableQuest
	tblDao.Columns = columnsQuest
	return tblDao
}

// TxQuestDao is generated quest table transaction.
type TxQuestDao struct {
	*sql.Tx
	TableName string
	Columns   []string
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
	tblDao.TableName = tableQuest
	tblDao.Columns = columnsQuest
	return tblDao
}

// SelectAll select quest table all recode.
func (g QuestDao) SelectAll() ([]entity.Quest, error) {
	return _QuestSelectAll(g)
}

// SelectAll transaction select quest table all recode.
func (g TxQuestDao) SelectAll() ([]entity.Quest, error) {
	return _QuestSelectAll(g)
}

func _QuestSelectAll(g QuestDaoQueryer) ([]entity.Quest, error) {
	queryString, args, err := queryArgs("quest", "selectAll", nil)
	if err != nil {
		return nil, err
	}

	var es []entity.Quest
	rows, err := g.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e entity.Quest
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
func (g QuestDao) SelectByID(id int) (entity.Quest, error) {
	return _QuestSelectByID(g, id)
}

// SelectByID transaction select quest table by primaryKey.
func (g TxQuestDao) SelectByID(id int) (entity.Quest, error) {
	return _QuestSelectByID(g, id)
}

func _QuestSelectByID(g QuestDaoQueryer, id int) (entity.Quest, error) {
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("quest", "selectByID", argsMap)
	if err != nil {
		return entity.Quest{}, err
	}

	rows, err := g.Query(queryString, args...)
	if err != nil {
		return entity.Quest{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Quest{}, sql.ErrNoRows
	}

	var e entity.Quest
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.Quest{}, err
	}

	return e, nil
}

// Insert insert quest table.
func (g QuestDao) Insert(e entity.Quest) (sql.Result, error) {
	return _QuestInsert(g, e)
}

// Insert transaction insert quest table.
func (g TxQuestDao) Insert(e entity.Quest) (sql.Result, error) {
	return _QuestInsert(g, e)
}

func _QuestInsert(g QuestDaoQueryer, e entity.Quest) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":        e.ID,
		"name":      e.Name,
		"detail":    e.Detail,
		"create_at": e.CreateAt,
	}
	queryString, args, err := queryArgs("quest", "insert", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (g QuestDao) Update(e entity.Quest) (sql.Result, error) {
	return _QuestUpdate(g, e)
}

// Update transaction update quest table.
func (g TxQuestDao) Update(e entity.Quest) (sql.Result, error) {
	return _QuestUpdate(g, e)
}

// Update update quest table.
func _QuestUpdate(g QuestDaoQueryer, e entity.Quest) (sql.Result, error) {
	argsMap := map[string]interface{}{
		"id":        e.ID,
		"name":      e.Name,
		"detail":    e.Detail,
		"create_at": e.CreateAt,
	}
	queryString, args, err := queryArgs("quest", "update", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
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
	argsMap := map[string]interface{}{
		"id": id,
	}
	queryString, args, err := queryArgs("quest", "delete", argsMap)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
