goma
====================
[![GoDoc](https://godoc.org/github.com/kyokomi/goma/goma?status.svg)](https://godoc.org/github.com/kyokomi/goma/goma)

goma is a Database access framework for golang（Go）

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
	"log"

	"github.com/kyokomi/goma/example/dao"
)

//go:generate goma -driver=mysql -user=admin -password=password -host=localhost -port=3306 -db=test -debug=true

func main() {
	fmt.Println("Hello goma!")

    g, err := Goma()
    if err != nil {
        log.Fatalln(err)
    }
    defer g.Close()

	// Insert

	_, err = dao.Quest(g).Insert(dao.QuestEntity{
		ID:       99,
		Name:     "test",
		Detail:   "test detail",
		CreateAt: time.Now(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("insert after: %+v\n", q)
	}
	
	// Update

	_, err = dao.Quest(g).Update(dao.QuestEntity{
		ID:       99,
		Name:     "test 2",
		Detail:   "test detail 2",
		CreateAt: time.Now(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("update after: %+v\n", q)
	}

	// Delete

	_, err = dao.Quest(g).Delete(99)
	if err != nil {
		log.Fatalln(err)
	}

	if q, err := dao.Quest(g).SelectByID(99); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("delete after: %+v\n", q)
	}
}

```

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
├── helper_gen.go
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

## Support Driver

- mysql（https://github.com/go-sql-driver/mysql）
- postgres（https://github.com/lib/pq）

## Author

[kyokomi](https://github.com/kyokomi)

## Licence

[Apache License Version 2.0](https://github.com/kyokomi/goma/blob/master/LICENSE)
