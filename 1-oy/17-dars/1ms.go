package main

import "fmt"

type Branch struct {
	Id      int
	Name    string
	Address string
}

type Transaction struct {
	Id        int
	BranchId  int
	ProductId int
	Quantity  int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func main() {
	var product = []Product{
		{Id: 1, Name: "Olma", Price: 12000},
		{Id: 2, Name: "Banan", Price: 22000},
		{Id: 3, Name: "Olcha", Price: 25000},
	}
	var transaction = []Transaction{
		{Id: 1, BranchId: 1, ProductId: 1, Quantity: 12},
		{Id: 2, BranchId: 2, ProductId: 2, Quantity: 10},
		{Id: 3, BranchId: 1, ProductId: 3, Quantity: 8},
		{Id: 4, BranchId: 2, ProductId: 1, Quantity: 14},
		{Id: 5, BranchId: 1, ProductId: 2, Quantity: 13},
		{Id: 6, BranchId: 2, ProductId: 3, Quantity: 7},
	}
	// var branches = []Branch{
	// 	{Id: 1, Name: "MGorkiy", Address: "Mirzo Ulug'bek 17 uy"},
	// 	{Id: 2, Name: "Chilonzor", Address: "Chilonzor Metro"},
	// 	{Id: 3, Name: "Chilonzor", Address: "Chilonzor Metro"},
	// }
	// for _, item1 := range branches {
	// 	a := 0
	// 	for _, s := range product {
	// 		for _, item := range transaction {
	// 			if item.BranchId == item1.Id && s.Id == item.ProductId {
	// 				a += item.Quantity * s.Price
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(item1.Id,"-id total sum= ",a)
	// }

	// counterMap:=make(map[int]int)
	// productMap:=make(map[int]Product)
	// branchesMap:=make(map[int]Branch)

	// for _,v:=range product{
	// 	productMap[v.Id]=v
	// }
	// for _,v:=range branches{
	// 	branchesMap[v.Id]=v
	// }
	// for _,v:=range transaction{
	// 	counterMap[v.BranchId]+=productMap[v.ProductId].Price
	// }

	slice := [10]int{}
	var a, b int
	for _, item1 := range product {
		a = 0
		for _, item := range transaction {
			if item1.Id == item.ProductId {
				a += item.Quantity
				// fmt.Println(item1.Id,"quantity",a)
				b = item.Id
			}
		}
		slice[b] = a
	}
	son1 := slice[0]
	for _, son := range slice {
		if son > son1 {
			son1 = son
		}
	}
	fmt.Println(son1)
}
