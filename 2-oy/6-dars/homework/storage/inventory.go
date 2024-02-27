package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/6-dars/homework/country"
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
func (i *Inventory) Update(c country.Country, name string, code int, id string) error {
	_, err := i.db.Exec(
		`UPDATE countries SET 
		    name=$2,
			code=$3,
			updated_at=NOW()
			WHERE id = $1`, id, name, code)
	if err != nil {
		fmt.Println("error while updating country err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Delete1(c country.Country, id string) error {
	_, err := i.db.Exec(
		`UPDATE countries SET 
		   deleted_at=NOW()
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while delete1 country err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Delete(id string) error {
	_, err := i.db.Exec(
		`DELETE FROM countries
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while deleting country: ", err)
		return err
	}

	return nil
}

func (i *Inventory) GetById(id string) (country.Country, error) {
	var countries country.Country
	err := i.db.QueryRow(`
	SELECT
	id,
	name,
	code,
	created_at
	FROM countries
	WHERE id=$1`, id).Scan(&countries.Id, &countries.Name, &countries.Code, &countries.CreatedAt)
	if err != nil {
		fmt.Println("error while get by id countries err:", err)
		return countries, err
	}
	return countries, nil

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



func (i *Inventory) GetId() ([]country.Country, error) {
	countries := []country.Country{}
	rows, err := i.db.Query(`select id from countries WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all countries err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := country.Country{}
		if err = rows.Scan(&c.Id); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, nil
}

func (i *Inventory) GetSearch(search string) ([]country.Country, error) {
	countries := []country.Country{}
	rows, err := i.db.Query(`SELECT id, name, code, created_at FROM countries WHERE name LIKE $1 AND deleted_at IS NULL`, search+"%")
	if err != nil {
		fmt.Println("Error while getting countries:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		c := country.Country{}
		if err := rows.Scan(&c.Id, &c.Name, &c.Code, &c.CreatedAt); err != nil {
			fmt.Println("Error while scanning country:", err)
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, nil
}
