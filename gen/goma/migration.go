package main

import (
	"io"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/migu"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func migration(opt goma.Options, modelsFilePath string) error {
	db, err := goma.OpenOptions(opt)
	if err != nil {
		return err
	}
	defer db.Close()

	var src io.Reader
	_, err = ioutil.ReadFile(modelsFilePath)
	if err != nil {
		// fileじゃないのでdir
		err = migu.SyncDir(db, modelsFilePath, src, true)
	} else {
		err = migu.Sync(db, modelsFilePath, src, true)
	}
	return err
}
