package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"io/ioutil"
)

func init() {
	settings.sqlFile = true
}

// Asset default read file
func Asset(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
