package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma/example/dao"

	"github.com/kyokomi/goma"
)

// Goma goma.Goma utils.
type Goma struct {
	*goma.Goma
}

// NewGoma is goma.Goma wrapper utils.
func NewGoma(configPath string) (Goma, error) {
	opts, err := goma.NewOptions(configPath)
	if err != nil {
		return Goma{}, err
	}
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
