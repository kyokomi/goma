package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kyokomi/goma"
)

func genAction(c *cli.Context) {
	generate(c.GlobalString("pkg"), scanGenFlags(c))
}

func initConfigAction(c *cli.Context) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	opt := goma.Options{}
	opt.CurrentDir = currentDir
	data, err := json.MarshalIndent(opt, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(opt.ConfigPath(), data, 0644); err != nil {
		log.Fatalln(err)
	}
}

func genConfigAction(c *cli.Context) {
	opt, err := goma.NewOptions(c.String("path"))
	if err != nil {
		log.Fatalln(err)
	}
	generate(c.GlobalString("pkg"), opt)
}

func scanGenFlags(c *cli.Context) goma.Options {
	opt := goma.Options{}
	opt.Driver = c.String("driver")
	opt.UserName = c.String("user")
	opt.PassWord = c.String("password")
	opt.Host = c.String("host")
	opt.Port = c.Int("port")
	opt.DBName = c.String("db")
	opt.Location = c.String("location")
	opt.SSLMode = c.String("ssl")
	opt.SQLRootDir = c.String("sql")
	opt.DaoRootDir = c.String("dao")
	opt.EntityRootDir = c.String("entity")
	opt.IsConfig = c.Bool("config")
	opt.Debug = c.GlobalBool("debug")
	return opt
}
