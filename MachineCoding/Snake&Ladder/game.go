package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	winTile    = 100
	numSnakes  = 6
	numLadders = 6
)

var (
	ladderTops  = []int{14, 31, 38, 84, 59, 92}
	ladderTiles = []int{4, 9, 20, 28, 40, 63}
	snakeHeads  = []int{7, 34, 19, 60, 36, 73}
	snakeTiles  = []int{17, 54, 62, 64, 87, 93}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerPos := 1

	for {
		diceRoll := rand.Intn(6) + 1
		fmt.Printf("You rolled a %d!\n", diceRoll)

		playerPos += diceRoll

		// Check for winning
		if playerPos >= winTile {
			fmt.Println("Congratulations, you won!")
			break
		}

		// Check for snakes and ladders
		for i := 0; i < numSnakes; i++ {
			if playerPos == snakeTiles[i] {
				playerPos = snakeHeads[i]
				fmt.Println("Oh no, you landed on a snake!")
				break
			}
		}

		for i := 0; i < numLadders; i++ {
			if playerPos == ladderTiles[i] {
				playerPos = ladderTops[i]
				fmt.Println("Great, you landed on a ladder!")
				break
			}
		}

		fmt.Printf("You are now on tile %d.\n", playerPos)
	}
}
