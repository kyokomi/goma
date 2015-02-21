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
		Driver:     "mysql",
		UserName:   "admin",
		PassWord:   "password",
		Host:       "localhost",
		Port:       3306,
		DBName:     "test",
		Debug:      true,
		SQLRootDir: "./sql",
		DaoRootDir: "./dao",
	}

	return goma.NewGoma(opts)
}
