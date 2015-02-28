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
		Driver:     "mysql",
		UserName:   "admin",
		PassWord:   "password",
		Host:       "localhost",
		Port:       3306,
		DBName:     "goma_test",
		Debug:      true,
		SQLRootDir: "sql",
		DaoRootDir: "dao",
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

func (g Goma) GomaBinaryTypes() dao.GomaBinaryTypesDao {
	return dao.GomaBinaryTypes(g.Goma)
}
func (g Goma) GomaDateTypes() dao.GomaDateTypesDao {
	return dao.GomaDateTypes(g.Goma)
}
func (g Goma) GomaNumericTypes() dao.GomaNumericTypesDao {
	return dao.GomaNumericTypes(g.Goma)
}
func (g Goma) GomaStringTypes() dao.GomaStringTypesDao {
	return dao.GomaStringTypes(g.Goma)
}
