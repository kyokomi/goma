package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma/example/dao"

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
		DBName:        "test",
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

// Quest is dao.QuestDao helper.
func (g Goma) Quest() dao.QuestDao {
	return dao.Quest(g.Goma)
}

// Sample is dao.SampleDao helper.
func (g Goma) Sample() dao.SampleDao {
	return dao.Sample(g.Goma)
}
