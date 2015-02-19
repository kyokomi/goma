package dao

// GENERATE CODE

import (
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
    Id int
    Name string
}

func (d *SampleDao) SelectAll() ([]*SampleEntity, error) {

    queryString := d.QueryArgs("sample", "selectAll", nil)

	var entitys []*SampleEntity
	rows, err := d.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity SampleEntity
		err = rows.Scan(&entity.Id, &entity.Name)
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
	if err := d.QueryRow(queryString).Scan(&entity.Id, &entity.Name); err != nil {
		return nil, err
	}
	
	return &entity, nil
}
