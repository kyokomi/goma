# goma
goma is a Database access framework for golang（Go）

I'm making based on [Doma](https://github.com/domaframework/doma);

## Usage

### Example `test` database `quest` table

```
mysql> SHOW COLUMNS FROM quest;
+--------+---------+------+-----+---------+-------+
| Field  | Type    | Null | Key | Default | Extra |
+--------+---------+------+-----+---------+-------+
| id     | int(11) | NO   | PRI | 0       |       |
| name   | text    | YES  |     | NULL    |       |
| detail | text    | YES  |     | NULL    |       |
+--------+---------+------+-----+---------+-------+
3 rows in set (0.00 sec)
```

### Example main.go

```go
package main

import (
	"fmt"
	"log"

	"github.com/kyokomi/goma/example/dao"
	"github.com/kyokomi/goma/goma"
)

//go:generate goma -driver=mysql -source=admin:password@tcp(localhost:3306)/test

func main() {
	fmt.Println("Hello goma!")

	opts := goma.Options{
		Driver: "mysql",
		Source: "admin:password@tcp(localhost:3306)/test",
		Debug:  false,
	}
	g, err := goma.NewGoma(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	q, err := dao.Quest(g).SelectByID(1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", q)
}

```

### Run

```
$ go generate
```

### Output

- `xxxxx`: TableName
 
```
.
├── dao
│   └── xxxxx_gen.go
├── main.go
└── sql
```

[example code](https://github.com/kyokomi/goma/blob/master/example)

## Author

[kyokomi](https://github.com/kyokomi)

## Licence

[Apache License Version 2.0](https://github.com/kyokomi/goma/blob/master/LICENSE)
