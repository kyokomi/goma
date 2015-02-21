package dao

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"database/sql"

	"github.com/kyokomi/goma/goma"
)

// SampleDao is generated sample table.
type SampleDao struct {
	*goma.Goma
}

var sample *SampleDao

// Sample is SampleDao singleton.
func Sample(g *goma.Goma) *SampleDao {
	if sample == nil {
		sample = &SampleDao{Goma: g}
	}
	return sample
}

// SampleEntity is generated sample table.
type SampleEntity struct {
	ID   int
	Name string
}

// SelectAll select sample table all recode.
func (d *SampleDao) SelectAll() ([]*SampleEntity, error) {

	queryString := d.QueryArgs("sample", "selectAll", nil)

	var entitys []*SampleEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity SampleEntity
		err = rows.Scan(&entity.ID, &entity.Name)
		if err != nil {
			break
		}

		entitys = append(entitys, &entity)
	}
	if err != nil {
		return nil, err
	}

	return entitys, nil
}

// SelectByID select sample table by primaryKey.
func (d *SampleDao) SelectByID(id int) (*SampleEntity, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("sample", "selectByID", args)

	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var entity SampleEntity
	if err := d.QueryRow(queryString).Scan(&entity.ID, &entity.Name); err != nil {
		return nil, err
	}

	return &entity, nil
}

// Insert insert sample table.
func (d *SampleDao) Insert(entity SampleEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":   entity.ID,
		"name": entity.Name,
	}
	queryString := d.QueryArgs("sample", "insert", args)

	return d.Exec(queryString)
}

// Update update sample table.
func (d *SampleDao) Update(entity SampleEntity) (sql.Result, error) {

	args := goma.QueryArgs{
		"id":   entity.ID,
		"name": entity.Name,
	}
	queryString := d.QueryArgs("sample", "update", args)

	return d.Exec(queryString)
}

// Delete delete sample table by primaryKey.
func (d *SampleDao) Delete(id int) (sql.Result, error) {

	args := goma.QueryArgs{
		"id": id,
	}
	queryString := d.QueryArgs("sample", "delete", args)

	return d.Exec(queryString)
}
