package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// Get players
	players := os.Args[1:]

	// Calculate the midpoint of the players array
	teamSize := (len(players) / 2)

	// Shuffle the players
	shuffledPlayers := shuffle(players)

	// Build teams
	frontRoom := shuffledPlayers[0:teamSize]
	backRoom := shuffledPlayers[teamSize:len(shuffledPlayers)]

	// Define maps
	halo_map := []string{
		"asylum",
		"cage",
		"pinnacle",
		"powerhouse",
		"sword base",
		"zealot",
	}

	// Define game types
	game_type := []string{
		"oddball",
		"slayer",
		"swat",
		"team slayer",
		"team swat",
	}

	// Print teams
	fmt.Println("\nHALO")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Front Room:", frontRoom)
	fmt.Println("Back Room:", backRoom, "\n")

	// Print starting conditions
	fmt.Println("Starting with:", sample(halo_map), "-", sample(game_type))
}

// Source: https://www.socketloop.com/tutorials/golang-shuffle-strings-array
func shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}

	return final
}

func sample(array []string) string {
	return array[rand.Intn(len(array))]
}
