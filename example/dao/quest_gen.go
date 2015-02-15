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

// NewQuestDao is QuestDao singleton.
func Quest(g *goma.Goma) *QuestDao {
	if quest == nil {
		quest = &QuestDao{Goma: g}
	}
	return quest
}

// QuestEntity is generated quest table.
type QuestEntity struct {
	ID     int
	Name   string
	Detail string
}

// TODO: あとで消す
const questSelectAll = `
select
  *
FROM
  quest
`

func (d *QuestDao) SelectAll() ([]*QuestEntity, error) {

	var entitys []*QuestEntity
	rows, err := d.Query(questSelectAll)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity QuestEntity
		err = rows.Scan(&entity.ID, &entity.Name, &entity.Detail)
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

// TODO: あとで消す
const questSelectByID = `
select
  *
FROM
  quest
WHERE
  id = /* id */1
and
  name = /* name */"hoge"
`

func (d *QuestDao) SelectByID(args goma.QueryArgs) (*QuestEntity, error) {

	// TODO: sqlファイルを読み込む（Dao生成時にmethod名、引数をKeyにmapで保持する
	queryString := d.QueryArgs(questSelectByID, args)

	var entity QuestEntity
	err := d.QueryRow(queryString).Scan(&entity.ID, &entity.Name, &entity.Detail)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
