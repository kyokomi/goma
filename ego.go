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
_, _ = fmt.Fprintf(w, "\"\n    \n    \"github.com/kyokomi/goma/goma\"\n)\n\n// ")
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
_, _ = fmt.Fprintf(w, " struct {\n\t*goma.Goma\n}\n\nvar ")
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.MemberName )
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, " *")
//line dao_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:23
_, _ = fmt.Fprintf(w, "\n\n// ")
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, " is ")
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:24
_, _ = fmt.Fprintf(w, " singleton.\nfunc ")
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "(g *goma.Goma) *")
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:25
_, _ = fmt.Fprintf(w, " {\n\tif ")
//line dao_template.go.ego:26
_, _ = fmt.Fprintf(w, "%v",  daoData.MemberName )
//line dao_template.go.ego:26
_, _ = fmt.Fprintf(w, " == nil {\n\t\t")
//line dao_template.go.ego:27
_, _ = fmt.Fprintf(w, "%v",  daoData.MemberName )
//line dao_template.go.ego:27
_, _ = fmt.Fprintf(w, " = &")
//line dao_template.go.ego:27
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:27
_, _ = fmt.Fprintf(w, "{Goma: g}\n\t}\n\treturn ")
//line dao_template.go.ego:29
_, _ = fmt.Fprintf(w, "%v",  daoData.MemberName )
//line dao_template.go.ego:30
_, _ = fmt.Fprintf(w, "\n}\n\n// SelectAll select ")
//line dao_template.go.ego:32
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:32
_, _ = fmt.Fprintf(w, " table all recode.\nfunc (d *")
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, ") SelectAll() ([]*")
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:33
_, _ = fmt.Fprintf(w, ", error) {\n\n    queryString := d.QueryArgs(\"")
//line dao_template.go.ego:35
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:35
_, _ = fmt.Fprintf(w, "\", \"selectAll\", nil)\n\n\tvar entitys []*")
//line dao_template.go.ego:37
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:37
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:37
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:38
_, _ = fmt.Fprintf(w, "\n\trows, err := d.Query(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tfor rows.Next() {\n\t\tvar entity ")
//line dao_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:44
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:44
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, "\n\t\terr = rows.Scan(")
//line dao_template.go.ego:45
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:45
 if idx != 0 { 
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:45
 } 
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, "&entity.")
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:45
 } 
//line dao_template.go.ego:45
_, _ = fmt.Fprintf(w, ")\n\t\tif err != nil {\n\t\t\tbreak\n\t\t}\n\n\t\tentitys = append(entitys, &entity)\n\t}\n\tif err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\n\treturn entitys, nil\n}\n\n// SelectByID select ")
//line dao_template.go.ego:60
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:60
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d *")
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, ") SelectByID(")
//line dao_template.go.ego:61
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:61
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:61
 continue 
//line dao_template.go.ego:61
 } 
//line dao_template.go.ego:61
 if idx != 0 { 
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:61
 } 
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:61
 } 
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, ") (*")
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:61
_, _ = fmt.Fprintf(w, ", error) {\n\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:64
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:64
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:64
 continue 
//line dao_template.go.ego:64
 } 
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:64
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:65
 } 
//line dao_template.go.ego:65
_, _ = fmt.Fprintf(w, "}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:66
_, _ = fmt.Fprintf(w, "\", \"selectByID\", args)\n\n\trows, err := d.Query(queryString)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer rows.Close()\n\n\tif !rows.Next() {\n\t\treturn nil, nil\n\t}\n\n\tvar entity ")
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:78
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:79
_, _ = fmt.Fprintf(w, "\n\tif err := d.QueryRow(queryString).Scan(")
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
_, _ = fmt.Fprintf(w, "); err != nil {\n\t    log.Println(err, queryString)\n\t\treturn nil, err\n\t}\n\t\n\treturn &entity, nil\n}\n\n// Insert insert ")
//line dao_template.go.ego:87
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:87
_, _ = fmt.Fprintf(w, " table.\nfunc (d *")
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, ") Insert(entity ")
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:88
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\n\targs := goma.QueryArgs{\n    ")
//line dao_template.go.ego:91
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:91
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:91
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:91
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:91
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:91
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:92
 } 
//line dao_template.go.ego:93
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:94
_, _ = fmt.Fprintf(w, "\", \"insert\", args)\n\t\n\tresult, err := d.Exec(queryString)\n    if err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Update update ")
//line dao_template.go.ego:103
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:103
_, _ = fmt.Fprintf(w, " table.\nfunc (d *")
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, ") Update(entity ")
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityPkgName )
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, ".")
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, "%v",  daoData.EntityName )
//line dao_template.go.ego:104
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\n\targs := goma.QueryArgs{\n\t")
//line dao_template.go.ego:107
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:107
_, _ = fmt.Fprintf(w, "  \"")
//line dao_template.go.ego:107
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:107
_, _ = fmt.Fprintf(w, "\": entity.")
//line dao_template.go.ego:107
_, _ = fmt.Fprintf(w, "%v",  column.TitleName )
//line dao_template.go.ego:107
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:108
 } 
//line dao_template.go.ego:109
_, _ = fmt.Fprintf(w, "\n\t}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:110
_, _ = fmt.Fprintf(w, "\", \"update\", args)\n\n\tresult, err := d.Exec(queryString)\n\tif err != nil {\n        log.Println(err, queryString)\n    }\n    return result, err\n}\n\n// Delete delete ")
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:119
_, _ = fmt.Fprintf(w, " table by primaryKey.\nfunc (d *")
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, ") Delete(")
//line dao_template.go.ego:120
 for idx, column := range daoData.Table.Columns { 
//line dao_template.go.ego:120
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:120
 continue 
//line dao_template.go.ego:120
 } 
//line dao_template.go.ego:120
 if idx != 0 { 
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, ", ")
//line dao_template.go.ego:120
 } 
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, " ")
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, "%v",  column.TypeName )
//line dao_template.go.ego:120
 } 
//line dao_template.go.ego:120
_, _ = fmt.Fprintf(w, ") (sql.Result, error) {\n\n    args := goma.QueryArgs{\n    ")
//line dao_template.go.ego:123
 for _, column := range daoData.Table.Columns { 
//line dao_template.go.ego:123
 if !column.IsPrimaryKey { 
//line dao_template.go.ego:123
 continue 
//line dao_template.go.ego:123
 } 
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "    \"")
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "\": ")
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, "%v",  column.Name )
//line dao_template.go.ego:123
_, _ = fmt.Fprintf(w, ",\n    ")
//line dao_template.go.ego:124
 } 
//line dao_template.go.ego:124
_, _ = fmt.Fprintf(w, "}\n\tqueryString := d.QueryArgs(\"")
//line dao_template.go.ego:125
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.Name )
//line dao_template.go.ego:125
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
//line entity_template.go.ego:16
_, _ = fmt.Fprintf(w, "\n")
//line entity_template.go.ego:16
 } 
//line entity_template.go.ego:17
_, _ = fmt.Fprintf(w, "\n}\n")
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
_, _ = fmt.Fprintf(w, "\"\n        \n\t\"github.com/kyokomi/goma/goma\"\n)\n\n// Goma goma.Goma utils.\ntype Goma struct {\n\t*goma.Goma\n\n\t// dao\n\t")
//line gomautils_template.go.ego:22
 for _, daoData := range helperData.DaoList { 
//line gomautils_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line gomautils_template.go.ego:22
_, _ = fmt.Fprintf(w, " *")
//line gomautils_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  helperData.DaoPkgName )
//line gomautils_template.go.ego:22
_, _ = fmt.Fprintf(w, ".")
//line gomautils_template.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  daoData.Name )
//line gomautils_template.go.ego:23
_, _ = fmt.Fprintf(w, "\n    ")
//line gomautils_template.go.ego:23
 } 
//line gomautils_template.go.ego:24
_, _ = fmt.Fprintf(w, "\n}\n\n// NewGoma is goma.Goma wrapper utils.\nfunc NewGoma() (Goma, error) {\n\n\topts := goma.Options{\n\t    ")
//line gomautils_template.go.ego:30
 for _, t := range helperData.Options { 
//line gomautils_template.go.ego:30
 for key, value := range t { 
//line gomautils_template.go.ego:30
_, _ = fmt.Fprintf(w, "%v",  key )
//line gomautils_template.go.ego:30
_, _ = fmt.Fprintf(w, ": ")
//line gomautils_template.go.ego:30
_, _ = fmt.Fprintf(w, "%v",  value )
//line gomautils_template.go.ego:30
_, _ = fmt.Fprintf(w, ",")
//line gomautils_template.go.ego:30
 } 
//line gomautils_template.go.ego:31
_, _ = fmt.Fprintf(w, "\n    ")
//line gomautils_template.go.ego:31
 } 
//line gomautils_template.go.ego:31
_, _ = fmt.Fprintf(w, "}\n    \n\tcurrentDir, err := os.Getwd()\n\tif err != nil {\n        return Goma{}, err\n\t}\n\topts.CurrentDir = currentDir\n\n    g, err := goma.NewGoma(opts)\n    if err != nil {\n        return Goma{}, err\n    }\n\n    gm := Goma{}\n    gm.Goma = g\n    ")
//line gomautils_template.go.ego:46
 for _, daoData := range helperData.DaoList { 
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, "gm.")
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, " = ")
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, "%v",  helperData.DaoPkgName )
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, ".")
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, "%v",  daoData.Table.TitleName )
//line gomautils_template.go.ego:46
_, _ = fmt.Fprintf(w, "(g)\n    ")
//line gomautils_template.go.ego:47
 } 
//line gomautils_template.go.ego:48
_, _ = fmt.Fprintf(w, "\n\treturn gm, nil\n} \n")
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
