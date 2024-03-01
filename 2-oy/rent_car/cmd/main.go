package main

import (
	"fmt"
	"homework/2-oy/10-dars/rent_car/config"
	"homework/2-oy/10-dars/rent_car/controller"
	"homework/2-oy/10-dars/rent_car/storage"

	_ "golang.org/x/text/cases"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	c := controller.NewController(store)

	fmt.Println("Choose : \n 1-Create row \n 2-Update row  \n 3-Delete row \n 4-Get all ")
	userType := 0
	_, err = fmt.Scan(&userType)
	if err != nil {
		panic(err)
	}
	switch userType {
	case 1:
	c.CreateCar()
	case 2:
	c.Update()
	case 3:
	// c.Delete()
	case 4:
		cars, err := store.Car.GetAll()
		if err != nil {
			return
		}
		fmt.Println("Cars: ", cars)

}


}

