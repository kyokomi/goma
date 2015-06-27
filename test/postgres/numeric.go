package main

import (
	"database/sql"
	"log"

	"github.com/kyokomi/goma/test/postgres/dao"
)

func gomaNumericTypesDeleteAll(g dao.GomaNumericTypesDaoQueryer) (sql.Result, error) {
	queryString, args, err := dao.GenerateQuery("goma_numeric_types", "deleteAll", nil)
	if err != nil {
		return nil, err
	}

	result, err := g.Exec(queryString, args...)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
