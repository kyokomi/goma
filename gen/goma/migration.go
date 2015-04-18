package main

import (
	"io"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/migu"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func migration(opt goma.Options, modelsFilePath string) error {
	db, err := goma.OpenOptions(opt)
	if err != nil {
		return err
	}

	var src io.Reader
	if err := migu.Sync(db, modelsFilePath, src, true); err != nil {
		return err
	}

	return nil
}
