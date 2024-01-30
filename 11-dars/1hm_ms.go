package main
import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name string
	Chanse int
}

type Game struct {
	Number int
	Player Player
}
func (g Game) Newgame (player Player)Game{
	return Game{
		Number:rand.Intn(10),
		Player:player,
	}
}
func (p Player) Newplayer (name string,chanse int) Player{
	return Player{
		Name:name,
		Chanse:chance,
	}
}

func main(){
	getPlayerName()
	player:=Player{}
	player=player.Newplayer(getPlayerName(),3)
	game := Game{}
	game=game.Newgame(player)
	game.StartGame()

}

func getPlayerName () string{
var playerName string
fmt.Println("Game is strating!")
fmt.Print("Plase enter your name:")	
_, err:=fmt.Scanln(&playerName)
if err !=nil{
	panic(err)
}
return playerName
}
func (g Game) StartGame(){
	var numb int
	chanse:=g.Player.Chanse
	for i:=1; i<=int(chanse);i++{
	fmt.Printf("Your %v try.Plase enter the  number:",i)
	if g.Number==numb{
		fmt.Println("You win")
		break
	}else if i==int(chanse){
		fmt.Printf("You are a loser xD,Predicted number is %v\n",g.Number)
	}else{
		g.Player.Chanse--
		fmt.Printf("You didnt find.You have %v chanse.Try again:)\n,",g.Player.Chanse)
	}
}
	}

