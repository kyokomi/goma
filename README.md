# goma
goma is a Database access framework for golang（Go）

I'm making based on [Doma](https://github.com/domaframework/doma);

## Usage

### Example main.go

```go
package main

import "fmt"

//go:generate goma

func main() {
	fmt.Println("hoge")
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

## TODO

- [ ] go generateで`mysql, admin:password@tcp(localhost:3306)/test`みたいなのをもらってEntityを生成する
    - xxxx_gen.goに書き込む（xxxxはテーブル名）
- [ ] DBにあるTable一覧で、sqlパッケージ下に一致するパッケージがあるか探す
    - みつからないものは新規生成（とりあえず、SelectByIDとSelectAll）
    - 見つかったら上書きとか考慮する
- [ ] NewGomaするときに全パッケージのsqlをcacheする（パッケージ、メソッド名、引数でKeyにする）
