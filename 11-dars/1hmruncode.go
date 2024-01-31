package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name          string
	Chance        int
	Favoritenumber int
}

type Game struct {
	Number int
	Player Player
}

func (g *Game) NewGame(player Player) {
	s := player.Chance*3 + 1
	g.Number = rand.Intn(s)
	g.Player = player
}

func (p *Player) NewPlayer(name string, chance int, favoritenumber int) {
	p.Name = name
	p.Chance = chance
	p.Favoritenumber = favoritenumber
}

func main() {
	playerName := getPlayerName()
	player := Player{}
	player.NewPlayer(playerName, getPlayerChance(), getPlayerFavoritenumber())
	game := Game{}
	game.NewGame(player)
	game.StartGame()
}

func getPlayerName() string {
	var playerName string
	fmt.Println("Game is starting!")
	fmt.Print("Please enter your name: ")
	_, err := fmt.Scanln(&playerName)
	if err != nil {
		panic(err)
	}
	return playerName
}

func getPlayerChance() int {
	var playerChance int
	fmt.Print("Please enter chance: ")
	_, err := fmt.Scanln(&playerChance)
	if err != nil {
		panic(err)
	}
	return playerChance
}

func getPlayerFavoritenumber() int {
	var playerFavoritenumber int
	fmt.Print("Please enter Favorite number: ")
	_, err := fmt.Scanln(&playerFavoritenumber)
	if err != nil {
		panic(err)
	}
	
	return playerFavoritenumber
}

func (g *Game) StartGame() {
	chance := g.Player.Chance
	s := g.Player.Chance*3 + 1
	fmt.Printf("Choose a number between 0 and %v\n", s)
	for i := 1; i <= chance; i++ {
		fmt.Printf("Your %v try. Please enter the number: ", i)
		var numb int
		_, err := fmt.Scanln(&numb)
		if err != nil {
			panic(err)
		}

		if g.Number == numb && g.Player.Favoritenumber != numb {
			fmt.Println(g.Player.Name,"You win!")
			break
		} else if g.Number == numb && g.Player.Favoritenumber == numb {
			fmt.Printf(g.Player.Name,"You win! The number I chose was your favorite number %v\n", g.Number)
			break
		} else if i == chance {
			fmt.Printf(g.Player.Name,"You are a loser xD. Predicted number is %v\n", g.Number)
		} else if g.Number != numb && numb > g.Number {
			g.Player.Chance--
			fmt.Printf("You didn't find. You have %v chance(s). Enter a smaller number. Try again:)\n", g.Player.Chance)
		} else {
			g.Player.Chance--
			fmt.Printf("You didn't find. You have %v chance(s). Enter a bigger number. Try again:)\n", g.Player.Chance)
		}
	}
}
