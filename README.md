goma (胡麻)
====================

[![Circle CI](https://circleci.com/gh/kyokomi/goma.svg?style=svg)](https://circleci.com/gh/kyokomi/goma)
[![Coverage Status](https://coveralls.io/repos/kyokomi/goma/badge.svg?branch=master)](https://coveralls.io/r/kyokomi/goma?branch=master)
[![GoDoc](https://godoc.org/github.com/kyokomi/goma?status.svg)](https://godoc.org/github.com/kyokomi/goma)

goma (胡麻) is a Database access framework for golang（Go）

I'm making based on [Doma](https://github.com/domaframework/doma);

*golang version 1.4.0 〜*

## Install

```
$ go get -u github.com/kyokomi/goma/gen/goma
```

## Usage

[GitBook Document](http://kyokomi.gitbooks.io/goma/content/)

### Command

```
$ goma
COMMANDS:
   init-config create example config file
   gen         generate code by params
   config      generate code by config
   help, h	Shows a list of commands or help for one command
```

## Support Driver

- mysql ( https://github.com/go-sql-driver/mysql )
- postgres ( https://github.com/lib/pq )

## Author

[kyokomi](https://github.com/kyokomi)

## Licence

[Apache License Version 2.0](https://github.com/kyokomi/goma/blob/master/LICENSE)
