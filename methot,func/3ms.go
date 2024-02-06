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
	Quantity int
}

func (c *Client) AddToBasket(product Product, quantity int) {
	for i, item := range c.Basket {
		if item.ProductId == product.Id {
			c.Basket[i].Quantity += quantity
			return
		}
	}

	c.Basket = append(c.Basket, BasketProduct{ProductId: product.Id, Quantity: quantity})
}

func (c *Client) Total() int {
	total := 0
	for _, item := range c.Basket {
		for _, product := range productDatabase {
			if item.ProductId == product.Id {
				total += item.Quantity * product.Price
			}
		}
	}
	return total
}


func (c *Client) Remove(product Product, quantity int) {
	for i, item := range c.Basket{
		if c.Basket[i].Id=product.Id{

		}
		
		}
}



var productDatabase = []Product{
	{Id: 1, Name: "Product1", Price: 100},
	{Id: 2, Name: "Product2", Price: 200},
	{Id: 3, Name: "Product3", Price: 300},
	{Id: 4, Name: "Product4", Price: 400},
}

func main() {
	client := Client{
		Id:   1,
		Name: "client1",
		Basket: []BasketProduct{
			BasketProduct{
			ProductId:3,
			Quantity :3,
			},
		},
	}

	client.AddToBasket(productDatabase[0], 50)
	client.AddToBasket(productDatabase[1], 75)

	fmt.Println("The client's basket:", client.Basket)
	fmt.Println("Total:", Client.Total())
}
