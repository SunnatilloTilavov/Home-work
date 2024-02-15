// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type Player struct {
// 	Name   string
// 	Chance int
// }

// type Game struct {
// 	Number int
// 	Player Player
// }

// func (g *Game) NewGame(player Player) {
// 	g.Number = rand.Intn(10)
// 	g.Player = player
// }

// func (p *Player) NewPlayer(name string, chance int) {
// 	p.Name = name
// 	p.Chance = chance
// }

// func main() {
// 	playerName := getPlayerName()
// 	player := Player{}
// 	player.NewPlayer(playerName, 3)
// 	game := Game{}
// 	game.NewGame(player)
// 	game.StartGame()
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type Player struct {
// 	Name   string
// 	Chance int
// }

// type Game struct {
// 	Number int
// 	Player Player
// }

// func (g *Game) NewGame(player Player) {
// 	g.Number = rand.Intn(player.Chance*player.Chance+1)
// 	g.Player = player
// }

// func (p *Player) NewPlayer(name string, chance int) {
// 	p.Name = name
// 	p.Chance = chance1
// }

// func main() {
// 	playerName := getPlayerName()
// 	player := Player{}
// 	player.NewPlayer(playerName, playerChance)
// 	game := Game{}
// 	game.NewGame(player)
// 	game.StartGame()
// }

// func getPlayerName() string {
// 	var playerName string
// 	fmt.Println("Game is starting!")
// 	fmt.Print("Please enter your name: ")
// 	_, err := fmt.Scanln(&playerName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return playerName
// }
// func getChance() string {
// 	var playerChance string
// 	fmt.Print("Please enter chanse:")
// 	_, err := fmt.Scanln(&playerChance)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return playerChance
// }

// func (g *Game) StartGame() {
// 	var numb int
// 	chance := g.Player.Chance
// 	for i := 1; i <= int(chance); i++ {
// 		fmt.Printf("Your %v try. Please enter the number: ", i)
// 		_, err := fmt.Scanln(&numb)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if g.Number == numb {
// 			fmt.Println("You win!")
// 			break
// 		} else if i == int(chance) {
// 			fmt.Printf("You are a loser xD. Predicted number is %v\n", g.Number)
// 		} else {
// 			g.Player.Chance--
// 			fmt.Printf("You didn't find. You have %v chance(s). Try again:)\n", g.Player.Chance)
// 		}
// 	}
// }


// func (g *Game) StartGame() {
// 	var numb int
// 	chance := g.Player.Chance
// 	for i := 1; i <= int(chance); i++ {
// 		fmt.Printf("Your %v try. Please enter the number: ", i)
// 		_, err := fmt.Scanln(&numb)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if g.Number == numb {
// 			fmt.Println("You win!")
// 			break
// 		} else if i == int(chance) {
// 			fmt.Printf("You are a loser xD. Predicted number is %v\n", g.Number)
// 		} else {
// 			g.Player.Chance--
// 			fmt.Printf("You didn't find. You have %v chance(s). Try again:)\n", g.Player.Chance)
// 		}
// 	}
// }
