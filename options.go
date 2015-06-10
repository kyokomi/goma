package goma

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
)

// Options is open sql.DB options.
type Options struct {
	// driver and dataSource
	Driver   string `json:"driver"`   // DriverName
	UserName string `json:"user"`     // access user name `admin`
	PassWord string `json:"password"` // access user password `password`
	Host     string `json:"host"`     // localhost
	Port     int    `json:"port"`     // 3306
	DBName   string `json:"db"`       // DataBaseName
	Location string `json:"location"` // Location name

	// postgres
	SSLMode string `json:"ssl"` // disable, verify-full

	// goma
	IsConfig      bool   `json:"isConfig"`      // goma config generate
	Debug         bool   `json:"debug"`         // goma debug mode (default false)
	SQLRootDir    string `json:"sqlRootDir"`    // goma sql root dir path (default './sql')
	DaoRootDir    string `json:"daoRootDir"`    // goma dao root dir path (default './dao')
	EntityRootDir string `json:"entityRootDir"` // goma entity root dir path (default './entity')
	CurrentDir    string `json:"currentDir"`    // goma currentDir path
}

// SQLRootDirPath result sql dir path.
func (o Options) SQLRootDirPath() string {
	return filepath.Join(o.CurrentDir, o.SQLRootDir)
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

// ConfigPath result config path.
func (o Options) ConfigPath() string {
	return filepath.Join(o.CurrentDir, "config.json")
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
			o.Location,
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
		var value interface{}
		val := v.Field(i)
		if val.Kind() == reflect.String {
			value = fmt.Sprintf(`"%s"`, val.String())
		} else if val.Kind() == reflect.Bool {
			value = val.Bool()
		} else {
			value = val.Int()
		}

		name := t.Field(i).Name
		res = append(res, map[string]interface{}{name: value})
	}

	return res
}

// NewOptions create read file Options.
func NewOptions(filePath string) (Options, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Options{}, err
	}

	var opts Options
	if strings.HasSuffix(filePath, ".json") {
		if err := json.Unmarshal(f, &opts); err != nil {
			return opts, err
		}
	}

	return opts, nil
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
