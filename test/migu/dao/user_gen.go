package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"log"

	"database/sql"

	"github.com/kyokomi/goma/test/migu/entity"
)

// UserDaoQueryer is interface
type UserDaoQueryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// UserDao is generated user table.
type UserDao struct {
	*sql.DB
	TableName string
}

// Query UserDao query
func (g UserDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.DB.Query(query, args...)
}

// Exec UserDao exec
func (g UserDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.DB.Exec(query, args...)
}

var _ UserDaoQueryer = (*UserDao)(nil)

// User is UserDao.
func User(db *sql.DB) UserDao {
	tblDao := UserDao{}
	tblDao.DB = db
	tblDao.TableName = "User"
	return tblDao
}

// TxUserDao is generated user table transaction.
type TxUserDao struct {
	*sql.Tx
	TableName string
}

// Query TxUserDao query
func (g TxUserDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return g.Tx.Query(query, args...)
}

// Exec TxUserDao exec
func (g TxUserDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return g.Tx.Exec(query, args...)
}

var _ UserDaoQueryer = (*TxUserDao)(nil)

// TxUser is UserDao.
func TxUser(tx *sql.Tx) TxUserDao {
	tblDao := TxUserDao{}
	tblDao.Tx = tx
	tblDao.TableName = "User"
	return tblDao
}

// SelectAll select user table all recode.
func (g UserDao) SelectAll() ([]entity.User, error) {
	return _UserSelectAll(g)
}

// SelectAll transaction select user table all recode.
func (g TxUserDao) SelectAll() ([]entity.User, error) {
	return _UserSelectAll(g)
}

func _UserSelectAll(g UserDaoQueryer) ([]entity.User, error) {
	queryString := `
select
  id
, name
, email
, create_at
, update_at
, age
FROM
  user`

	var es []entity.User
	rows, err := g.Query(queryString)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	for rows.Next() {
		var e entity.User
		if err := e.Scan(rows); err != nil {
			break
		}

		es = append(es, e)
	}
	if err != nil {
		log.Println(err, queryString)
		return nil, err
	}

	return es, nil
}

// SelectByID select user table by primaryKey.
func (g UserDao) SelectByID(id int64) (entity.User, error) {
	return _UserSelectByID(g, id)
}

// SelectByID transaction select user table by primaryKey.
func (g TxUserDao) SelectByID(id int64) (entity.User, error) {
	return _UserSelectByID(g, id)
}

func _UserSelectByID(g UserDaoQueryer, id int64) (entity.User, error) {
	queryString := `
select
  id
, name
, email
, create_at
, update_at
, age
FROM
  user
WHERE
  id = ?
`
	rows, err := g.Query(queryString,
		id,
	)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.User{}, sql.ErrNoRows
	}

	var e entity.User
	if err := e.Scan(rows); err != nil {
		log.Println(err, queryString)
		return entity.User{}, err
	}

	return e, nil
}

// Insert insert user table.
func (g UserDao) Insert(entity entity.User) (sql.Result, error) {
	return _UserInsert(g, entity)
}

// Insert transaction insert user table.
func (g TxUserDao) Insert(entity entity.User) (sql.Result, error) {
	return _UserInsert(g, entity)
}

func _UserInsert(g UserDaoQueryer, entity entity.User) (sql.Result, error) {
	queryString := `
insert into user(
  id
, name
, email
, create_at
, update_at
, age
) values(
  ?
, ?
, ?
, ?
, ?
, ?
);`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.Name,
		entity.Email,
		entity.CreateAt,
		entity.UpdateAt,
		entity.Age,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Update update user table.
func (g UserDao) Update(entity entity.User) (sql.Result, error) {
	return _UserUpdate(g, entity)
}

// Update transaction update user table.
func (g TxUserDao) Update(entity entity.User) (sql.Result, error) {
	return _UserUpdate(g, entity)
}

// Update update user table.
func _UserUpdate(g UserDaoQueryer, entity entity.User) (sql.Result, error) {
	queryString := `
update user set
    id = ?
,   name = ?
,   email = ?
,   create_at = ?
,   update_at = ?
,   age = ?
 where
    id = ?

`
	result, err := g.Exec(queryString,
		entity.ID,
		entity.Name,
		entity.Email,
		entity.CreateAt,
		entity.UpdateAt,
		entity.Age,

		entity.ID,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}

// Delete delete user table.
func (g UserDao) Delete(id int64) (sql.Result, error) {
	return _UserDelete(g, id)
}

// Delete transaction delete user table.
func (g TxUserDao) Delete(id int64) (sql.Result, error) {
	return _UserDelete(g, id)
}

// Delete delete user table by primaryKey.
func _UserDelete(g UserDaoQueryer, id int64) (sql.Result, error) {
	queryString := `
delete
from
  user
where
  id = ?

`
	result, err := g.Exec(queryString,
		id,
	)
	if err != nil {
		log.Println(err, queryString)
	}
	return result, err
}
