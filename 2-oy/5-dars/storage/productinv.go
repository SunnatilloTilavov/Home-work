package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/5-dars/models"
	"github.com/google/uuid"
)

type Productinv struct {
	db *sql.DB
}

func NewProductinv(db *sql.DB) Productinv {
	return Productinv{
		db: db,
	}
}

func (i *Productinv) Create1(c models.Product) error {
	id := uuid.NewString()
	_, err := i.db.Exec(
		`INSERT INTO Product (id,name,category_id,created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP)`, id, c.Name,c.Category_id)
	if err != nil {
		fmt.Println("error while creating Product err: ", err)
		return err
	}

	return nil
}

func (i *Productinv) Update(c models.Product, name string, id string) error {
	_, err := i.db.Exec(
		`UPDATE Product SET 
		    name=$2,
			updated_at=NOW()
			WHERE id = $1`, id, name)
	if err != nil {
		fmt.Println("error while updating Product err: ", err)
		return err
	}

	return nil
}

func (i *Productinv) Delete1(c models.Product, id string) error {
	_, err := i.db.Exec(
		`UPDATE Product SET 
		   deleted_at=NOW()
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while delete1 Product err: ", err)
		return err
	}

	return nil
}

func (i *Productinv) Delete(id string) error {
	_, err := i.db.Exec(
		`DELETE FROM Product
			WHERE id = $1`, id)
	if err != nil {
		fmt.Println("error while deleting Product: ", err)
		return err
	}

	return nil
}

func (i *Productinv) GetById(id string) (models.Product, error) {
	var countries models.Product
	err := i.db.QueryRow(`
	SELECT
	id,
	name,
	created_at
	FROM Product
	WHERE id=$1`, id).Scan(&countries.Id, &countries.Name, &countries.CreatedAt)
	if err != nil {
		fmt.Println("error while get by id Product err:", err)
		return countries, err
	}
	return countries, nil

}

func (i *Productinv) GetAll() ([]models.Product, error) {
	Products := []models.Product{}
	rows, err := i.db.Query(`select 
	id,
	name,
	created_at from Product WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all Product err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := models.Product{}
		if err = rows.Scan(&c.Id, &c.Name, &c.CreatedAt); err != nil {
			fmt.Println("error while scanning Product err: ", err)
			return nil, err
		}
		Products = append(Products, c)
	}

	return Products, nil
}

func (i *Productinv) GetId() ([]models.Product, error) {
	Products := []models.Product{}
	rows, err := i.db.Query(`select id from Product WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all Product err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := models.Product{}
		if err = rows.Scan(&c.Id); err != nil {
			fmt.Println("error while scanning Product err: ", err)
			return nil, err
		}
		Products = append(Products, c)
	}

	return Products, nil
}
2
