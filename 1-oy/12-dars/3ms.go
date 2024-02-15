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

//  for i, item := range b.Basket {
//   if item.ProductId == product.Id {
// 	sum:=b.Client.BasketProducts.Quantity*Product.Price
//    return sum

//   }
//  }
//  b.Basket = append(b.Basket, BasketProducts{ProductId: product.Id, Quantity: quantity})

// }

// func main() {
//  sum := 1

//  product1 := Product{
//   Id:    12,
//   Name:  "melon",
//   Price: 1000,
//  }
//  product2 := Product{
//   Id:    13,
//   Name:  "watermelon",
//   Price: 1020,
//  }
//  product3 := Product{
//   Id:    14,
//   Name:  "banana",
//   Price: 1200,
//  }

//  client1 := Client{
//   Name: "Abror",
//   Id:   123,
//   Basket: []BasketProducts{
//    BasketProducts{
//     ProductId: 12,
//     Quantity:  1,
//    },
//   },
//  }
//  client2 := Client{
//   Name: "Bexruz",
//   Id:   133,
//   Basket: []BasketProducts{
//    BasketProducts{
//     ProductId: 13,
//     Quantity:  2,
//    },
//   },
//  }
//  client3 := Client{
//   Name: "Alisa",
//   Id:   143,
//   Basket: []BasketProducts{
//    BasketProducts{
//     ProductId: 14,
//     Quantity:  3,
//    },
//   },
//  }
//  fmt.Printf("the price of the product in the basket:")
//  add(&client1, product1, sum)
//  add(&client2, product2, sum)
//  add(&client3, product3, sum)

//  fmt.Println("Client 1 the price of the product in the basket:", sum)
//  fmt.Println("Client 2 the price of the product in the basket:", sum)
//  fmt.Println("Client 3 the price of the product in the basket:", sum)

// }
