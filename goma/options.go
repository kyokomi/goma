package goma

import "fmt"

// Options is open sql.DB options.
type Options struct {
	// driver and dataSource
	Driver   string // DriverName
	UserName string // access user name `admin`
	PassWord string // access user password `password`
	Host     string // localhost
	Port     int    // 3306
	DBName   string // DataBaseName

	// goma
	Debug      bool   // goma debug mode (default false)
	SQLRootDir string // goma sql root dir path (default './sql')
	DaoRootDir string // goma dao root dir path (default './dao')
}

// Source create driver databaseSource.
func (o Options) Source() string {

	switch o.Driver {
	case "mysql":
		// admin:password@tcp(localhost:3306)/test?parseTime=true&loc=Local
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v&loc=%s",
			o.UserName,
			o.PassWord,
			o.Host,
			o.Port,
			o.DBName,
			true,
			"Local",
		)
	case "postgres":
		return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
			o.UserName,
			o.PassWord,
			o.Host,
			o.Port,
			o.DBName,
		)
	}

	panic("not support driver name: " + o.Driver)
}

func (o Options) Tuples() []map[string]interface{} {
	
	// mapそのままだと順番がgenerateするごとに変わるの配列のmapにしてる
	
	// TODO: あとでなんとかする... reflect とか使えばなんとかなりそう
	t := make([]map[string]interface{}, 0)
	t = append(t, map[string]interface{}{"Driver": fmt.Sprintf(`"%s"`, o.Driver)})
	t = append(t, map[string]interface{}{"UserName": fmt.Sprintf(`"%s"`, o.UserName)})
	t = append(t, map[string]interface{}{"PassWord": fmt.Sprintf(`"%s"`, o.PassWord)})
	t = append(t, map[string]interface{}{"Host": fmt.Sprintf(`"%s"`, o.Host)})
	t = append(t, map[string]interface{}{"Port": o.Port})
	t = append(t, map[string]interface{}{"DBName": fmt.Sprintf(`"%s"`, o.DBName)})
	t = append(t, map[string]interface{}{"Debug": o.Debug})
	t = append(t, map[string]interface{}{"SQLRootDir": fmt.Sprintf(`"%s"`, o.SQLRootDir)})
	t = append(t, map[string]interface{}{"DaoRootDir": fmt.Sprintf(`"%s"`, o.DaoRootDir)})
	return t
}
