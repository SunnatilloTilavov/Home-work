package main

import (
	"fmt"
	"homework/2-oy/13-dars/rent_car/config"
	"homework/2-oy/13-dars/rent_car/controller"
	"homework/2-oy/13-dars/rent_car/storage"
	"net/http"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/car", con.Car)
	http.HandleFunc("/customer", con.Customer)

	fmt.Println("programm is running on localhost:8080...")
	http.ListenAndServe(":8080", nil)

}
