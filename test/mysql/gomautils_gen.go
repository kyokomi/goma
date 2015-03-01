package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma/test/mysql/dao"

	"github.com/kyokomi/goma"
)

// Goma goma.Goma utils.
type Goma struct {
	*goma.Goma
}

// NewGoma is goma.Goma wrapper utils.
func NewGoma() (Goma, error) {
	opts := goma.Options{
		Driver:        "mysql",
		UserName:      "admin",
		PassWord:      "password",
		Host:          "localhost",
		Port:          3306,
		DBName:        "goma_test",
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

// GomaBinaryTypes is dao.GomaBinaryTypesDao helper.
func (g Goma) GomaBinaryTypes() dao.GomaBinaryTypesDao {
	return dao.GomaBinaryTypes(g.Goma)
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
