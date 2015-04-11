package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// HelperTemplateData helper data.
type HelperTemplateData struct {
	PkgName      string
	DriverImport string
	DaoImport    string
	DaoPkgName   string
	DaoList      []DaoTemplateData
	Options      []map[string]interface{}
}

// DaoTemplateData dao data.
type DaoTemplateData struct {
	Name          string
	MemberName    string
	EntityName    string
	Table         TableTemplateData
	Imports       []string
	DaoPkgName    string
	EntityPkgName string
	EntityImport  string
}

// TableTemplateData table data.
type TableTemplateData struct {
	Name      string
	TitleName string
	Columns   []ColumnTemplateData
}

// ColumnTemplateData column data.
type ColumnTemplateData struct {
	Name         string
	TitleName    string
	TypeName     string
	TypeDetail   string
	IsPrimaryKey bool
	Sample       string
}

type AssetTemplateData struct {
	DaoPkgName string
}

type QueryArgsTemplateData struct {
	DaoPkgName string
	SQLRootDir string
}

func (d AssetTemplateData) execAssetTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := AssetTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, "asset_gen.go", buf.Bytes())
}

func (d QueryArgsTemplateData) execQueryArgsTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := QueryArgsTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, "queryargs_gen.go", buf.Bytes())
}

func (d HelperTemplateData) execHelperTemplate(rootDir string) error {
	//	var buf bytes.Buffer
	//	if err := HelperTemplate(&buf, d); err != nil {
	//		return err
	//	}
	//	return formatFileWrite(rootDir, "gomautils_gen.go", buf.Bytes())
	return nil
}

func (d DaoTemplateData) execDaoTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := DaoTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, d.Table.Name+"_gen.go", buf.Bytes())
}

func (d DaoTemplateData) execEntityTemplate(entityRootDir string) error {
	var buf bytes.Buffer
	if err := EntityTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(entityRootDir, d.Table.Name+"_gen.go", buf.Bytes())
}

func (t TableTemplateData) execTableTemplate(sqlRootDir string) error {
	sqlTableDir := filepath.Join(sqlRootDir, t.Name)
	if err := os.MkdirAll(sqlTableDir, 0755); err != nil {
		return err
	}

	var err error
	var buf bytes.Buffer
	fileWriteFunc := func(templateFunc func(io.Writer, TableTemplateData) error, fileName string) {
		if err != nil {
			return
		}

		buf.Reset()
		err = templateFunc(&buf, t)
		if err != nil {
			return
		}

		filePath := filepath.Join(sqlTableDir, fileName)
		err = ioutil.WriteFile(filePath, buf.Bytes(), 0644)
		if err != nil {
			err = fmt.Errorf("file write error: %s \n%s", err, filePath)
			return
		}

		debugPrintln("generate file:", filePath)
	}

	fileWriteFunc(SelectAllTemplate, "selectAll.sql")
	fileWriteFunc(SelectByIDTemplate, "selectByID.sql")
	fileWriteFunc(InsertTemplate, "insert.sql")
	fileWriteFunc(UpdateTemplate, "update.sql")
	fileWriteFunc(DeleteTemplate, "delete.sql")

	return err
}

func formatFileWrite(path, fileName string, data []byte) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	bts, err := format.Source(data)
	if err != nil {
		return fmt.Errorf("go format error: %s \n%s", err, string(data))
	}

	filePath := filepath.Join(path, fileName)
	if err := ioutil.WriteFile(filePath, bts, 0644); err != nil {
		return fmt.Errorf("file write error: %s \n%s", err, filePath)
	}

	debugPrintln("generate file:", filePath)
	return nil
}

// Println Debug Println
func debugPrintln(v ...interface{}) {
	// SliceInsert (https://code.google.com/p/go-wiki/wiki/SliceTricks)
	v = append(v[:0], append([]interface{}{"[goma]", "[debug]"}, v[0:]...)...)

	fmt.Println(v...)
}
