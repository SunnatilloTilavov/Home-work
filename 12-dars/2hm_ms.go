package main

import "fmt"

type Client struct {
	Name   string
	Id     int
	Basket []BasketProducts
}

type BasketProducts struct {
	ProductId int
	Quantity  int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func add(b *Client, product Product, quantity int) {
	for i, item := range b.Basket {
		if item.ProductId == product.Id {
		
			b.Basket[i].Quantity += quantity
			return
		}
	}
	
	b.Basket = append(b.Basket, BasketProducts{ProductId: product.Id, Quantity: quantity})
}

func countOccurrences(basket []BasketProducts, productId int) int {
	count := 0
	for _, item := range basket {
		if item.ProductId == productId {
			count += item.Quantity
		}
	}
	return count
}

func main() {
	product1 := Product{
		Id:    12,
		Name:  "melon",
		Price: 10,
	}
	product2 := Product{
		Id:    13,
		Name:  "watermelon",
		Price: 15,
	}
	product3 := Product{
		Id:    14,
		Name:  "banana",
		Price: 20,
	}

	client1 := Client{
		Name: "Abror",
		Id:   123,
		Basket: []BasketProducts{
			BasketProducts{
				ProductId: 12,
				Quantity:  1,
			},
			BasketProducts{
				ProductId: 13,
				Quantity:  1,
			},
		},
	}

	client2 := Client{
		Name: "Bexruz",
		Id:   133,
		Basket: []BasketProducts{
			BasketProducts{
				ProductId: 13,
				Quantity:  2,
			},
			BasketProducts{
				ProductId: 14,
				Quantity:  2,
			},
		},
	}

	client3 := Client{
		Name: "Alisa",
		Id:   143,
		Basket: []BasketProducts{
			BasketProducts{
				ProductId: 14,
				Quantity:  3,
			},
			BasketProducts{
				ProductId: 12,
				Quantity:  5,
			},
		},
	}


	add(&client1, product1, client1.Basket[0].Quantity)
	add(&client2, product2, client2.Basket[0].Quantity)
	add(&client3, product3, client3.Basket[0].Quantity)


	count1 := countOccurrences(client1.Basket, product1.Id)
	count2 := countOccurrences(client2.Basket, product2.Id)
	count3 := countOccurrences(client3.Basket, product3.Id)

	
	fmt.Printf("%d %s\n", count1, product1.Name)
	fmt.Printf("%d %s\n",  count2, product2.Name)
	fmt.Printf(" %d %s\n",  count3, product3.Name)
}
