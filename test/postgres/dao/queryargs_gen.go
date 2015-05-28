package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"path/filepath"

	"github.com/kyokomi/goma"
)

type queryArgSettings struct {
	rootDir string
}

var settings queryArgSettings

func (s *queryArgSettings) SetRootDir(rootDir string) {
	s.rootDir = rootDir
}

func queryArgs(tableName string, queryName string, args goma.QueryArgs) string {
	return settings.queryArgs(tableName, queryName, args)
}

func (s queryArgSettings) queryArgs(tableName string, queryName string, args goma.QueryArgs) string {
	filePath := createSqlFilePath(s.rootDir, tableName, queryName)
	return goma.PostgresGenerateQuery(assetSQL(filePath), args)
}

func assetSQL(filePath string) string {
	data, err := Asset(filePath)
	if err != nil {
		// Asset was not found.
	}
	return string(data)
}

func createSqlFilePath(rootDir string, tableName string, queryName string) string {
	return filepath.Join(rootDir, "sql", tableName, queryName+".sql")
}
