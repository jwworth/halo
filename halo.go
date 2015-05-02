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

	// Team size
	teamSize := (len(os.Args) / 2)

	// Shuffle the array
	shuffledArray := shuffle(players)

	// Build the first team
	teamA := shuffledArray[0:teamSize]
	fmt.Println("Team A:", teamA)

	// Build the second team
	teamB := shuffledArray[teamSize:len(shuffledArray)]
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
