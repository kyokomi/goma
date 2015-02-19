package dao

// GENERATE CODE

import (
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
    Id int
    Name string
    Detail string
}

func (d *QuestDao) SelectAll() ([]*QuestEntity, error) {

    queryString := d.QueryArgs("quest", "selectAll", nil)

	var entitys []*QuestEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity QuestEntity
		err = rows.Scan(&entity.Id, &entity.Name, &entity.Detail)
		if err != nil {
			break
		}

		entitys = append(entitys, &entity)
	}
	if err != nil {
		return nil, err
	}

	return entitys, nil
}

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
	if err := d.QueryRow(queryString).Scan(&entity.Id, &entity.Name, &entity.Detail); err != nil {
		return nil, err
	}
	
	return &entity, nil
}
