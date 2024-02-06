// package main

// import "fmt"

// type Client struct {
// 	Id     int
// 	Name   string
// 	Basket []BasketProduct
// }

// type BasketProduct struct {
// 	ProductId int
// 	Quantity  int
// }

// type Product struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

// func (c *Client) AddToBasket(product Product, quantity int) {
// 	for i, basketProduct := range c.Basket {
// 		if basketProduct.ProductId == product.Id {
// 			c.Basket[i].Quantity += quantity
// 			return
// 		}
// 	}

	
// 	c.Basket = append(c.Basket, BasketProduct{ProductId: product.Id, Quantity: quantity})
// }

// func main() {
// 	client := Client{
// 		Id:   1,
// 		Name: "client1",
// 	}

// 	var productDatabase = []Product{
// 		{Id: 1, Name: "Product1", Price: 1},
// 		{Id: 2, Name: "Product2", Price: 2},
// 		{Id: 3, Name: "Product3", Price: 3},
// 		{Id: 4, Name: "Product4", Price: 4},
// 	}

// 	client.AddToBasket(productDatabase[1], 1)
// 	client.AddToBasket(productDatabase[2], 2)
// 	fmt.Println("the client in the basket", client.Basket)
// }
