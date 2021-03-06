package game

import (
	"fmt"
	"strings"
)

func (g *Game) DisplayBoard() {
	fmt.Print("\x1b7") // save cursor position
	for i := 0; i < maxGuesses; i++ {
		guessesCount := len(g.guesses)
		if i < guessesCount {
			for j, c := range g.guesses[i] {
				if c == rune(g.Word[j]) {
					fmt.Printf("\x1b[30;42m[%c]\x1b[0m", c)
				} else if strings.ContainsRune(g.Word, c) {
					fmt.Printf("\x1b[30;43m[%c]\x1b[0m", c)
				} else {
					fmt.Printf("\x1b[30;47m[%c]\x1b[0m", c)
				}
			}
		} else if i == guessesCount {
			for j := 0; j < wordLength; j++ {
				c := '_'
				currguesslen := len(g.currentGuess)
				if currguesslen > j {
					c = rune(g.currentGuess[j])
				}
				if currguesslen == j {
					fmt.Printf("\x1b[1;30;46m[_]\x1b[0m")
				} else {
					fmt.Printf("\x1b[1;30;47m[%c]\x1b[0m", c)
				}
			}
		} else {
			fmt.Print("\x1b[100;47m")
			for i := 0; i < wordLength; i++ {
				fmt.Print("[ ]")
			}
			fmt.Print("\x1b[0m")
		}
		fmt.Printf("\n")
	}
}
