package goma

import "github.com/jmoiron/sqlx"

func GenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	if len(argsMap) <= 0 {
		return queryString, nil, nil
	}
	return sqlx.Named(queryString, argsMap)
}

// MySQLGenerateQuery generate bind args query
func MySQLGenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	return GenerateQuery(queryString, argsMap)
}

// PostgresGenerateQuery generate bind args query
func PostgresGenerateQuery(queryString string, argsMap map[string]interface{}) (string, []interface{}, error) {
	query, args, err := GenerateQuery(queryString, argsMap)
	return sqlx.Rebind(sqlx.DOLLAR, query), args, err
}
