package main

import "testing"

func TestRollingDice(t *testing.T) {
	var diceNumbers = []int{1, 2, 3, 4, 5, 6}
	rolledDice := rollingDice()

	for i := 0; i < len(diceNumbers); i++ {
		if rolledDice == diceNumbers[i] {
			t.Logf("Success, rolled number is %d, output rolled number should contains %v", rolledDice, diceNumbers)
			return
		}
	}

	t.Errorf("Failed, rolled number is %d, output rolled number should contains %v", rolledDice, diceNumbers)
}

func TestCalculateScore(t *testing.T) {
	var score = 0
	var rolledNumbers = []int{3, 6, 1, 4, 5}

	game := Game{name: "Roll the Dice", round: 5, Point: Point{plus: 5, minus: 3}}
	var gameActivity Activity
	gameActivity = game

	for _, rolledNumber := range rolledNumbers {
		gameActivity.calculateScore(rolledNumber, &score)
	}

	if score == 9 {
		t.Logf("Success, score is %d, expected score is %d", score, 9)
	} else {
		t.Errorf("Failed, score is %d, expected score is %d", score, 9)
	}
}

func TestPlay(t *testing.T) {
	game := Game{name: "Roll the Dice", round: 5, Point: Point{plus: 5, minus: 3}}
	var gameActivity Activity
	gameActivity = game

	player := gameActivity.play(1)

	if (Player{} == player) {
		t.Errorf("Failed, function should return a struct Player")
	} else {
		t.Logf("Success, function should return a struct Player")
	}
}

func TestSortingLeaderboards(t *testing.T) {
	leaderboards = []Player{
		{name: "Heri", score: 15},
		{name: "Hakim", score: 20},
		{name: "Setiawan", score: 17},
	}

	sortingLeaderboards()

	for i := 0; i < len(leaderboards)-1; i++ {
		if leaderboards[i].score < leaderboards[i+1].score {
			t.Errorf("Failed, leaderboards should sorted by score - %v", leaderboards)
			return
		}
	}

	t.Logf("Success, leaderboards should sorted by score - %v", leaderboards)
}
