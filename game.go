package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Player is representation for a person and his score
type Player struct {
	name  string
	score int
}

// Game is representation for the game
type Game struct {
	name  string
	round int
	Point
}

// Point is representation of the point rule
type Point struct {
	plus  int
	minus int
}

// Activity is the Interface for Games Activity
type Activity interface {
	showHeader()
	play(playerTurn int) Player
	calculateScore(rolledDice int, score *int) int
}

func rollingDice() int {
	rand.Seed(time.Now().UnixNano())
	var dice = []int{1, 2, 3, 4, 5, 6}
	return dice[rand.Intn(len(dice))]
}

func (game Game) showHeader() {
	fmt.Printf("====================================\n")
	fmt.Printf("Welcome to %s Game\n", game.name)
	fmt.Printf("====================================\n\n")
}

func (game Game) calculateScore(rolledDice int, score *int) int {
	if rolledDice%2 == 0 {
		*score -= game.minus
	} else {
		*score += game.plus
	}
	return *score
}

func (game Game) play(playerTurn int) Player {
	scanner := bufio.NewScanner(os.Stdin)

	var score int = 0

	fmt.Printf("\nPlayer %d, Input your name : ", playerTurn)
	scanner.Scan()
	playerName := scanner.Text()

	for currentRound := 1; currentRound <= game.round; currentRound++ {
		fmt.Printf("\nRound %d.\nOkay %s, press enter to roll the dice", currentRound, playerName)
		scanner.Scan()
		rolledDice := rollingDice()
		fmt.Printf("You've got %d\n", rolledDice)
		game.calculateScore(rolledDice, &score)
	}
	fmt.Printf("====================================\n")
	fmt.Printf("Your score is : %d\n", score)
	fmt.Printf("====================================\n")

	// Create a Player
	var currentPlayer Player
	currentPlayer.name = playerName
	currentPlayer.score = score

	return currentPlayer
}

func viewLeaderboards() {
	fmt.Printf("\n----- Leaderboards -----\n")
	for i, leaderboard := range leaderboards {
		fmt.Printf("%d. %s with Score: %d\n", i+1, leaderboard.name, leaderboard.score)
	}
}

func sortingLeaderboards() {
	for i := 0; i < len(leaderboards); i++ {
		for j := 0; j < len(leaderboards)-i; j++ {
			if j < len(leaderboards)-1 && leaderboards[j].score < leaderboards[j+1].score {
				temp := leaderboards[j]
				leaderboards[j] = leaderboards[j+1]
				leaderboards[j+1] = temp
			}
		}
	}
}
