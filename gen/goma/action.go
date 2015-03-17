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
	opt := goma.Options{}
	opt.Driver = c.String("driver")
	opt.UserName = c.String("user")
	opt.PassWord = c.String("password")
	opt.Host = c.String("host")
	opt.Port = c.Int("port")
	opt.DBName = c.String("db")
	opt.SSLMode = c.String("ssl")
	opt.SQLRootDir = c.String("sql")
	opt.DaoRootDir = c.String("dao")
	opt.EntityRootDir = c.String("entity")
	opt.Debug = c.GlobalBool("debug")

	generate(c.GlobalString("pkg"), opt)
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

func configAction(c *cli.Context) {
	// TODO: config.jsonを読み込んでgenerate呼ぶ
}
