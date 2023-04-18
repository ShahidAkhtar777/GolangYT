package player

import (
	"fmt"
)

// Player represents a player of the game
type Player struct {
	Name   string
	Number int
	Pos    int
}

// Move moves the player by a given number of steps
func (p *Player) Move(steps int) {
	p.Pos += steps
	if p.Pos > 100 {
		p.Pos = 100
	}
}

// NewPlayer creates a new player with the given name and number
func NewPlayer(name string, number int) *Player {
	return &Player{
		Name:   name,
		Number: number,
		Pos:    1,
	}
}

// String returns a string representation of the player
func (p *Player) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Pos)
}
