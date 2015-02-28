package main
import (
"fmt"
"io"
)
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
_, _ = fmt.Fprintf(w, "\"\n    \n    \"github.com/kyokomi/goma\"\n)\n\n// ")
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, " is generated ")
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:17
_, _ = fmt.Fprintf(w, " table.\ntype ")
//line dao_template.go.ego:18
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:18
_, _ = fmt.Fprintf(w, " struct {\n\t*goma.Goma\n\ttx *sql.Tx\n\tTableName string\n}\n\n// ")
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, " is ")
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, ".\nfunc ")
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "(g *goma.Goma) ")
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, " {\n    tblDao := ")
//line dao_template.go.ego:26
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:26
_, _ = fmt.Fprintf(w, "{}\n    tblDao.Goma = g\n    tblDao.tx = nil\n    tblDao.TableName = \"")
//line dao_template.go.ego:29
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:29
_, _ = fmt.Fprintf(w, "\"\n\treturn tblDao\n}\n\n// IsTx started transaction?\nfunc (d *")
//line dao_template.go.ego:34
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:34
_, _ = fmt.Fprintf(w, ") IsTx() bool {\n\treturn d.tx != nil\n}\n\n// SetTx set transaction.\nfunc (d *")
//line dao_template.go.ego:39
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:39
_, _ = fmt.Fprintf(w, ") SetTx(tx *sql.Tx) {\n\td.tx = tx\n}\n\n// ResetTx reset transaction.\n// Call after Commit of Rollback.\nfunc (d *")
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, ") ResetTx() {\n\td.tx = nil\n}\n\nfunc (d *")
//line dao_template.go.ego:49
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:49
_, _ = fmt.Fprintf(w, ") daoQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {\n\tif d.IsTx() {\n\t\trows, err = d.tx.Query(query, args...)\n\t} else {\n\t\trows, err = d.Query(query, args...)\n\t}\n\treturn\n}\n\nfunc (d *")
//line dao_template.go.ego:58
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:58
_, _ = fmt.Fprintf(w, ") daoExec(query string, args ...interface{}) (result sql.Result, err error) {\n\tif d.IsTx() {\n\t\tresult, err = d.tx.Exec(query, args...)\n\t} else {\n\t\tresult, err = d.Exec(query, args...)\n\t}\n\treturn\n}\n\n// SelectAll select ")
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:67
_, _ = fmt.Fprintf(w, " table all recode.\nfunc (d *")
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, ") SelectAll() ([]*")
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:68
_, _ = fmt.Fprintf(w, ", error) {\n    queryString := d.QueryArgs(\"")
//line dao_template.go.ego:69
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:69
_, _ = fmt.Fprintf(w, "\", \"selectAll\", nil)\n\n\tvar entitys []*")
//line dao_template.go.ego:71
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:71
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:71
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:72
_, _ = fmt.Fprintf(w, "\n\trows, err := d.daoQuery(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tfor rows.Next() {\n\t\tvar entity ")
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, "\n\t\terr = rows.Scan(")
//line dao_template.go.ego:79
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:79
 if idx != 0 { 
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:79
 } 
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, "&entity.")
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:79
 } 
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, ")\n\t\tif err != nil {\n\t\t\tbreak\n\t\t}\n\n\t\tentitys = append(entitys, &entity)\n\t}\n\tif err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\n\treturn entitys, nil\n}\n\n// SelectByID select ")
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d *")
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, ") SelectByID(")
//line dao_template.go.ego:95
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:95
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:95
 continue 
//line dao_template.go.ego:95
 } 
//line dao_template.go.ego:95
 if idx != 0 { 
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:95
 } 
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:95
 } 
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, ") (*")
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:95
_, _ = fmt.Fprintf(w, ", error) {\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:97
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:97
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:97
 continue 
//line dao_template.go.ego:97
 } 
//line dao_template.go.ego:97
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:97
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:97
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:97
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:97
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:98
 } 
//line dao_template.go.ego:98
_, _ = fmt.Fprintf(w, "}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:99
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:99
_, _ = fmt.Fprintf(w, "\", \"selectByID\", args)\n\n\trows, err := d.daoQuery(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer rows.Close()\n\n\tif !rows.Next() {\n\t\treturn nil, nil\n\t}\n\n\tvar entity ")
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:111
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:112
_, _ = fmt.Fprintf(w, "\n\tif err := rows.Scan(")
//line dao_template.go.ego:112
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:112
 if idx != 0 { 
//line dao_template.go.ego:112
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:112
 } 
//line dao_template.go.ego:112
_, _ = fmt.Fprintf(w, "&entity.")
//line dao_template.go.ego:112
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:112
 } 
//line dao_template.go.ego:112
_, _ = fmt.Fprintf(w, "); err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\t\n\treturn &entity, nil\n}\n\n// Insert insert ")
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, " table.\nfunc (d *")
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, ") Insert(entity ")
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:121
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\targs := goma.QueryArgs{\n    ")
//line dao_template.go.ego:123
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:124
 } 
//line dao_template.go.ego:125
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:126
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:126
_, _ = fmt.Fprintf(w, "\", \"insert\", args)\n\t\n\tresult, err := d.daoExec(queryString)\n    if err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Update update ")
//line dao_template.go.ego:135
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:135
_, _ = fmt.Fprintf(w, " table.\nfunc (d *")
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, ") Update(entity ")
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:136
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\targs := goma.QueryArgs{\n\t")
//line dao_template.go.ego:138
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:138
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:138
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:138
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:138
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:138
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:139
 } 
//line dao_template.go.ego:140
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:141
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:141
_, _ = fmt.Fprintf(w, "\", \"update\", args)\n\n\tresult, err := d.daoExec(queryString)\n\tif err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Delete delete ")
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:150
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d *")
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, ") Delete(")
//line dao_template.go.ego:151
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:151
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:151
 continue 
//line dao_template.go.ego:151
 } 
//line dao_template.go.ego:151
 if idx != 0 { 
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:151
 } 
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:151
 } 
//line dao_template.go.ego:151
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:153
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:153
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:153
 continue 
//line dao_template.go.ego:153
 } 
//line dao_template.go.ego:153
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:153
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:153
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:153
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:153
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:154
 } 
//line dao_template.go.ego:154
_, _ = fmt.Fprintf(w, "}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:155
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:155
_, _ = fmt.Fprintf(w, "\", \"delete\", args)\n\n    result, err := d.daoExec(queryString)\n\tif err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n")
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
_, _ = fmt.Fprintf(w, "\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\n")
//line entity_template.go.ego:7
 if len(daoData.Imports) > 0 { 
//line entity_template.go.ego:7
_, _ = fmt.Fprintf(w, "import (\n\t")
//line entity_template.go.ego:8
 for _, importName := range daoData.Imports { 
//line entity_template.go.ego:9
_, _ = fmt.Fprintf(w, "\n\t    \"")
//line entity_template.go.ego:9
_, _ = fmt.Fprintf(w, "%v",  importName )
//line entity_template.go.ego:9
_, _ = fmt.Fprintf(w, "\"\n    ")
//line entity_template.go.ego:10
 } 
//line entity_template.go.ego:11
_, _ = fmt.Fprintf(w, "\n)")
//line entity_template.go.ego:11
 } 
//line entity_template.go.ego:12
_, _ = fmt.Fprintf(w, "\n    \n// ")
//line entity_template.go.ego:13
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:13
_, _ = fmt.Fprintf(w, " is generated ")
//line entity_template.go.ego:13
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line entity_template.go.ego:13
_, _ = fmt.Fprintf(w, " table.\ntype ")
//line entity_template.go.ego:14
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line entity_template.go.ego:14
_, _ = fmt.Fprintf(w, " struct {\n")
//line entity_template.go.ego:15
 for _, column := range daoData.Table.Columns { 
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, " ")
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, " //")
//line entity_template.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  column.TypeDetail )
//line entity_template.go.ego:16
_, _ = fmt.Fprintf(w, "\n")
//line entity_template.go.ego:16
 } 
//line entity_template.go.ego:16
_, _ = fmt.Fprintf(w, "}\n")
return nil
}
//line gomautils_template.go.ego:1
func HelperTemplate(w io.Writer, helperData HelperTemplateData) error  {
//line gomautils_template.go.ego:1
_, _ = fmt.Fprintf(w, "package ")
//line gomautils_template.go.ego:1
_, _ = fmt.Fprintf(w, "%v",  helperData.PkgName )
//line gomautils_template.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n// NOTE: THIS FILE WAS PRODUCED BY THE\n// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)\n// DO NOT EDIT\n\nimport (\n    \"os\"\n\n    ")
//line gomautils_template.go.ego:10
_, _ = fmt.Fprintf(w, "%v",  helperData.DriverImport )
//line gomautils_template.go.ego:11
_, _ = fmt.Fprintf(w, "\n\t\n\t\"")
//line gomautils_template.go.ego:12
_, _ = fmt.Fprintf(w, "%v",  helperData.DaoImport )
//line gomautils_template.go.ego:12
_, _ = fmt.Fprintf(w, "\"\n        \n\t\"github.com/kyokomi/goma\"\n)\n\n// Goma goma.Goma utils.\ntype Goma struct {\n\t*goma.Goma\n}\n\n// NewGoma is goma.Goma wrapper utils.\nfunc NewGoma() (Goma, error) {\n\topts := goma.Options{\n\t    ")
//line gomautils_template.go.ego:25
 for _, t := range helperData.Options { 
//line gomautils_template.go.ego:25
 for key, value := range t { 
//line gomautils_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  key )
//line gomautils_template.go.ego:25
_, _ = fmt.Fprintf(w, ": ")
//line gomautils_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  value )
//line gomautils_template.go.ego:25
_, _ = fmt.Fprintf(w, ",")
//line gomautils_template.go.ego:25
 } 
//line gomautils_template.go.ego:26
_, _ = fmt.Fprintf(w, "\n    ")
//line gomautils_template.go.ego:26
 } 
//line gomautils_template.go.ego:26
_, _ = fmt.Fprintf(w, "}\n    \n\tcurrentDir, err := os.Getwd()\n\tif err != nil {\n        return Goma{}, err\n\t}\n\topts.CurrentDir = currentDir\n\n    g, err := goma.NewGoma(opts)\n    if err != nil {\n        return Goma{}, err\n    }\n\n    gm := Goma{}\n    gm.Goma = g\n\treturn gm, nil\n} \n\n")
//line gomautils_template.go.ego:44
 for _, daoData := range helperData.DaoList { 
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, "func (g Goma) ")
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, "() ")
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  helperData.DaoPkgName )
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, ".")
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line gomautils_template.go.ego:44
_, _ = fmt.Fprintf(w, " {\n    return ")
//line gomautils_template.go.ego:45
_, _ = fmt.Fprintf(w, "%v",  helperData.DaoPkgName )
//line gomautils_template.go.ego:45
_, _ = fmt.Fprintf(w, ".")
//line gomautils_template.go.ego:45
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line gomautils_template.go.ego:45
_, _ = fmt.Fprintf(w, "(g.Goma)\n}\n")
//line gomautils_template.go.ego:47
 } 
//line gomautils_template.go.ego:48
_, _ = fmt.Fprintf(w, "\n")
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
