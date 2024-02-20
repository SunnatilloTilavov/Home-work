package main

import (
	"database/sql"
	"fmt"
	"homework/2-oy/4-dars/homework/country"
	"homework/2-oy/4-dars/homework/storage"

	_ "github.com/lib/pq"
)

func main() {
	var id, name string
	var code int
	db := connectDB()
	defer db.Close()
	fmt.Println("Choose : \n 1-Create row \n 2-Update row  \n 3-Delete row \n 4-Get all")
	userType := 0
	_, err := fmt.Scan(&userType)
	if err != nil {
		panic(err)
	}
	inv := storage.NewInventory(db)
	switch userType {
	case 1:
		var name string
		code := 0
		fmt.Println("insert country name")
		fmt.Scan(&name)
		fmt.Println("insert country code")
		fmt.Scan(&code)
		country := country.Country{
			Name: name,
			Code: code,
		}
		err := inv.Create(country)
		if err != nil {
			return
		}
		fmt.Println("Country created successfully")
	case 3:
		fmt.Println("Choose : \n 1-Soft delete \n 2-hard delete ")
		userType := 0
		_, err = fmt.Scan(&userType)
		if err != nil {
			panic(err)
		}
		inv := storage.NewInventory(db)

		countries, err := inv.GetAll()
		if err != nil {
			return
		}
		fmt.Println("Countries: ", countries)


		switch userType {
		case 1:

			var id string
			fmt.Println("IDni tanlang")
			fmt.Scan(&id)
			err = inv.Delete(id)
			if err != nil {
				return
			}
			fmt.Println("Country DELETE successfully")

		case 2:
		fmt.Println("IDni tanlang")
		fmt.Scan(&id)
		err = inv.Delete1(country.Country{},id)
		if err != nil {
			return
		}
		fmt.Println("Country delete successfully")

		}

	case 2:
		countries, err := inv.GetAll()
		if err != nil {
			return
		}
		fmt.Println("Countries: ", countries)

	
		fmt.Println("IDni tanlang")
		fmt.Scan(&id)
		fmt.Println("updete country name")
		fmt.Scan(&name)
		fmt.Println("updete country code")
		fmt.Scan(&code)
		err = inv.Update(country.Country{}, name, code, id)
		if err != nil {
			return
		}
		fmt.Println("Country update successfully")

	case 4:

		countries, err := inv.GetAll()
		if err != nil {
			return
		}
		fmt.Println("Countries: ", countries)
	}

}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=sunnatillo password=1111 database=sqldatabase sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
