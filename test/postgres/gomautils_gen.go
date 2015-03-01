package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"os"

	_ "github.com/lib/pq"

	"github.com/kyokomi/goma/test/postgres/dao"

	"github.com/kyokomi/goma"
)

// Goma goma.Goma utils.
type Goma struct {
	*goma.Goma
}

// NewGoma is goma.Goma wrapper utils.
func NewGoma() (Goma, error) {
	opts := goma.Options{
		Driver:        "postgres",
		UserName:      "admin",
		PassWord:      "password",
		Host:          "localhost",
		Port:          5432,
		DBName:        "goma-test",
		SSLMode:       "disable",
		Debug:         true,
		SQLRootDir:    "sql",
		DaoRootDir:    "dao",
		EntityRootDir: "entity",
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return Goma{}, err
	}
	opts.CurrentDir = currentDir

	g, err := goma.NewGoma(opts)
	if err != nil {
		return Goma{}, err
	}

	gm := Goma{}
	gm.Goma = g
	return gm, nil
}

// GomaDateTypes is dao.GomaDateTypesDao helper.
func (g Goma) GomaDateTypes() dao.GomaDateTypesDao {
	return dao.GomaDateTypes(g.Goma)
}

// GomaNumericTypes is dao.GomaNumericTypesDao helper.
func (g Goma) GomaNumericTypes() dao.GomaNumericTypesDao {
	return dao.GomaNumericTypes(g.Goma)
}

// GomaStringTypes is dao.GomaStringTypesDao helper.
func (g Goma) GomaStringTypes() dao.GomaStringTypesDao {
	return dao.GomaStringTypes(g.Goma)
}
