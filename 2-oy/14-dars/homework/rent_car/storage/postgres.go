package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/13-dars/rent_car/config"

	_ "github.com/lib/pq"
)

type Store struct {
	DB  *sql.DB
	Car carRepo
	Customer customerRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

newCar := NewCar(db)
newCustomer := NewCustomer(db)

return Store{
    DB:       db,
    Car:      newCar,
    Customer: newCustomer,
}, nil
 

}
