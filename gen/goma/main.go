package main

import (
	"os"

	"github.com/codegangsta/cli"
)

//go:generate ego -package main templates

func main() {
	app := cli.NewApp()
	app.Name = "goma"
	app.Version = "2.0"
	app.Usage = ""
	app.Author = "kyokomi"
	app.Email = "kyoko1220adword@gmail.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "debug", Usage: "goma debug mode", EnvVar: ""},
		cli.StringFlag{Name: "file", Value: os.Getenv("GOFILE"), Usage: "input file", EnvVar: ""},
		cli.StringFlag{Name: "pkg", Value: os.Getenv("GOPACKAGE"), Usage: "output package", EnvVar: ""},
	}

	genFlags := []cli.Flag{
		cli.StringFlag{Name: "driver", Value: "mysql", Usage: "sql driver", EnvVar: ""},
		cli.StringFlag{Name: "user", Value: "admin", Usage: "database access user's name", EnvVar: ""},
		cli.StringFlag{Name: "password", Value: "", Usage: "database access user's password", EnvVar: ""},
		cli.StringFlag{Name: "host", Value: "localhost", Usage: "database host", EnvVar: ""},
		cli.IntFlag{Name: "port", Value: 3306, Usage: "database port", EnvVar: ""},
		cli.StringFlag{Name: "db", Value: "test", Usage: "database name", EnvVar: ""},
		cli.StringFlag{Name: "location", Value: "UTC", Usage: "database location name", EnvVar: ""},

		cli.StringFlag{Name: "ssl", Value: "disable", Usage: "postgres ssl mode", EnvVar: ""},

		cli.StringFlag{Name: "sql", Value: "sql", Usage: "generate sql root dir", EnvVar: ""},
		cli.StringFlag{Name: "dao", Value: "dao", Usage: "generate dao root dir", EnvVar: ""},
		cli.StringFlag{Name: "entity", Value: "entity", Usage: "generate entity root dir", EnvVar: ""},
		cli.BoolFlag{Name: "config", Usage: "generate config", EnvVar: ""},
	}

	miguGenFlags := make([]cli.Flag, len(genFlags))
	for idx, c := range genFlags {
		miguGenFlags[idx] = c
	}
	miguGenFlags = append(miguGenFlags,
		cli.StringFlag{Name: "models", Value: "models.go", Usage: "models struct file path", EnvVar: ""},
	)

	app.Commands = []cli.Command{
		{
			Name:   "init-config",
			Action: initConfigAction,
			Usage:  "create example config file",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "out", Value: "", Usage: "output dir path", EnvVar: ""},
			},
		},
		{
			Name:   "gen",
			Action: genAction,
			Usage:  "generate code by params",
			Flags:  genFlags,
		},
		{
			Name:   "gen-config",
			Action: genConfigAction,
			Usage:  "generate code by config",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "path", Value: "config.json", Usage: "config path", EnvVar: ""},
			},
		},
	}
	app.Run(os.Args)
}
