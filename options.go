package goma

import (
	"fmt"
	"path/filepath"
	"strings"
	"reflect"
)

// Options is open sql.DB options.
type Options struct {
	// driver and dataSource
	Driver   string // DriverName
	UserName string // access user name `admin`
	PassWord string // access user password `password`
	Host     string // localhost
	Port     int    // 3306
	DBName   string // DataBaseName

	// postgres
	SSLMode string // disable, verify-full

	// goma
	Debug         bool   // goma debug mode (default false)
	SQLRootDir    string // goma sql root dir path (default './sql')
	DaoRootDir    string // goma dao root dir path (default './dao')
	EntityRootDir string // goma entity root dir path (default './entity')
	CurrentDir    string // goma currentDir path
}

// DaoImportPath result dao package import.
func (o Options) DaoImportPath() string {
	return importPath(o.CurrentDir, o.DaoPkgName())
}

// EntityImportPath result entity package import.
func (o Options) EntityImportPath() string {
	return importPath(o.CurrentDir, o.EntityPkgName())
}

// DaoPkgName result dao package name.
func (o Options) DaoPkgName() string {
	return packageName(o.DaoRootDir)
}

// EntityPkgName result entity package name.
func (o Options) EntityPkgName() string {
	return packageName(o.EntityRootDir)
}

func importPath(currentDir, pkgName string) string {
	// /User/xxxxx/src/github.com/yyyyy => github.com/yyyyy
	srcIdx := strings.Index(currentDir, "src/")
	// github.com/yyyyy + dao => github.com/yyyyy/dao
	return filepath.Join(currentDir[srcIdx+len("src/"):], pkgName)
}

func packageName(pkgRootDir string) string {
	// ./dao/ => dao
	pkgName := strings.Replace(pkgRootDir, "./", "", -1)
	return strings.Replace(pkgName, "/", "", -1)
}

// Source create driver databaseSource.
func (o Options) Source() string {
	var source string
	switch o.Driver {
	default:
		panic("not support driver name: " + o.Driver)
	case "mysql":
		// admin:password@tcp(localhost:3306)/test?parseTime=true&loc=Local
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v&loc=%s",
			o.UserName,
			o.PassWord,
			o.Host,
			o.Port,
			o.DBName,
			true,
			"Local",
		)
	case "postgres":
		source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
			o.UserName,
			o.PassWord,
			o.Host,
			o.Port,
			o.DBName,
			o.SSLMode,
		)
	}

	if o.Debug {
		fmt.Println(source)
	}

	return source
}

// Tuples opions to string array map.
func (o Options) Tuples() []map[string]interface{} {
	// mapそのままだと順番がgenerateするごとに変わるの配列のmapにしてる
	var res []map[string]interface{}

	fmt.Printf("%+v\n", o)

	v := reflect.ValueOf(o)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name
		if name == "CurrentDir" {
			continue
		}

		var value interface{}
		val := v.Field(i)
		if val.Kind() == reflect.String{
			value = fmt.Sprintf(`"%s"`, val.String())
		} else if val.Kind() == reflect.Bool{
			value = val.Bool()
		} else {
			value = val.Int()
		}
		res = append(res, map[string]interface{}{name: value})
	}

	return res
}
