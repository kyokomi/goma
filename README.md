goma (胡麻)
====================

[![Circle CI](https://circleci.com/gh/kyokomi/goma.svg?style=svg)](https://circleci.com/gh/kyokomi/goma)
[![GoDoc](https://godoc.org/github.com/kyokomi/goma?status.svg)](https://godoc.org/github.com/kyokomi/goma)

goma (胡麻) is a Database access framework for golang（Go）

I'm making based on [Doma](https://github.com/domaframework/doma);

*golang version 1.4.0 〜*

## Install

```
$ go get gopkg.in/kyokomi/goma.v1/gen/goma
```

latest Ver

```
$ go get github.com.kyokomi/goma/gen/goma
```


## Usage

[GitBook Document](http://kyokomi.gitbooks.io/goma/content/)

### Command

```
   init-config	create example config file
   gen		generate code by params
   gen-config	generate code by config
   help, h	Shows a list of commands or help for one command
```

## Support Driver

- mysql ( https://github.com/go-sql-driver/mysql )
- postgres ( https://github.com/lib/pq )

## Author

[kyokomi](https://github.com/kyokomi)

## Licence

[Apache License Version 2.0](https://github.com/kyokomi/goma/blob/master/LICENSE)
