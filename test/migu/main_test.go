package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"testing"

	"github.com/kyokomi/goma"
	"github.com/kyokomi/goma/test/migu/dao"
	"github.com/kyokomi/goma/test/migu/models"
)

const testID = int64(1234567896)

func TestSelectByID(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	_, err = dao.Quest(db).SelectByID(testID)
	if err != sql.ErrNoRows {
		t.Errorf("ERROR: rows")
	}
}

func TestInsert(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	q := models.Quest{}
	q.ID = testID
	q.Title = "test"
	q.Detail = "detail test"
	q.UpdateAt = time.Now()
	q.CreateAt = time.Now()
	if _, err := dao.Quest(db).Insert(q); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	_, err = dao.Quest(db).SelectByID(testID)
	if err == sql.ErrNoRows {
		t.Errorf("ERROR: %s", err)
	}
}

func TestSelectAll(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	list, err := dao.Quest(db).SelectAll()
	if len(list) != 1 {
		t.Errorf("ERROR: len != %d", len(list))
	}

	if list[0].ID != testID {
		t.Errorf("ERROR: %ld != %ld", list[0].ID, testID)
	}
}

func TestUpdate(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	q, err := dao.Quest(db).SelectByID(testID)
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}

	q.Title = "hogehoge"
	if _, err := dao.Quest(db).Update(q); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	q2, err := dao.Quest(db).SelectByID(testID)
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	if q.Title != q2.Title {
		t.Errorf("ERROR: update miss %s != %s", q.Title, q2.Title)
	}
}

func TestDelete(t *testing.T) {
	db, err := goma.Open("config.json")
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
	defer db.Close()

	if _, err := dao.Quest(db).Delete(testID); err != nil {
		t.Errorf("ERROR: %s", err)
	}

	_, err = dao.Quest(db).SelectByID(testID)
	if err != sql.ErrNoRows {
		t.Errorf("ERROR: %s", err)
	}
}
