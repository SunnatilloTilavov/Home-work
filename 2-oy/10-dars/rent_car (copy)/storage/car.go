package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/10-dars/rent_car/models"
	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCar(db *sql.DB) carRepo {
	return carRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *carRepo) Create(car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7) 
	`

	res, err := c.db.Exec(query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", res)

	return id.String(), nil
}


func (c *carRepo) Update(car models.Car) (string, error) {

	id := uuid.New()

	query := ` UPDATE cars SET 
		name=$2,
		brand=$3,
		model=$4,
		hourse_power=$5,
		colour=$6,
		engine_cap=$7,
		updated_at=NOW()	
	where id=$1`

	res, err := c.db.Exec(query,
		car.Id,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	fmt.Printf("Update db %+v\n", res)

	return id.String(), nil
}



func (i *carRepo) Delete(car models.Car) error {
	_, err := i.db.Exec(
		`DELETE FROM cars
			WHERE id = $1`,car.Id)
	if err != nil {
		fmt.Println("Error while deleting cars: ", err)
		return err
	}

	return nil
}

func (i *carRepo) GetAll() ([]models.Car, error) {
	cars := []models.Car{}
	rows, err := i.db.Query(`select 
	id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap,
		created_at from cars WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all cars err: ", err)
		return nil, err
	}

	for rows.Next() {
		car := models.Car{}
		if err = rows.Scan(car.Id,
			car.Name, car.Brand,
			car.Model, car.HoursePower,
			car.Colour, car.EngineCap,car.CreatedAt); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}
