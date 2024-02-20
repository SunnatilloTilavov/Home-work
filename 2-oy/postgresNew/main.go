package main

import (
	"database/sql"
	"fmt"
	"homework/2-oy/4-dars/postgresNew/country"
	"homework/2-oy/4-dars/postgresNew/storage"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()
	defer db.Close()

	inv := storage.NewInventory(db)
	country := country.Country{
		Name: "Uzbekistan",
		Code: 998,
	}
	err := inv.Create(country)
	if err != nil {
		return
	}
	fmt.Println("Country created successfully")

		countries, err := inv.GetAll()
		if err != nil {
			return
		}
		fmt.Println("Countries: ", countries)

}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=sunnatillo password=1111 database=sqldatabase sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
