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

// AssetFile default read file
func AssetFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
