package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/example/entity"

	"github.com/kyokomi/goma/goma"
)

// QuestDao is generated quest table.
type QuestDao struct {
	*goma.Goma
	TableName string
}

var sQuest *QuestDao

// Quest is QuestDao singleton.
func Quest(g *goma.Goma) *QuestDao {
	if sQuest == nil {
		sQuest = &QuestDao{
			Goma:      g,
			TableName: "Quest",
		}
	}
	return sQuest
}

// SelectAll select quest table all recode.
func (d *QuestDao) SelectAll() ([]*entity.QuestEntity, error) {

	queryString := d.QueryArgs("quest", "selectAll", nil)

	var entitys []*entity.QuestEntity
	rows, err := d.Query(queryString)
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
func (d *QuestDao) SelectByID(id int) (*entity.QuestEntity, error) {

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

	var entity entity.QuestEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.Name, &entity.Detail, &entity.CreateAt); err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return &entity, nil
}

// Insert insert quest table.
func (d *QuestDao) Insert(entity entity.QuestEntity) (sql.Result, error) {

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
func (d *QuestDao) Update(entity entity.QuestEntity) (sql.Result, error) {

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
