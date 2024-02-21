package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/5-dars/models"
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

func (i *Inventory) Create(c models.Category) error {
	id := uuid.NewString()
	_, err := i.db.Exec(
		`INSERT INTO Category (id,name,created_at)
		VALUES($1,$2,CURRENT_TIMESTAMP)`, id, c.Name)
	if err != nil {
		fmt.Println("error while creating Category err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Update(c models.Category, name string, id string) error {
	_, err := i.db.Exec(
		`UPDATE Category SET 
		    name=$2,
			updated_at=NOW()
			WHERE id = $1`, id, name)
	if err != nil {
		fmt.Println("error while updating Category err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Delete1(c models.Category, id string) error {
	_, err := i.db.Exec(
		`UPDATE Category SET 
		   deleted_at=NOW()
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while delete1 Category err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Delete(id string) error {
	_, err := i.db.Exec(
		`DELETE FROM Category
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while deleting Category: ", err)
		return err
	}

	return nil
}

func (i *Inventory) GetById(id string) (models.Category, error) {
	var countries models.Category
	err := i.db.QueryRow(`
	SELECT
	id,
	name,
	created_at
	FROM Category
	WHERE id=$1`, id).Scan(&countries.Id, &countries.Name, &countries.CreatedAt)
	if err != nil {
		fmt.Println("error while get by id Category err:", err)
		return countries, err
	}
	return countries, nil

}

func (i *Inventory) GetAll() ([]models.Category, error) {
	Categorys := []models.Category{}
	rows, err := i.db.Query(`select 
	id,
	name,
	created_at from Category WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all Category err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := models.Category{}
		if err = rows.Scan(&c.Id, &c.Name, &c.CreatedAt); err != nil {
			fmt.Println("error while scanning Category err: ", err)
			return nil, err
		}
		Categorys = append(Categorys, c)
	}

	return Categorys, nil
}

func (i *Inventory) GetId() ([]models.Category, error) {
	Categorys := []models.Category{}
	rows, err := i.db.Query(`select id from Category WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all Category err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := models.Category{}
		if err = rows.Scan(&c.Id); err != nil {
			fmt.Println("error while scanning Category err: ", err)
			return nil, err
		}
		Categorys = append(Categorys, c)
	}

	return Categorys, nil
}
