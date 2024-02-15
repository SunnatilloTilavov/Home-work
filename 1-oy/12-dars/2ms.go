// package main

// import "log"

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

// // func add(b *Client, product Product, quantity int) {

// //  for i, item := range b.Basket {
// //   if item.ProductId == product.Id {

// //    b.Basket[i].Quantity += quantity
// //    return

// //   }
// //  }
// //  b.Basket = append(b.Basket, BasketProducts{ProductId: product.Id, Quantity: quantity})

// // }

// func remove(client *Client, idToRemove int) {
//  for i, item := range client.Basket {
//   if item.ProductId == idToRemove {
//    client.Basket = append(client.Basket[:i], client.Basket[i+1:]...)
//    return
//   }
//  }

// }

// func main() {

//  // client := Client{
//  //  Id:   1,
//  //  Name: "javohir",
//  //  Basket: []BasketProducts{
//  //   {ProductId: 1, Quantity: 2},
//  //   {ProductId: 2, Quantity: 1},
//  //  },
//  // }
//  // productIdToRemove := 1
//  // remove(&client, productIdToRemove)

//  // log.Printf("Client's Basket after removing a product: %+v\n", client.Basket)

//  // product1 := []Product{
//  //  Product{
//  //   Id:    12,
//  //   Name:  "melon",
//  //   Price: 1000,
//  //  },
//  // }
//  // product2 := Product{
//  //  Id:    13,
//  //  Name:  "watermelon",
//  //  Price: 1020,
//  // }
//  // product3 := Product{
//  //  Id:    14,
//  //  Name:  "banana",
//  //  Price: 1200,
//  // }

//  client1 := Client{
//   Name: "Abror",
//   Id:   123,
//   Basket: []BasketProducts{
//    BasketProducts{
//     ProductId: 12,
//     Quantity:  1,
//    },
//    BasketProducts{
//     ProductId: 13,
//     Quantity:  1,
//    },
//   },
//  }
//  // client2 := Client{
//  //  Name: "Bexruz",
//  //  Id:   133,
//  //  Basket: []BasketProducts{
//  //   BasketProducts{
//  //    ProductId: 13,
//  //    Quantity:  2,
//  //   },
//  //  },
//  // }
//  // client3 := Client{
//  //  Name: "Alisa",
//  //  Id:   143,
//  //  Basket: []BasketProducts{
//  //   BasketProducts{
//  //    ProductId: 14,
//  //    Quantity:  3,
//  //   },
//  //  },
//  // }
//  removee := 12
//  remove(&client1, removee)
//  log.Printf("Client's Basket after removing a product: %+v\n", client1.Basket)


//  // add(&client1, product1, quantity)
//  // add(&client2, product2, quantity)
//  // add(&client3, product3, quantity)

//  // fmt.Println("Client 1 Basket:", client1.Basket)
//  // fmt.Println("Client 2 Basket:", client2.Basket)
//  // fmt.Println("Client 3 Basket:", client3.Basket)

// }

