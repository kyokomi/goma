package dao

import (
	"database/sql"
	"log"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/example/entity"
)

// UpdateAll update quest table.
func (d *QuestDao) UpdateAll(entity entity.QuestEntity) (sql.Result, error) {
	args := goma.QueryArgs{
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := queryArgs("quest", "updateAll", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
