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

	// dao
	Quest  *dao.QuestDao
	Sample *dao.SampleDao
}

// NewGoma is goma.Goma wrapper utils.
func NewGoma() (Goma, error) {

	opts := goma.Options{
		Driver:     "mysql",
		UserName:   "admin",
		PassWord:   "password",
		Host:       "localhost",
		Port:       3306,
		DBName:     "test",
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
	gm.Quest = dao.Quest(g)
	gm.Sample = dao.Sample(g)

	return gm, nil
}
