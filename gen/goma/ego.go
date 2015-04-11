package main
import (
"fmt"
"io"
)
//line asset_template.go.ego:1
func AssetTemplate(w io.Writer, assetData AssetTemplateData) error  {
//line asset_template.go.ego:1
_, _ = fmt.Fprintf(w, "package  ")
//line asset_template.go.ego:1
_, _ = fmt.Fprintf(w, "%v",  assetData.DaoPkgName )
//line asset_template.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\nimport (\n\t\"io/ioutil\"\n)\n\n// Asset default read file\nfunc Asset(filePath string) ([]byte, error) {\n\treturn ioutil.ReadFile(filePath)\n}\n")
return nil
}
//line dao_template.go.ego:1
func DaoTemplate(w io.Writer, daoData DaoTemplateData) error  {
//line dao_template.go.ego:1
_, _ = fmt.Fprintf(w, "package  ")
//line dao_template.go.ego:1
_, _ = fmt.Fprintf(w, "%v",  daoData.DaoPkgName )
//line dao_template.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\nimport (\n    \"log\"\n    \n    \"database/sql\"\n    \n    \"")
//line dao_template.go.ego:12
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityImport )
//line dao_template.go.ego:12
_, _ = fmt.Fprintf(w, "\"\n    \n    \"github.com/kyokomi/goma\"\n)\n\ntype ")
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, "Queryer interface {\n\tQuery(query string, args ...interface{}) (*sql.Rows, error)\n\tExec(query string, args ...interface{}) (sql.Result, error)\n}\n\n// ")
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, " is generated ")
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, " table.\ntype ")
//line dao_template.go.ego:23
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:23
_, _ = fmt.Fprintf(w, " struct {\n\t*sql.DB\n\tTableName string\n}\n\n// Query ")
//line dao_template.go.ego:28
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:28
_, _ = fmt.Fprintf(w, " query\nfunc (g ")
//line dao_template.go.ego:29
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:29
_, _ = fmt.Fprintf(w, ") Query(query string, args ...interface{}) (*sql.Rows, error) {\n\treturn g.DB.Query(query, args...)\n}\n\n// Exec ")
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, " exec\nfunc (g ")
//line dao_template.go.ego:34
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:34
_, _ = fmt.Fprintf(w, ") Exec(query string, args ...interface{}) (sql.Result, error) {\n\treturn g.DB.Exec(query, args...)\n}\n\nvar _ ")
//line dao_template.go.ego:38
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:38
_, _ = fmt.Fprintf(w, "Queryer = (*")
//line dao_template.go.ego:38
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:38
_, _ = fmt.Fprintf(w, ")(nil)\n\n// ")
//line dao_template.go.ego:40
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:40
_, _ = fmt.Fprintf(w, " is ")
//line dao_template.go.ego:40
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:40
_, _ = fmt.Fprintf(w, ".\nfunc ")
//line dao_template.go.ego:41
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:41
_, _ = fmt.Fprintf(w, "(db *sql.DB) ")
//line dao_template.go.ego:41
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:41
_, _ = fmt.Fprintf(w, " {\n    tblDao := ")
//line dao_template.go.ego:42
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:42
_, _ = fmt.Fprintf(w, "{}\n    tblDao.DB = db\n    tblDao.TableName = \"")
//line dao_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:44
_, _ = fmt.Fprintf(w, "\"\n\treturn tblDao\n}\n\n// ")
//line dao_template.go.ego:48
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:48
_, _ = fmt.Fprintf(w, "Tx is generated ")
//line dao_template.go.ego:48
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:48
_, _ = fmt.Fprintf(w, " table transaction.\ntype Tx")
//line dao_template.go.ego:49
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:49
_, _ = fmt.Fprintf(w, " struct {\n\t*sql.Tx\n\tTableName string\n}\n\n// Query Tx")
//line dao_template.go.ego:54
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:54
_, _ = fmt.Fprintf(w, " query\nfunc (g Tx")
//line dao_template.go.ego:55
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:55
_, _ = fmt.Fprintf(w, ") Query(query string, args ...interface{}) (*sql.Rows, error) {\n\treturn g.Tx.Query(query, args...)\n}\n\n// Exec Tx")
//line dao_template.go.ego:59
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:59
_, _ = fmt.Fprintf(w, " exec\nfunc (g Tx")
//line dao_template.go.ego:60
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:60
_, _ = fmt.Fprintf(w, ") Exec(query string, args ...interface{}) (sql.Result, error) {\n\treturn g.Tx.Exec(query, args...)\n}\n\nvar _ ")
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "Queryer = (*Tx")
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, ")(nil)\n\n// Tx")
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, " is ")
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, ".\nfunc Tx")
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, "(tx *sql.Tx) Tx")
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, " {\n    tblDao := Tx")
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, "{}\n    tblDao.Tx = tx\n    tblDao.TableName = \"")
//line dao_template.go.ego:70
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:70
_, _ = fmt.Fprintf(w, "\"\n\treturn tblDao\n}\n\n// SelectAll select ")
//line dao_template.go.ego:74
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:74
_, _ = fmt.Fprintf(w, " table all recode.\nfunc (d ")
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, ") SelectAll() ([]*")
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:75
_, _ = fmt.Fprintf(w, ", error) {\n\treturn ")
//line dao_template.go.ego:76
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:76
_, _ = fmt.Fprintf(w, "SelectAll(d)\n}\n\n// SelectAll transaction select ")
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, " table all recode.\nfunc (d Tx")
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, ") SelectAll() ([]*")
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:80
_, _ = fmt.Fprintf(w, ", error) {\n\treturn ")
//line dao_template.go.ego:81
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:81
_, _ = fmt.Fprintf(w, "SelectAll(d)\n}\n\nfunc ")
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "SelectAll(d ")
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "Queryer) ([]*")
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:84
_, _ = fmt.Fprintf(w, ", error) {\n    queryString := queryArgs(\"")
//line dao_template.go.ego:85
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:85
_, _ = fmt.Fprintf(w, "\", \"selectAll\", nil)\n\n\tvar es []*")
//line dao_template.go.ego:87
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:87
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:87
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, "\n\trows, err := d.Query(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tfor rows.Next() {\n\t\tvar e ")
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "\n\t\tif err := e.Scan(rows); err != nil {\n\t\t\tbreak\n\t\t}\n\n\t\tes = append(es, &e)\n\t}\n\tif err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\n\treturn es, nil\n}\n\n// SelectByID select ")
//line dao_template.go.ego:109
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:109
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d ")
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, ") SelectByID(")
//line dao_template.go.ego:110
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:110
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:110
 continue 
//line dao_template.go.ego:110
 } 
//line dao_template.go.ego:110
 if idx != 0 { 
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:110
 } 
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:110
 } 
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, ") (*")
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, ", error) {\n\treturn ")
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, "SelectByID(d, ")
//line dao_template.go.ego:111
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:111
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:111
 continue 
//line dao_template.go.ego:111
 } 
//line dao_template.go.ego:111
 if idx != 0 { 
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:111
 } 
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:111
 } 
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, ")\n}\n\n// SelectByID transaction select ")
//line dao_template.go.ego:114
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:114
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d Tx")
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, ") SelectByID(")
//line dao_template.go.ego:115
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:115
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:115
 continue 
//line dao_template.go.ego:115
 } 
//line dao_template.go.ego:115
 if idx != 0 { 
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:115
 } 
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:115
 } 
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, ") (*")
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:115
_, _ = fmt.Fprintf(w, ", error) {\n\treturn ")
//line dao_template.go.ego:116
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:116
_, _ = fmt.Fprintf(w, "SelectByID(d, ")
//line dao_template.go.ego:116
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:116
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:116
 continue 
//line dao_template.go.ego:116
 } 
//line dao_template.go.ego:116
 if idx != 0 { 
//line dao_template.go.ego:116
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:116
 } 
//line dao_template.go.ego:116
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:116
 } 
//line dao_template.go.ego:116
_, _ = fmt.Fprintf(w, ")\n}\n\nfunc ")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "SelectByID(d ")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "Queryer, ")
//line dao_template.go.ego:119
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:119
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:119
 continue 
//line dao_template.go.ego:119
 } 
//line dao_template.go.ego:119
 if idx != 0 { 
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:119
 } 
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:119
 } 
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, ") (*")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, ", error) {\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:121
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:121
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:121
 continue 
//line dao_template.go.ego:121
 } 
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:122
 } 
//line dao_template.go.ego:122
_, _ = fmt.Fprintf(w, "}\n\tqueryString := queryArgs(\"")
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "\", \"selectByID\", args)\n\n\trows, err := d.Query(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer rows.Close()\n\n\tif !rows.Next() {\n\t\treturn nil, nil\n\t}\n\n\tvar e ")
//line dao_template.go.ego:135
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:135
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:135
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, "\n\tif err := e.Scan(rows); err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\t\n\treturn &e, nil\n}\n\n// Insert insert ")
//line dao_template.go.ego:144
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:144
_, _ = fmt.Fprintf(w, " table.\nfunc (d ")
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, ") Insert(entity ")
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:145
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:146
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:146
_, _ = fmt.Fprintf(w, "Insert(d, entity)\n}\n\n// Insert transaction insert ")
//line dao_template.go.ego:149
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:149
_, _ = fmt.Fprintf(w, " table.\nfunc (d Tx")
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, ") Insert(entity ")
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, "Insert(d, entity)\n}\n\nfunc ")
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "Insert(d ")
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "Queryer, entity ")
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\targs := goma.QueryArgs{\n    ")
//line dao_template.go.ego:156
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:156
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:156
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:156
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:156
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:156
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:157
 } 
//line dao_template.go.ego:158
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := queryArgs(\"")
//line dao_template.go.ego:159
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:159
_, _ = fmt.Fprintf(w, "\", \"insert\", args)\n\t\n\tresult, err := d.Exec(queryString)\n    if err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Update update ")
//line dao_template.go.ego:168
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:168
_, _ = fmt.Fprintf(w, " table.\nfunc (d ")
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, ") Update(entity ")
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:169
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:170
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:170
_, _ = fmt.Fprintf(w, "Update(d, entity)\n}\n\n// Update transaction update ")
//line dao_template.go.ego:173
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:173
_, _ = fmt.Fprintf(w, " table.\nfunc (d Tx")
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, ") Update(entity ")
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:174
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:175
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:175
_, _ = fmt.Fprintf(w, "Update(d, entity)\n}\n\n// Update update ")
//line dao_template.go.ego:178
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:178
_, _ = fmt.Fprintf(w, " table.\nfunc ")
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "Update(d ")
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "Queryer, entity ")
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:179
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\targs := goma.QueryArgs{\n\t")
//line dao_template.go.ego:181
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:181
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:181
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:181
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:181
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:181
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:182
 } 
//line dao_template.go.ego:183
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := queryArgs(\"")
//line dao_template.go.ego:184
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:184
_, _ = fmt.Fprintf(w, "\", \"update\", args)\n\n\tresult, err := d.Exec(queryString)\n\tif err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Delete delete ")
//line dao_template.go.ego:193
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:193
_, _ = fmt.Fprintf(w, " table.\nfunc (d ")
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, ") Delete(")
//line dao_template.go.ego:194
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:194
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:194
 continue 
//line dao_template.go.ego:194
 } 
//line dao_template.go.ego:194
 if idx != 0 { 
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:194
 } 
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:194
 } 
//line dao_template.go.ego:194
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:195
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:195
_, _ = fmt.Fprintf(w, "Delete(d, ")
//line dao_template.go.ego:195
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:195
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:195
 continue 
//line dao_template.go.ego:195
 } 
//line dao_template.go.ego:195
 if idx != 0 { 
//line dao_template.go.ego:195
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:195
 } 
//line dao_template.go.ego:195
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:195
 } 
//line dao_template.go.ego:195
_, _ = fmt.Fprintf(w, ")\n}\n\n// Delete transaction delete ")
//line dao_template.go.ego:198
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:198
_, _ = fmt.Fprintf(w, " table.\nfunc (d Tx")
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, ") Delete(")
//line dao_template.go.ego:199
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:199
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:199
 continue 
//line dao_template.go.ego:199
 } 
//line dao_template.go.ego:199
 if idx != 0 { 
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:199
 } 
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:199
 } 
//line dao_template.go.ego:199
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\treturn ")
//line dao_template.go.ego:200
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:200
_, _ = fmt.Fprintf(w, "Delete(d, ")
//line dao_template.go.ego:200
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:200
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:200
 continue 
//line dao_template.go.ego:200
 } 
//line dao_template.go.ego:200
 if idx != 0 { 
//line dao_template.go.ego:200
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:200
 } 
//line dao_template.go.ego:200
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:200
 } 
//line dao_template.go.ego:200
_, _ = fmt.Fprintf(w, ")\n}\n\n// Delete delete ")
//line dao_template.go.ego:203
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:203
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc ")
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "Delete(d ")
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "Queryer, ")
//line dao_template.go.ego:204
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:204
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:204
 continue 
//line dao_template.go.ego:204
 } 
//line dao_template.go.ego:204
 if idx != 0 { 
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:204
 } 
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:204
 } 
//line dao_template.go.ego:204
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:206
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:206
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:206
 continue 
//line dao_template.go.ego:206
 } 
//line dao_template.go.ego:206
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:206
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:206
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:206
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:206
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:207
 } 
//line dao_template.go.ego:207
_, _ = fmt.Fprintf(w, "}\n\tqueryString := queryArgs(\"")
//line dao_template.go.ego:208
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:208
_, _ = fmt.Fprintf(w, "\", \"delete\", args)\n\n    result, err := d.Exec(queryString)\n\tif err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n")
return nil
}
//line delete_template.sql.ego:1
func DeleteTemplate(w io.Writer, tableData TableTemplateData) error  {
//line delete_template.sql.ego:1
_, _ = fmt.Fprintf(w, "delete\nfrom\n  ")
//line delete_template.sql.ego:3
_, _ = fmt.Fprintf(w, "%v",  tableData.Name )
//line delete_template.sql.ego:4
_, _ = fmt.Fprintf(w, "\nwhere\n")
//line delete_template.sql.ego:5
 for idx, column := range tableData.Columns { 
//line delete_template.sql.ego:5
 if !column.IsPrimaryKey { 
//line delete_template.sql.ego:5
 continue 
//line delete_template.sql.ego:5
 } 
//line delete_template.sql.ego:5
 if idx != 0 { 
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, "and")
//line delete_template.sql.ego:5
 } 
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, "  ")
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, " = /* ")
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line delete_template.sql.ego:5
_, _ = fmt.Fprintf(w, " */1\n")
//line delete_template.sql.ego:6
 } 
//line delete_template.sql.ego:7
_, _ = fmt.Fprintf(w, "\n")
return nil
}
//line entity_template.go.ego:1
func EntityTemplate(w io.Writer, daoData DaoTemplateData) error  {
//line entity_template.go.ego:1
_, _ = fmt.Fprintf(w, "package ")
//line entity_template.go.ego:1
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line entity_template.go.ego:2
_, _ = fmt.Fprintf(w, "\n\nimport \"database/sql\"\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\n")
//line entity_template.go.ego:9
 if len(daoData.Imports) > 0 { 
//line entity_template.go.ego:9
_, _ = fmt.Fprintf(w, "import (\n\t")
//line entity_template.go.ego:10
 for _, importName := range daoData.Imports { 
//line entity_template.go.ego:11
_, _ = fmt.Fprintf(w, "\n\t    \"")
//line entity_template.go.ego:11
_, _ = fmt.Fprintf(w, "%v",  importName )
//line entity_template.go.ego:11
_, _ = fmt.Fprintf(w, "\"\n    ")
//line entity_template.go.ego:12
 } 
//line entity_template.go.ego:13
_, _ = fmt.Fprintf(w, "\n)")
//line entity_template.go.ego:13
 } 
//line entity_template.go.ego:14
_, _ = fmt.Fprintf(w, "\n    \n// ")
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, " is generated ")
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, " table.\ntype ")
//line entity_template.go.ego:16
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:16
_, _ = fmt.Fprintf(w, " struct {\n")
//line entity_template.go.ego:17
 for _, column := range daoData.Table.Columns { 
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, " ")
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, " //")
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  column.TypeDetail )
//line entity_template.go.ego:18
_, _ = fmt.Fprintf(w, "\n")
//line entity_template.go.ego:18
 } 
//line entity_template.go.ego:18
_, _ = fmt.Fprintf(w, "}\n\n// Scan ")
//line entity_template.go.ego:20
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:20
_, _ = fmt.Fprintf(w, " all scan\nfunc (e *")
//line entity_template.go.ego:21
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:21
_, _ = fmt.Fprintf(w, ") Scan(rows *sql.Rows) error {\n\treturn rows.Scan(")
//line entity_template.go.ego:22
 for idx, column := range daoData.Table.Columns { 
//line entity_template.go.ego:22
 if idx != 0 { 
//line entity_template.go.ego:22
_, _ = fmt.Fprintf(w, ", ")
//line entity_template.go.ego:22
 } 
//line entity_template.go.ego:22
_, _ = fmt.Fprintf(w, "&e.")
//line entity_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line entity_template.go.ego:22
 } 
//line entity_template.go.ego:22
_, _ = fmt.Fprintf(w, ")\n}\n")
return nil
}
//line insert_template.sql.ego:1
func InsertTemplate(w io.Writer, tableData TableTemplateData) error  {
//line insert_template.sql.ego:1
_, _ = fmt.Fprintf(w, "insert into ")
//line insert_template.sql.ego:1
_, _ = fmt.Fprintf(w, "%v",  tableData.Name )
//line insert_template.sql.ego:1
_, _ = fmt.Fprintf(w, "(\n  ")
//line insert_template.sql.ego:2
 for idx, column := range tableData.Columns { 
//line insert_template.sql.ego:2
 if idx != 0 { 
//line insert_template.sql.ego:2
_, _ = fmt.Fprintf(w, ", ")
//line insert_template.sql.ego:2
 } 
//line insert_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line insert_template.sql.ego:3
_, _ = fmt.Fprintf(w, "\n")
//line insert_template.sql.ego:3
 } 
//line insert_template.sql.ego:3
_, _ = fmt.Fprintf(w, ") values(\n  ")
//line insert_template.sql.ego:4
 for idx, column := range tableData.Columns { 
//line insert_template.sql.ego:4
 if idx != 0 { 
//line insert_template.sql.ego:4
_, _ = fmt.Fprintf(w, ", ")
//line insert_template.sql.ego:4
 } 
//line insert_template.sql.ego:4
_, _ = fmt.Fprintf(w, "/* ")
//line insert_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line insert_template.sql.ego:4
_, _ = fmt.Fprintf(w, " */")
//line insert_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  column.Sample )
//line insert_template.sql.ego:5
_, _ = fmt.Fprintf(w, "\n")
//line insert_template.sql.ego:5
 } 
//line insert_template.sql.ego:5
_, _ = fmt.Fprintf(w, ");\n")
return nil
}
//line queryargs_template.go.ego:1
func QueryArgsTemplate(w io.Writer, queryArgsData QueryArgsTemplateData) error  {
//line queryargs_template.go.ego:1
_, _ = fmt.Fprintf(w, "package  ")
//line queryargs_template.go.ego:1
_, _ = fmt.Fprintf(w, "%v",  queryArgsData.DaoPkgName )
//line queryargs_template.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\nimport (\n\t\"path/filepath\"\n\n\t\"github.com/kyokomi/goma\"\n)\n\ntype queryArgSettings struct {\n\trootDir string\n}\n\nvar settings queryArgSettings\n\nfunc (s *queryArgSettings) SetRootDir(rootDir string) {\n\ts.rootDir = rootDir\n}\n\nfunc queryArgs(tableName string, queryName string, args goma.QueryArgs) string {\n\treturn settings.queryArgs(tableName, queryName, args)\n}\n\nfunc (s queryArgSettings) queryArgs(tableName string, queryName string, args goma.QueryArgs) string {\n\tfilePath := createSqlFilePath(s.rootDir, tableName, queryName)\n\treturn goma.GenerateQuery(assetSQL(filePath), args)\n}\n\nfunc assetSQL(filePath string) string {\n\tdata, err := Asset(filePath)\n\tif err != nil {\n\t\t// Asset was not found.\n\t}\n\treturn string(data)\n}\n\nfunc createSqlFilePath(rootDir string, tableName string, queryName string) string {\n\treturn filepath.Join(rootDir, \"")
//line queryargs_template.go.ego:41
_, _ = fmt.Fprintf(w, "%v",  queryArgsData.SQLRootDir )
//line queryargs_template.go.ego:41
_, _ = fmt.Fprintf(w, "\", tableName, queryName+\".sql\")\n}\n")
return nil
}
//line selectAll_template.sql.ego:1
func SelectAllTemplate(w io.Writer, tableData TableTemplateData) error  {
//line selectAll_template.sql.ego:1
_, _ = fmt.Fprintf(w, "select\n  ")
//line selectAll_template.sql.ego:2
 for idx, column := range tableData.Columns { 
//line selectAll_template.sql.ego:2
 if idx != 0 { 
//line selectAll_template.sql.ego:2
_, _ = fmt.Fprintf(w, ", ")
//line selectAll_template.sql.ego:2
 } 
//line selectAll_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line selectAll_template.sql.ego:3
_, _ = fmt.Fprintf(w, "\n")
//line selectAll_template.sql.ego:3
 } 
//line selectAll_template.sql.ego:3
_, _ = fmt.Fprintf(w, "FROM\n  ")
//line selectAll_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  tableData.Name )
//line selectAll_template.sql.ego:5
_, _ = fmt.Fprintf(w, "\n")
return nil
}
//line selectByID_template.sql.ego:1
func SelectByIDTemplate(w io.Writer, tableData TableTemplateData) error  {
//line selectByID_template.sql.ego:1
_, _ = fmt.Fprintf(w, "select\n  ")
//line selectByID_template.sql.ego:2
 for idx, column := range tableData.Columns { 
//line selectByID_template.sql.ego:2
 if idx != 0 { 
//line selectByID_template.sql.ego:2
_, _ = fmt.Fprintf(w, ", ")
//line selectByID_template.sql.ego:2
 } 
//line selectByID_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line selectByID_template.sql.ego:3
_, _ = fmt.Fprintf(w, "\n")
//line selectByID_template.sql.ego:3
 } 
//line selectByID_template.sql.ego:3
_, _ = fmt.Fprintf(w, "FROM\n  ")
//line selectByID_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  tableData.Name )
//line selectByID_template.sql.ego:5
_, _ = fmt.Fprintf(w, "\nWHERE\n")
//line selectByID_template.sql.ego:6
 for idx, column := range tableData.Columns { 
//line selectByID_template.sql.ego:6
 if !column.IsPrimaryKey { 
//line selectByID_template.sql.ego:6
 continue 
//line selectByID_template.sql.ego:6
 } 
//line selectByID_template.sql.ego:6
 if idx != 0 { 
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, "and")
//line selectByID_template.sql.ego:6
 } 
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, "  ")
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, " = /* ")
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, " */")
//line selectByID_template.sql.ego:6
_, _ = fmt.Fprintf(w, "%v",  column.Sample )
//line selectByID_template.sql.ego:7
_, _ = fmt.Fprintf(w, "\n")
//line selectByID_template.sql.ego:7
 } 
//line selectByID_template.sql.ego:8
_, _ = fmt.Fprintf(w, "\n")
return nil
}
//line update_template.sql.ego:1
func UpdateTemplate(w io.Writer, tableData TableTemplateData) error  {
//line update_template.sql.ego:1
_, _ = fmt.Fprintf(w, "update ")
//line update_template.sql.ego:1
_, _ = fmt.Fprintf(w, "%v",  tableData.Name )
//line update_template.sql.ego:1
_, _ = fmt.Fprintf(w, " set\n  ")
//line update_template.sql.ego:2
 for idx, column := range tableData.Columns { 
//line update_template.sql.ego:2
 if idx != 0 { 
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, ", ")
//line update_template.sql.ego:2
 } 
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, "  ")
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, " = /* ")
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, " */")
//line update_template.sql.ego:2
_, _ = fmt.Fprintf(w, "%v",  column.Sample )
//line update_template.sql.ego:3
_, _ = fmt.Fprintf(w, "\n")
//line update_template.sql.ego:3
 } 
//line update_template.sql.ego:3
_, _ = fmt.Fprintf(w, " where\n  ")
//line update_template.sql.ego:4
 for idx, column := range tableData.Columns { 
//line update_template.sql.ego:4
 if !column.IsPrimaryKey { 
//line update_template.sql.ego:4
 continue 
//line update_template.sql.ego:4
 } 
//line update_template.sql.ego:4
 if idx != 0 { 
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, "and")
//line update_template.sql.ego:4
 } 
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, "  ")
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, " = /* ")
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line update_template.sql.ego:4
_, _ = fmt.Fprintf(w, " */1\n")
//line update_template.sql.ego:5
 } 
//line update_template.sql.ego:6
_, _ = fmt.Fprintf(w, "\n")
return nil
}
