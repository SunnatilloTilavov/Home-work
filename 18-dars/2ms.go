package main  //omitepty--- bo'sh bolsa tashlab ketadi 

import (
	"encoding/json"
	"fmt"
)

// Foydalanuvchi struct ni JSON ma'lumotlariga o'xshash qilib tanlab olish uchun ishlatamiz
type Foydalanuvchi struct {
	Ism    string `json:"ism"`
	Yosh   int    `json:"yosh"`
	Manzil Manzil  `json:"manzil"`
}


// Manzil struct ni o'z ichiga oladi
type Manzil struct {
	Shahar  string `json:"shahar"`
	Viloyat string `json:"viloyat"`
}


func main() {
	// JSON ma'lumotni kodlash (marshalling)
	foydalanuvchi := Foydalanuvchi{
		Ism:  "John",
		Yosh: 30,
		Manzil: Manzil{
			Shahar:  "Toshkent",
			Viloyat: "Toshkent",
		},
	}

	jsonData, err := json.MarshalIndent(foydalanuvchi," ","  ")
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	// JSON ma'lumotni ekranga chiqarish
	fmt.Println("Kodlangan JSON ma'lumot:", string(jsonData))

	// JSON ma'lumotni o'qish (unmarshalling)
	var yangiFoydalanuvchi Foydalanuvchi
	err = json.Unmarshal(jsonData, &yangiFoydalanuvchi)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	// O'qilgan JSON ma'lumotni ekranga chiqarish
	fmt.Printf("O'qilgan JSON ma'lumot: %+v", yangiFoydalanuvchi)
}
