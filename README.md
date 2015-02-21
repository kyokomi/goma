goma (胡麻)
====================
[![GoDoc](https://godoc.org/github.com/kyokomi/goma/goma?status.svg)](https://godoc.org/github.com/kyokomi/goma/goma)

goma (胡麻) is a Database access framework for golang（Go）

I'm making based on [Doma](https://github.com/domaframework/doma);

*golang version 1.4.0 〜*

## Install

```
$ go get -u github.com/kyokomi/goma
```

## Usage

### Example `test` database `quest` table

```
mysql> SHOW COLUMNS FROM quest;
+-----------+----------+------+-----+---------+-------+
| Field     | Type     | Null | Key | Default | Extra |
+-----------+----------+------+-----+---------+-------+
| id        | int(11)  | NO   | PRI | 0       |       |
| name      | text     | YES  |     | NULL    |       |
| detail    | text     | YES  |     | NULL    |       |
| create_at | datetime | YES  |     | NULL    |       |
+-----------+----------+------+-----+---------+-------+
4 rows in set (0.00 sec)
```

### Example main.go（mysql）

```go
package main

import (
	"fmt"
)

//go:generate goma -driver=mysql -user=admin -password=password -host=localhost -port=3306 -db=test -debug=true

func main() {
	fmt.Println("Hello goma!")
}
```

driver and dataSource params.

- `driver="mysql"`: sql driver (`mysql` or `postgres`)
- `user="admin"`: database access user's name
- `password="password"`: database access user's password
- `host="localhost"`: database host
- `port=3306`: database port
- `db="test"`: database name

goma options params.

- `debug=false`: goma debug mode (`true`: log output)
- `dao="dao"`: generate dao root dirs
- `sql="sql"`: generate sql root dir
- `entity="entity"`: generate entity root dir

### Run

```
$ go generate
```

### Output

- `xxxxx1`,`xxxxx2`: TableName
 
```
├── dao
│   ├── xxxxx1_gen.go
│   └── xxxxx2_gen.go
├── entity
│   ├── xxxxx1_gen.go
│   └── xxxxx2_gen.go
├── gomautils_gen.go
├── main.go
└── sql
    ├── xxxxx1
    │   ├── delete.sql
    │   ├── insert.sql
    │   ├── selectAll.sql
    │   ├── selectByID.sql
    │   └── update.sql
    └── xxxxx2
        ├── delete.sql
        ├── insert.sql
        ├── selectAll.sql
        ├── selectByID.sql
        └── update.sql
```

[example code](https://github.com/kyokomi/goma/blob/master/example)

### Generate Code Usage

import your generate dao package.

```go
import "xxxxxxx/yyyyyy/zzzz/entity"
```

#### GomaOpen

```go
// ひらけゴマ!!
goma, err := NewGoma() // gomautils_gen.go
if err != nil {
    log.Fatalln(err)
}
defer g.Close()
```

#### Insert
 
```go
_, err = goma.Quest.Insert(entity.QuestEntity{
    ID:       99,
    Name:     "test",
    Detail:   "test detail",
    CreateAt: time.Now(),
})
if err != nil {
    log.Fatalln(err)
}
```

#### Select

```go
questEntity, err := goma.Quest.SelectByID(99); err != nil {
    log.Fatalln(err)
}

fmt.Printf("insert after: %+v\n", questEntity)
```

#### Update

```go
_, err = goma.Quest.Update(entity.QuestEntity{
    ID:       99,
    Name:     "test 2",
    Detail:   "test detail 2",
    CreateAt: time.Now(),
})
if err != nil {
    log.Fatalln(err)
}
```

#### Delete

```go
_, err = goma.Quest.Delete(99)
if err != nil {
    log.Fatalln(err)
}
```

### Add Custom Query Example

create new file [sql/updateAll.sql](https://github.com/kyokomi/goma/blob/master/example/sql/updateAll.sql).

```sql
update quest set
    name = /* name */'test'
,   detail = /* detail */'test'
,   create_at = /* create_at */'2006/01/02 13:40:00'
```

create new file [dao/quest_custom.go](https://github.com/kyokomi/goma/blob/master/example/dao/quest_custom.go).

```go
package dao

import (
	"database/sql"
	"log"

	"github.com/kyokomi/goma/example/entity"
	"github.com/kyokomi/goma/goma"
)

// Update update quest table.
func (d *QuestDao) UpdateAll(entity entity.QuestEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"name":      entity.Name,
		"detail":    entity.Detail,
		"create_at": entity.CreateAt,
	}
	queryString := d.QueryArgs("quest", "updateAll", args)

	result, err := d.Exec(queryString)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
```

#### Call

```go
_, err = goma.Quest.UpdateAll(entity.QuestEntity{
    Name:     "test 2",
    Detail:   "test detail 2",
    CreateAt: time.Now(),
})
if err != nil {
    log.Fatalln(err)
}
```

## Support Driver

- mysql（https://github.com/go-sql-driver/mysql）
- postgres（https://github.com/lib/pq）

## Author

[kyokomi](https://github.com/kyokomi)

## Licence

[Apache License Version 2.0](https://github.com/kyokomi/goma/blob/master/LICENSE)
