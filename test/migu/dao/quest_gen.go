package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/migu/models"
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
func (g QuestDao) SelectAll() ([]models.Quest, error) {
	return _QuestSelectAll(g)
}

// SelectAll transaction select quest table all recode.
func (g TxQuestDao) SelectAll() ([]models.Quest, error) {
	return _QuestSelectAll(g)
}

func _QuestSelectAll(g QuestDaoQueryer) ([]models.Quest, error) {
	queryString := `
select
  id
, title
, detail
, create_at
, update_at
FROM
  quest`

	var es []models.Quest
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e models.Quest
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
func (g QuestDao) SelectByID(id int64) (models.Quest, error) {
	return _QuestSelectByID(g, id)
}

// SelectByID transaction select quest table by primaryKey.
func (g TxQuestDao) SelectByID(id int64) (models.Quest, error) {
	return _QuestSelectByID(g, id)
}

func _QuestSelectByID(g QuestDaoQueryer, id int64) (models.Quest, error) {
	queryString := `
select
  id
, title
, detail
, create_at
, update_at
FROM
  quest
WHERE
  id = ?
`
	rows, err := g.Query(queryString,
		id,
	)
	if err != nil {
		return models.Quest{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return models.Quest{}, sql.ErrNoRows
	}

	var e models.Quest
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return models.Quest{}, err
	}

	return e, nil
}

// Insert insert quest table.
func (g QuestDao) Insert(e models.Quest) (sql.Result, error) {
	return _QuestInsert(g, e)
}

// Insert transaction insert quest table.
func (g TxQuestDao) Insert(e models.Quest) (sql.Result, error) {
	return _QuestInsert(g, e)
}

func _QuestInsert(g QuestDaoQueryer, e models.Quest) (sql.Result, error) {
	queryString := `
insert into quest(
  id
, title
, detail
, create_at
, update_at
) values(
  ?
, ?
, ?
, ?
, ?
)`
	result, err := g.Exec(queryString,
		e.ID,
		e.Title,
		e.Detail,
		e.CreateAt,
		e.UpdateAt,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (g QuestDao) Update(e models.Quest) (sql.Result, error) {
	return _QuestUpdate(g, e)
}

// Update transaction update quest table.
func (g TxQuestDao) Update(e models.Quest) (sql.Result, error) {
	return _QuestUpdate(g, e)
}

// Update update quest table.
func _QuestUpdate(g QuestDaoQueryer, e models.Quest) (sql.Result, error) {
	queryString := `
update quest set
    id = ?
,   title = ?
,   detail = ?
,   create_at = ?
,   update_at = ?
 where
    id = ?

`
	result, err := g.Exec(queryString,
		e.ID,
		e.Title,
		e.Detail,
		e.CreateAt,
		e.UpdateAt,

		e.ID,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete quest table.
func (g QuestDao) Delete(id int64) (sql.Result, error) {
	return _QuestDelete(g, id)
}

// Delete transaction delete quest table.
func (g TxQuestDao) Delete(id int64) (sql.Result, error) {
	return _QuestDelete(g, id)
}

// Delete delete quest table by primaryKey.
func _QuestDelete(g QuestDaoQueryer, id int64) (sql.Result, error) {
	queryString := `
delete
from
  quest
where
  id = ?

`
	result, err := g.Exec(queryString,
		id,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
