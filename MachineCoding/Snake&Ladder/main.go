// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"/MachineCoding/Snake&Ladder/board/board.go"
// 	"github.com/username/snakes-and-ladders/game"
// )

// func main() {
// 	// Read input from command line arguments
// 	numSnakes, snakes, err := readSnakes(os.Args[1:])
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	numLadders, ladders, err := readLadders(os.Args[numSnakes+2:])
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	numPlayers, players, err := readPlayers(os.Args[numSnakes+numLadders+3:])
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Create game board
// 	b := board.NewBoard(100, snakes, ladders)

// 	// Create game instance
// 	g := game.NewGame(b, players)

// 	// Start game
// 	g.Play()
// }

// func readSnakes(args []string) (int, []board.Snake, error) {
// 	numSnakes, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		return 0, nil, fmt.Errorf("invalid number of snakes: %s", args[0])
// 	}

// 	snakes := make([]board.Snake, numSnakes)
// 	for i := 0; i < numSnakes; i++ {
// 		h, err1 := strconv.Atoi(args[i*2+1])
// 		t, err2 := strconv.Atoi(args[i*2+2])
// 		if err1 != nil || err2 != nil {
// 			return 0, nil, fmt.Errorf("invalid snake position: %s, %s", args[i*2+1], args[i*2+2])
// 		}
// 		snakes[i] = board.Snake{Head: h, Tail: t}
// 	}

// 	return numSnakes, snakes, nil
// }

// func readLadders(args []string) (int, []board.Ladder, error) {
// 	numLadders, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		return 0, nil, fmt.Errorf("invalid number of ladders: %s", args[0])
// 	}

// 	ladders := make([]board.Ladder, numLadders)
// 	for i := 0; i < numLadders; i++ {
// 		s, err1 := strconv.Atoi(args[i*2+1])
// 		e, err2 := strconv.Atoi(args[i*2+2])
// 		if err1 != nil || err2 != nil {
// 			return 0, nil, fmt.Errorf("invalid ladder position: %s, %s", args[i*2+1], args[i*2+2])
// 		}
// 		ladders[i] = board.Ladder{Start: s, End: e}
// 	}

// 	return numLadders, ladders, nil
// }

// func readPlayers(args []string) (int, []game.Player, error) {
// 	numPlayers, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		return 0, nil, fmt.Errorf("invalid number of players: %s", args[0])
// 	}

// 	players := make([]game.Player, numPlayers)
// 	for i := 0; i < numPlayers; i++ {
// 		players[i] = game.Player{name: args[i+1], position: 0}
// 	}

// 	return numPlayers, players, nil
// }

package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/username/snakes-and-ladders/board"
	"github.com/username/snakes-and-ladders/player"
)

func main() {
	// Read input from file or command line
	input, err := readInput()
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Create the board
	b := board.NewBoard(input.Snakes, input.Ladders)

	// Create the players
	players := make([]*player.Player, input.NumPlayers)
	for i := 0; i < input.NumPlayers; i++ {
		players[i] = player.NewPlayer(input.PlayerNames[i], i+1)
	}

	// Play the game
	currentPlayer := 0
	for {
		// Roll the die
		steps := RollDie()

		// Move the player
		players[currentPlayer].Move(steps)

		// Check for snake or ladder
		pos, found := b.SnakeOrLadder(players[currentPlayer].Pos)
		if found {
			players[currentPlayer].Pos = pos
		}

		// Print the move
		fmt.Printf("%s rolled a %d and moved from %d to %d\n",
			players[currentPlayer], steps, players[currentPlayer].Pos-steps, players[currentPlayer].Pos)

		// Check if player won
		if players[currentPlayer].Pos == 100 {
			fmt.Printf("%s has won the game!\n", players[currentPlayer])
			break
		}

		// Next player's turn
		currentPlayer = (currentPlayer + 1) % input.NumPlayers
	}
}

func RollDie() int {
	return 1 + (rand.Intn(6))
}

type Input struct {
	Snakes      []board.Snake
	Ladders     []board.Ladder
	NumPlayers  int
	PlayerNames []string
}

func readInput() (Input, error) {
	var input Input

	// Read number of snakes
	var numSnakes int
	_, err := fmt.Scanln(&numSnakes)
	if err != nil {
		return input, err
	}

	// Read snakes
	input.Snakes = make([]board.Snake, numSnakes)
	for i := 0; i < numSnakes; i++ {
		var head, tail int
		_, err := fmt.Scanln(&head, &tail)
		if err != nil {
			return input, err
		}
		input.Snakes[i] = board.Snake{Head: head, Tail: tail}
	}
}
