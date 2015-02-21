package goma

import "fmt"

// Options is open sql.DB options.
type Options struct {
	Driver   string // DriverName
	UserName string // access user name `admin`
	PassWord string // access user password `password`
	Host     string // localhost
	Port     int    // 3306
	DBName   string // DataBaseName
	Debug    bool
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

func (o Options) Map() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["Driver"] = fmt.Sprintf(`"%s"`, o.Driver)
	m["UserName"] = fmt.Sprintf(`"%s"`, o.UserName)
	m["PassWord"] = fmt.Sprintf(`"%s"`, o.PassWord)
	m["Host"] = fmt.Sprintf(`"%s"`, o.Host)
	m["Port"] = o.Port
	m["DBName"] = fmt.Sprintf(`"%s"`, o.DBName)
	return m
}
