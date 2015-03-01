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

// QuestDao is generated quest table.
type QuestDao struct {
	*goma.Goma
	tx        *sql.Tx
	TableName string
}

// Quest is QuestDao.
func Quest(g *goma.Goma) QuestDao {
	tblDao := QuestDao{}
	tblDao.Goma = g
	tblDao.tx = nil
	tblDao.TableName = "Quest"
	return tblDao
}

// IsTx started transaction?
func (d QuestDao) IsTx() bool {
	return d.tx != nil
}

// SetTx set transaction.
func (d *QuestDao) SetTx(tx *sql.Tx) {
	d.tx = tx
}

// ResetTx reset transaction.
// Call after Commit of Rollback.
func (d *QuestDao) ResetTx() {
	d.tx = nil
}

func (d QuestDao) daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if d.IsTx() {
		rows, err = d.tx.Query(query, args...)
	} else {
		rows, err = d.Query(query, args...)
	}
	return
}

func (d QuestDao) daoExec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.IsTx() {
		result, err = d.tx.Exec(query, args...)
	} else {
		result, err = d.Exec(query, args...)
	}
	return
}

// SelectAll select quest table all recode.
func (d QuestDao) SelectAll() ([]*entity.QuestEntity, error) {
	queryString := d.QueryArgs("quest", "selectAll", nil)

	var entitys []*entity.QuestEntity
	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity entity.QuestEntity
		err = rows.Scan(&entity.ID, &entity.Name, &entity.Detail, &entity.CreateAt)
		if err != nil {
			break
		}

		entitys = append(entitys, &entity)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return entitys, nil
}

// SelectByID select quest table by primaryKey.
func (d QuestDao) SelectByID(id int) (*entity.QuestEntity, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("quest", "selectByID", args)

	rows, err := d.daoQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity entity.QuestEntity
	if err := rows.Scan(&entity.ID, &entity.Name, &entity.Detail, &entity.CreateAt); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert quest table.
func (d QuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("quest", "insert", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (d QuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("quest", "update", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete quest table by primaryKey.
func (d QuestDao) Delete(id int) (sql.Result, error) {
	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("quest", "delete", args)

	result, err := d.daoExec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
