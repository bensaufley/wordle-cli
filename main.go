package main

import (
	"fmt"

	"github.com/bensaufley/wordle-cli/internal/game"
	"github.com/bensaufley/wordle-cli/internal/input"
	"github.com/pkg/term"
)

func main() {

	// Check for TTY
	t, err := term.Open("/dev/tty")
	if err != nil {
		panic(err)
	}
	t.Close()

	for {
		g := game.New()

		g.Play()

		if g.Won() {
			fmt.Println("You won! 🎉")
		} else {
			fmt.Printf("Too bad. 😢 Want to know the word? (Yn) ")
			if input.Confirm(true) {
				fmt.Println(g.Word)
			}
		}

		fmt.Printf("Would you like to play again? (yN) ")
		if input.Confirm(false) {
			fmt.Printf("\n")
		} else {
			fmt.Printf("\n")
			break
		}
	}
}
