package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var leaderboards = []Player{}

func main() {
	game := Game{name: "Roll the Dice", round: 5, Point: Point{plus: 5, minus: 3}}
	var gameActivity Activity
	gameActivity = game

	gameActivity.showHeader()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("How many players who want to play this game : ")
	scanner.Scan()
	totalPlayers, _ := strconv.Atoi(scanner.Text())

	for playerTurn := 1; playerTurn <= totalPlayers; playerTurn++ {
		newPlayer := gameActivity.play(playerTurn)
		leaderboards = append(leaderboards, newPlayer)
	}

	sortingLeaderboards()
	viewLeaderboards()
}
