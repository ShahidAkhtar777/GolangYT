package board

import (
	"errors"
)

// Board represents a snake and ladder board
type Board struct {
	size    int
	snakes  map[int]int
	ladders map[int]int
}

// NewBoard creates a new Board object with the given size, snakes and ladders
func NewBoard(size int, snakes map[int]int, ladders map[int]int) (*Board, error) {
	if size <= 0 {
		return nil, errors.New("board size must be greater than zero")
	}

	if err := validateSnakes(snakes, size); err != nil {
		return nil, err
	}

	if err := validateLadders(ladders, size); err != nil {
		return nil, err
	}

	return &Board{
		size:    size,
		snakes:  snakes,
		ladders: ladders,
	}, nil
}

// validateSnakes validates that all snakes have valid positions
func validateSnakes(snakes map[int]int, size int) error {
	for head, tail := range snakes {
		if head <= 0 || head > size {
			return errors.New("snake head position is out of bounds")
		}

		if tail <= 0 || tail >= head {
			return errors.New("snake tail position is out of bounds or greater than or equal to head position")
		}
	}

	return nil
}

// validateLadders validates that all ladders have valid positions
func validateLadders(ladders map[int]int, size int) error {
	for start, end := range ladders {
		if start <= 0 || start >= size {
			return errors.New("ladder start position is out of bounds")
		}

		if end <= 0 || end <= start {
			return errors.New("ladder end position is out of bounds or less than or equal to start position")
		}
	}

	return nil
}

// GetNewPosition calculates the new position on the board after rolling the die
func (b *Board) GetNewPosition(currentPosition, dieRoll int) int {
	newPosition := currentPosition + dieRoll

	// check if the new position is beyond the board size
	if newPosition > b.size {
		return currentPosition
	}

	// check if the new position has a snake or ladder
	if tail, ok := b.snakes[newPosition]; ok {
		return tail
	} else if end, ok := b.ladders[newPosition]; ok {
		return end
	}

	return newPosition
}
