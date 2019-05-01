// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.

// Player - Baseball player struct
type Player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (player Player) average() float32 {
	return float32(player.atBats) / float32(player.hits)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []Player{
		Player{name: "Player1", atBats: 22, hits: 10},
		Player{name: "Player2", atBats: 33, hits: 10},
		Player{name: "Player3", atBats: 44, hits: 10},
	}

	// Display the batting average for each player in the slice.
	for _, player := range players {
		fmt.Printf("name = %v, avg = %v\n", player.name, player.average())
	}
}
