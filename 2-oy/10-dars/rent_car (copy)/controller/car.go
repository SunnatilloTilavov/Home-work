package controller

import (
	"fmt"
	"homework/2-oy/10-dars/rent_car/models"

	"time"
)

func (c *Controller) CreateCar() {
	car := getCarInfo()
	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year intput is not correct")
		return
	}
	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car created successfully with ID: %v\n", id)

}

func (c *Controller) Update(){
	car :=getCarInfoUpdate()
	id, err := c.Store.Car.Update(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car Update successfully with ID: %v\n", id)



}
// func (c *Controller) Delete(){
// 	car :=getCarInfoDelete()
// 	id, err := c.Store.Car.Delete(car)
// 	if err != nil {
// 		fmt.Println("error while creating car, err: ", err)
// 		return
// 	}
// 	fmt.Printf("Car Update successfully with ID: %v\n", id)



// }
// func getCarInfoDelete() models.Car {
// 	car := models.Car{}
// 	fmt.Println(`Update Enter the  car datas`)
// 	fmt.Println(` Name`)
// 	fmt.Scan(&car.Name)
// 	fmt.Println(` Year`)
// 	fmt.Scan(&car.Year)
// 	fmt.Println(` Brand`)
// 	fmt.Scan( &car.Brand)
// 	fmt.Println(` Model`)
// 	fmt.Scan(&car.Model)
// 	fmt.Println(` HoursePower`)
// 	fmt.Scan(&car.HoursePower)
// 	fmt.Println(` Colour`)
// 	fmt.Scan(&car.Colour)
// 	fmt.Println(` EngineCap`)
// 	fmt.Scan(&car.EngineCap)
// 	return car
// }


func getCarInfo() models.Car {
	car := models.Car{}
	fmt.Println(`Delete row`)
	fmt.Println("Id=")
	fmt.Scan(&car.Id)
	return car
}


func getCarInfoUpdate() models.Car {
	car := models.Car{}
	fmt.Println("Idni tanlang")
	fmt.Scan(&car.Id)
	fmt.Println(`Update Enter the  car datas`)
	fmt.Println(` Name`)
	fmt.Scan(&car.Name)
	fmt.Println(` Year`)
	fmt.Scan(&car.Year)
	fmt.Println(` Brand`)
	fmt.Scan( &car.Brand)
	fmt.Println(` Model`)
	fmt.Scan(&car.Model)
	fmt.Println(` HoursePower`)
	fmt.Scan(&car.HoursePower)
	fmt.Println(` Colour`)
	fmt.Scan(&car.Colour)
	fmt.Println(` EngineCap`)
	fmt.Scan(&car.EngineCap)

	return car
}

