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
	teamA := shuffledPlayers[0:teamSize]
	teamB := shuffledPlayers[teamSize:len(shuffledPlayers)]

	// Print teams
	fmt.Println("\nHALO")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Team A:", teamA)
	fmt.Println("Team B:", teamB)
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
