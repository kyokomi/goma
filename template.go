package main

// DaoTemplateData dao data.
type DaoTemplateData struct {
	Name       string
	MemberName string
	EntityName string
	Table      TableTemplateData
	Imports    []string
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
	IsPrimaryKey bool
	Sample       string
}

// HelperTemplateData helper data.
type HelperTemplateData struct {
	PkgName      string
	DriverImport string
	Options      map[string]interface{}
}
