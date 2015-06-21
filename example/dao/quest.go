package dao

import (
	"log"

	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kyokomi/goma/example/entity"
)

// SelectLimit select quest table by primaryKey.
func (g QuestDao) SelectLimit(limit int) ([]entity.Quest, error) {
	argsMap := map[string]interface{}{
		"limit": limit,
	}

	queryString, args, err := queryArgs("quest", "selectLimit", argsMap)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(queryString, args)

	// TODO: DriverNameなんとかしたい
	rows, err := sqlx.NewDb(g.DB, "mysql").Queryx(queryString, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	es := []entity.Quest{}
	for rows.Next() {
		e := entity.Quest{}
		if err := rows.StructScan(&e); err != nil {
			log.Println(err)
			continue
		}
		es = append(es, e)
	}

	return es, nil
}
