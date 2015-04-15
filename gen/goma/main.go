package main

import (
	"os"

	"github.com/codegangsta/cli"
)

//go:generate ego -package main templates

func main() {
	app := cli.NewApp()
	app.Name = "goma"
	app.Version = "1.1"
	app.Usage = ""
	app.Author = "kyokomi"
	app.Email = "kyoko1220adword@gmail.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"debug", "goma debug mode", ""},
		cli.StringFlag{"file", os.Getenv("GOFILE"), "input file", ""},
		cli.StringFlag{"pkg", os.Getenv("GOPACKAGE"), "output package", ""},
	}
	app.Commands = []cli.Command{
		{
			Name:   "init-config",
			Action: initConfigAction,
			Usage:  "create example config file",
			Flags: []cli.Flag{
				cli.StringFlag{"out", "", "output dir path", ""},
			},
		},
		{
			Name:   "gen",
			Action: genAction,
			Usage:  "generate code by params",
			Flags: []cli.Flag{
				cli.StringFlag{"driver", "mysql", "sql driver", ""},
				cli.StringFlag{"user", "admin", "database access user's name", ""},
				cli.StringFlag{"password", "", "database access user's password", ""},
				cli.StringFlag{"host", "localhost", "database host", ""},
				cli.IntFlag{"port", 3306, "database port", ""},
				cli.StringFlag{"db", "test", "database name", ""},

				cli.StringFlag{"ssl", "disable", "postgres ssl mode", ""},

				cli.StringFlag{"sql", "sql", "generate sql root dir", ""},
				cli.StringFlag{"dao", "dao", "generate dao root dir", ""},
				cli.StringFlag{"entity", "entity", "generate entity root dir", ""},
			},
		},
		{
			Name:   "gen-simple",
			Action: genSimpleAction,
			Usage:  "generate code by params",
			Flags: []cli.Flag{
				cli.StringFlag{"driver", "mysql", "sql driver", ""},
				cli.StringFlag{"user", "admin", "database access user's name", ""},
				cli.StringFlag{"password", "", "database access user's password", ""},
				cli.StringFlag{"host", "localhost", "database host", ""},
				cli.IntFlag{"port", 3306, "database port", ""},
				cli.StringFlag{"db", "test", "database name", ""},

				cli.StringFlag{"ssl", "disable", "postgres ssl mode", ""},

				cli.StringFlag{"sql", "sql", "generate sql root dir", ""},
				cli.StringFlag{"dao", "dao", "generate dao root dir", ""},
				cli.StringFlag{"entity", "entity", "generate entity root dir", ""},
			},
		},
		{
			Name:   "gen-config",
			Action: genConfigAction,
			Usage:  "generate code by config",
			Flags: []cli.Flag{
				cli.StringFlag{"path", "config.json", "config path", ""},
			},
		},
	}
	app.Run(os.Args)
}
