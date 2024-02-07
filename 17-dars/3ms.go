package main

import (
 "fmt"
)

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

var products = []Product{
 {Id: 1, Name: "Olma", Price: 12000},
 {Id: 2, Name: "Banan", Price: 22000},
 {Id: 3, Name: "Olcha", Price: 25000},
}

var transactions = []Transaction{
 {Id: 1, BranchId: 1, ProductId: 1, Quantity: 12},
 {Id: 2, BranchId: 2, ProductId: 2, Quantity: 10},
 {Id: 3, BranchId: 1, ProductId: 3, Quantity: 8},
 {Id: 4, BranchId: 2, ProductId: 1, Quantity: 14},
 {Id: 5, BranchId: 1, ProductId: 2, Quantity: 13},
 {Id: 6, BranchId: 2, ProductId: 3, Quantity: 7},
}

var branches = []Branch{
 {Id: 1, Name: "MGorkiy", Address: "Mirzo Ulug'bek 17 uy"},
 {Id: 2, Name: "Chilonzor", Address: "Chilonzor Metro"},
}

func main() {
counterMap:=make(map[int]int)
counterMap2:=make(map[int]int)
	productMap:=make(map[int]Product)
	branchesMap:=make(map[int]Branch)

	for _,v:=range products{
		productMap[v.Id]=v
	}
	for _,v:=range branches{
		branchesMap[v.Id]=v
	}
	for _,v:=range transactions{
		
		counterMap[v.ProductId]+=v.Quantity
		counterMap2[v.BranchId]+=v.Quantity
	}
	for i:=range branchesMap{
		fmt.Println("max min",counterMap[i])
		fmt.Println(branchesMap[i].Name,counterMap2[i])
	}

 
}