package main

import (
	"database/sql"
	"fmt"
	"homework/2-oy/5-dars/models"
	"homework/2-oy/5-dars/storage"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	var id, name string
	db := connectDB()
	defer db.Close()
	fmt.Println("Choose :\n 1-Category \n 2-Product   \n 3-Get all")
	userType := 0
	_, err := fmt.Scan(&userType)
	if err != nil {
		panic(err)
	}

	switch userType {
	//case product
	case 1:
		inv := storage.NewInventory(db)
		fmt.Println("Choose : \n 1-Create row \n 2-Update row  \n 3-Delete row \n 4-Get all \n 5-GetBYid")
		userType := 0
		_, err := fmt.Scan(&userType)
		if err != nil {
			panic(err)
		}
		switch userType {
		case 1:

			var name string
			fmt.Println("insert Category name")
			fmt.Scan(&name)
			models := models.Category{
				Name: name,
			}
			err := inv.Create(models)
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

			Products, err := inv.GetId()
			if err != nil {
				return
			}
			fmt.Println("Countries: ", Products)

			switch userType {
			case 2:

				var id string
				fmt.Println("IDni tanlang")
				fmt.Scan(&id)
				err = inv.Delete(id)
				if err != nil {
					return
				}
				fmt.Println("Country DELETE successfully")

			case 1:
				fmt.Println("IDni tanlang")
				fmt.Scan(&id)
				err = inv.Delete1(models.Category{}, id)
				if err != nil {
					return
				}
				fmt.Println("Country delete successfully")

			}

		case 2:
			Categorys, err := inv.GetId()
			if err != nil {
				return
			}
			fmt.Println("Countries: ", Categorys)

			fmt.Println("IDni tanlang")
			fmt.Scan(&id)
			fmt.Println("updete country name")
			fmt.Scan(&name)
			err = inv.Update(models.Category{}, name, id)
			if err != nil {
				return
			}
			fmt.Println("Country update successfully")

		case 4:

			Category, err := inv.GetAll()
			if err != nil {
				return
			}
			fmt.Println("Category: ", Category)

		case 5:
			Category, err := inv.GetId()
			if err != nil {
				return
			}
			fmt.Println("Products: ", Category)
			fmt.Println("IDni tanlang")
			fmt.Scan(&id)
			onecountry, err := inv.GetById(id)
			if err != nil {
				return
			}
			fmt.Println(onecountry)

		}
///////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
		// case category
	case 2:
		inv2 := storage.NewProductinv(db)
		fmt.Println("Choose : \n 1-Create row \n 2-Update row  \n 3-Delete row \n 4-Get all \n 5-GetBYid")
		userType := 0
		_, err := fmt.Scan(&userType)
		if err != nil {
			panic(err)
		}
		switch userType {
		case 1:

			var name, category_idd string
			var category_id uuid.UUID
			var code int
			fmt.Println("insert Product name")
			fmt.Scan(&name)
			fmt.Println("insert Product code")
			fmt.Scan(&code)

			inv1 := storage.NewProductinv(db)
			Category, err := inv1.GetAll()
			if err != nil {
				return
			}
			fmt.Println("Category: ", Category)

			fmt.Println("choose category id")
			fmt.Scan(&category_idd)
			category_id, err = uuid.Parse(category_idd)
			if err != nil {
				fmt.Println("error parsing UUID:", err)
				return
			}
			models := models.Product{
				Name:        name,
				Code:        code,
				Category_id: category_id,
			}
			err = inv2.Create1(models)
			if err != nil {
				return
			}
			fmt.Println("Product created successfully")
		case 3:
			fmt.Println("Choose : \n 1-Soft delete \n 2-hard delete ")
			userType := 0
			_, err = fmt.Scan(&userType)
			if err != nil {
				panic(err)
			}
			inv2 := storage.NewProductinv(db)

			Products, err := inv2.GetId()
			if err != nil {
				return
			}
			fmt.Println("Product: ", Products)

			switch userType {
			case 2:

				var id string
				fmt.Println("IDni tanlang")
				fmt.Scan(&id)
				err = inv2.Delete(id)
				if err != nil {
					return
				}
				fmt.Println("Product DELETE successfully")

			case 1:
				fmt.Println("IDni tanlang")
				fmt.Scan(&id)
				err = inv2.Delete1(models.Product{}, id)
				if err != nil {
					return
				}
				fmt.Println("Product delete successfully")

			}

		case 2:
			Products, err := inv2.GetId()
			if err != nil {
				return
			}
			fmt.Println("Product: ", Products)

			fmt.Println("IDni tanlang")
			fmt.Scan(&id)
			fmt.Println("updete product name")
			fmt.Scan(&name)
			err = inv2.Update(models.Product{}, name, id)
			if err != nil {
				return
			}
			fmt.Println("Product update successfully")

		case 4:

			Category, err := inv2.GetAll()
			if err != nil {
				return
			}
			fmt.Println("Product: ", Category)

		case 5:
			Category, err := inv2.GetId()
			if err != nil {
				return
			}
			fmt.Println("Product: ", Category)
			fmt.Println("IDni tanlang")
			fmt.Scan(&id)
			onecountry, err := inv2.GetById(id)
			if err != nil {
				return
			}
			fmt.Println(onecountry)

		}
	case 3:
		inv := storage.NewInventory(db)
		Category1, err := inv.GetAll()
			if err != nil {
				return
			}
			fmt.Println("Category: ", Category1)


		inv1 := storage.NewProductinv(db)
		Category, err := inv1.GetAll()
			if err != nil {
				return
			}
			fmt.Println("Product: ", Category)

	}

}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=sunnatillo password=1111 database=models sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
