package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/etgryphon/stringUp"
)

// HelperTemplateData helper data.
type HelperTemplateData struct {
	PkgName      string
	DriverImport string
	DaoImport    string
	DaoPkgName   string
	DriverName   string
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
	DriverName    string
	EntityBlock   string
}

// Placeholder driver placeholder
func (d DaoTemplateData) Placeholder(num int) string {
	if d.DriverName == "postgres" {
		return "$" + strconv.Itoa(num)
	}
	return "?"
}

func (d DaoTemplateData) IsDriverPkg() bool {
	for _, c := range d.Table.Columns {
		if c.EnumData.TypeName != "" {
			return true
		}
	}
	return false
}

// TableTemplateData table data.
type TableTemplateData struct {
	Name      string
	TitleName string
	Columns   []ColumnTemplateData
}

// ColumnTemplateData column data.
type ColumnTemplateData struct {
	Name            string
	TitleName       string
	TypeName        string
	TypePkgName     string
	TypePointer     bool
	TypeDetail      string
	IsPrimaryKey    bool
	Sample          string
	IsAutoIncrement bool
	EnumData        EnumTemplateData // enumカラムのみ
}

func (c ColumnTemplateData) TypeFullName(pkgName string) string {
	typeName := c.TypeName
	if len(c.TypePkgName) > 0 && pkgName != c.TypePkgName {
		typeName = strings.Join([]string{c.TypePkgName, c.TypeName}, ".")
	}
	if c.TypePointer {
		typeName = "*" + typeName
	}
	return typeName
}

// CamelName convert go lint CamelCase
func (c ColumnTemplateData) CamelName() string {
	return lintName(stringUp.CamelCase(c.Name))
}

// AssetTemplateData asset data.
type AssetTemplateData struct {
	DaoPkgName string
}

// QueryArgsTemplateData query args data.
type QueryArgsTemplateData struct {
	DaoPkgName string
	SQLRootDir string
	DriverName string
}

// EnumData enumを構成する1要素
type EnumData struct {
	Name  string // HogeDataType
	Value string // HOGE
}

// EnumTemplateData enumのTypeとenumの要素
type EnumTemplateData struct {
	TypeName string     // DataType
	Enums    []EnumData // const () の中身
}

func (d AssetTemplateData) execAssetTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := AssetTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, "asset_gen.go", buf.Bytes())
}

func (d AssetTemplateData) execAssetFileTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := AssetFileTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, "asset_file_gen.go", buf.Bytes())
}

func (d QueryArgsTemplateData) execQueryArgsTemplate(daoRootDir string) error {
	var buf bytes.Buffer
	if err := QueryArgsTemplate(&buf, d); err != nil {
		return err
	}
	return formatFileWrite(daoRootDir, "queryargs_gen.go", buf.Bytes())
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
