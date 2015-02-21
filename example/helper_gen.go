package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/kyokomi/goma/goma"
)

func Goma() (*goma.Goma, error) {

	opts := goma.Options{
		Debug:      "true",
		Driver:     "mysql",
		UserName:   "admin",
		PassWord:   "password",
		Port:       3306,
		Host:       "localhost",
		DBName:     "test",
		SQLRootDir: "./sql",
		DaoRootDir: "./dao",
	}

	return goma.NewGoma(opts)
}
