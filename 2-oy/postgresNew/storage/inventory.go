package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/4-dars/postgresNew/country"
	"github.com/google/uuid"
)

type Inventory struct {
	db *sql.DB
}

func NewInventory(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}

func (i *Inventory) Create(c country.Country) error {
	id := uuid.NewString()
	_, err := i.db.Exec(
		`INSERT INTO countries (id,name,code,created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP)`, id, c.Name, c.Code)
	if err != nil {
		fmt.Println("error while creating country err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Update(c country.Country) error {
	return nil
}

func (i *Inventory) GetAll() ([]country.Country, error) {
	countries := []country.Country{}
	rows, err := i.db.Query(`select 
	id,
	name,
	code,
	created_at from countries WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all countries err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := country.Country{}
		if err = rows.Scan(&c.Id, &c.Name, &c.Code, &c.CreatedAt); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, nil
}
