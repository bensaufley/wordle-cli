package game

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/bensaufley/wordle/internal/input"
)

//go:embed words.json
var wordbytes []byte

var words []string
var wordlen *big.Int

func init() {
	if err := json.Unmarshal(wordbytes, &words); err != nil {
		panic(err)
	}
	wordlen = big.NewInt(int64(len(words)))
}

const maxGuesses = 6
const wordLength = 5

type Game struct {
	Word         string
	guesses      []string
	currentGuess string
}

func (g *Game) Play() {
	fmt.Printf("\x1b[?25l") // hide cursor
	for g.CanGuess() {
		g.DisplayBoard()

		char, err := input.GetChar()
		if err != nil {
			panic(err)
		}
		if char == 3 { // sigint
			os.Exit(130)
		}

		fmt.Print("\x1b8")
		currguesslen := len(g.currentGuess)
		switch char {
		case 13: // return
			if currguesslen < wordLength {
				continue
			}
			if guessInWords(g.currentGuess) {
				g.guesses = append(g.guesses, g.currentGuess)
				g.currentGuess = ""
			}
		case 127: // backspace
			if currguesslen > 0 {
				g.currentGuess = g.currentGuess[0:(currguesslen - 1)]
			}
		default:
			if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
				continue
			}
			if currguesslen < wordLength {
				g.currentGuess = strings.ToLower(fmt.Sprintf("%s%c", g.currentGuess, char))
			}
		}
	}
	g.DisplayBoard()
	fmt.Printf("\x1b[?25h") // show cursor
}
