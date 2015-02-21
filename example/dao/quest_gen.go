package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"time"

	"database/sql"

	"github.com/kyokomi/goma/goma"
)

// QuestDao is generated quest table.
type QuestDao struct {
	*goma.Goma
}

var quest *QuestDao

// Quest is QuestDao singleton.
func Quest(g *goma.Goma) *QuestDao {
	if quest == nil {
		quest = &QuestDao{Goma: g}
	}
	return quest
}

// QuestEntity is generated quest table.
type QuestEntity struct {
	ID       int
	Name     string
	Detail   string
	CreateAt time.Time
}

// SelectAll select quest table all recode.
func (d *QuestDao) SelectAll() ([]*QuestEntity, error) {

	queryString := d.QueryArgs("quest", "selectAll", nil)

	var entitys []*QuestEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity QuestEntity
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
func (d *QuestDao) SelectByID(id int) (*QuestEntity, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("quest", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity QuestEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.Name, &entity.Detail, &entity.CreateAt); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert quest table.
func (d *QuestDao) Insert(entity QuestEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("quest", "insert", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update quest table.
func (d *QuestDao) Update(entity QuestEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":        entity.ID,
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("quest", "update", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete quest table by primaryKey.
func (d *QuestDao) Delete(id int) (sql.Result, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("quest", "delete", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
