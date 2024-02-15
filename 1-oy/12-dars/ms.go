// package main

// import "fmt"

// type Client struct {
//  Name   string
//  Id     int
//  Basket []BasketProducts
// }
// type BasketProducts struct {
//  ProductId int
//  Quantity  int
// }
// type Product struct {
//  Id    int
//  Name  string
//  Price int
// }
// func add(b *Client, product Product, quantity int) {

// 	for i, item := range b.Basket {
// 	 if item.ProductId == product.Id {
   
// 	  b.Basket[i].Quantity += quantity
// 	  return
   
// 	 }
// 	}
// 	b.Basket = append(b.Basket, BasketProducts{ProductId: product.Id, Quantity: quantity})
   
//    }

// func remove(d *Client, Id int) {
// for i, item := range d.Basket {
// 	if item.ProductId==Id{
// 	Client.Basket=append(Client.Basket[:i],Client.Basket[i+1:]...)

// 	}
// }
// }

// func main() {
// 	quantity := 4
   
// 	product1 := Product{
// 	 Id:    12,
// 	 Name:  "melon",
// 	 Price: 1000,
// 	}
// 	product2 := Product{
// 	 Id:    13,
// 	 Name:  "watermelon",
// 	 Price: 1020,
// 	}
// 	product3 := Product{
// 	 Id:    14,
// 	 Name:  "banana",
// 	 Price: 1200,
// 	}
   
// 	client1 := Client{
// 	 Name: "Abror",
// 	 Id:   123,
// 	 Basket: []BasketProducts{
// 	  BasketProducts{
// 	   ProductId: 12,
// 	   Quantity:  1,
// 	  },
// 	 },
// 	}
// 	client2 := Client{
// 	 Name: "Bexruz",
// 	 Id:   133,
// 	 Basket: []BasketProducts{
// 	  BasketProducts{
// 	   ProductId: 13,
// 	   Quantity:  2,
// 	  },
// 	 },
// 	}
// 	client3 := Client{
// 	 Name: "Alisa",
// 	 Id:   143,
// 	 Basket: []BasketProducts{
// 	  BasketProducts{
// 	   ProductId: 14,
// 	   Quantity:  3,
// 	  },
// 	 },
// 	}
	
// 	a:=0
// 	fmt.Print(&client1,a)
// 	add(&client1, product1, quantity)
//     add(&client2, product2, quantity)
//     add(&client3, product3, quantity)

// 	fmt.Println("Client 1 Basket:", client1.Basket)
// 	fmt.Println("Client 2 Basket:", client2.Basket)
// 	fmt.Println("Client 3 Basket:", client3.Basket)
   
