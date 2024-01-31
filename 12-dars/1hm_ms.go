// package main

// import "fmt"

// type Client struct {
// 	Name   string
// 	Id     int
// 	Basket []BasketProducts
// }

// type BasketProducts struct {
// 	ProductId int
// 	Quantity  int
// }

// type Product struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

// func add(b *Client, product Product, quantity int) {
// 	for i, item := range b.Basket {
// 		if item.ProductId == product.Id {
// 			b.Basket[i].Quantity += quantity-1
// 			return
// 		}
// 	}
// 	b.Basket = append(b.Basket, BasketProducts{ProductId: product.Id, Quantity: quantity})
// }

// func main() {
// 	sum1 := 0
// 	sum2 := 0
// 	sum3 := 0

// 	product1 := Product{
// 		Id:    12,
// 		Name:  "melon",
// 		Price: 10,
// 	}
// 	product2 := Product{
// 		Id:    13,
// 		Name:  "watermelon",
// 		Price: 15,
// 	}
// 	product3 := Product{
// 		Id:    14,
// 		Name:  "banana",
// 		Price: 5,
// 	}

// 	client1 := Client{
// 		Name: "Abror",
// 		Id:   123,
// 		Basket: []BasketProducts{
// 			BasketProducts{
// 				ProductId: 12,
// 				Quantity:  1,
// 			},
// 			BasketProducts{
// 				ProductId: 12,
// 				Quantity:  1,
// 			},
// 		},
// 	}

// 	client2 := Client{
// 		Name: "Bexruz",
// 		Id:   133,
// 		Basket: []BasketProducts{
// 			BasketProducts{
// 				ProductId: 13,
// 				Quantity:  2,
// 			},
// 			BasketProducts{
// 				ProductId: 14,
// 				Quantity:  2,
// 			},
// 		},
// 	}

// 	client3 := Client{
// 		Name: "Alisa",
// 		Id:   143,
// 		Basket: []BasketProducts{
// 			BasketProducts{
// 				ProductId: 14,
// 				Quantity:  3,
// 			},
// 			BasketProducts{
// 				ProductId: 14,
// 				Quantity:  3,
// 			},
// 		},
// 	}

// 	add(&client1, product1, client1.Basket[0].Quantity)
// 	add(&client2, product2, client2.Basket[0].Quantity)
// 	add(&client3, product3, client3.Basket[0].Quantity)

// 	for _, item := range client1.Basket {
// 		sum1 += item.Quantity * product1.Price
// 	}
// 	for _, item := range client2.Basket {
// 		sum2 += item.Quantity * product2.Price
// 	}
// 	for _, item := range client3.Basket {
// 		sum3 += item.Quantity * product3.Price
// 	}

// 	fmt.Println(client1.Name, "Client 1 the price of the product in the basket:", sum1)
// 	fmt.Println(client2.Name, "Client 2 the price of the product in the basket:", sum2)
// 	fmt.Println(client3.Name, "Client 3 the price of the product in the basket:", sum3)
// }
