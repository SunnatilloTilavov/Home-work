package main

import "fmt"

type Client struct {
	Id     int
	Name   string
	Basket []BasketProduct
}

type BasketProduct struct {
	ProductId int
	Quantity  int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func (c *Client) AddToBasket(product Product, quantity int) {
	for i, ifor := range c.Basket {
		if ifor.ProductId == product.Id {
			c.Basket[i].Quantity += quantity
			return
		}
	}
	c.Basket = append(c.Basket, BasketProduct{ProductId: product.Id, Quantity: quantity})
			
}

func (c *Client) Total(product Product,quantity int)int {
	total:=0
	for _, ifor := range c.Basket {
		if ifor.ProductId == product.Id {
			total+=quantity*product.Price
			return total
		}	
	}			
}

func main() {
	client := Client{
		Id:   1,
		Name: "client1",
	

	}

	var productDatebase = []Product{
		{Id: 1, Name: "Product1", Price: 1},
		{Id: 2, Name: "Product2", Price: 2},
		{Id: 3, Name: "Product3", Price: 3},
		{Id: 4, Name: "Product4", Price: 4},
	}

	client.AddToBasket(productDatebase[2], 1,)
	client.AddToBasket(productDatebase[3], 2)
	fmt.Println("the client in the basket",client.Basket)
	fmt.Println("Total",client.Total(productDatebase[3], 2))
}
