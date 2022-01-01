package game

import (
	"crypto/rand"
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

func randomWord() string {
	index, _ := rand.Int(rand.Reader, wordlen)
	return words[index.Int64()]
}

func guessInWords(guess string) bool {
	for _, word := range words {
		if word == guess {
			return true
		}
	}
	return false
}

type Game struct {
	Word         string
	guesses      []string
	currentGuess string
}

func New() *Game {
	g := &Game{
		Word:    randomWord(),
		guesses: []string{},
	}
	return g
}

func (g *Game) CanGuess() bool {
	guesslen := len(g.guesses)
	if guesslen == 0 {
		return true
	}
	if guesslen == 5 || g.Won() {
		return false
	}
	return true
}

func (g *Game) Won() bool {
	return g.guesses[len(g.guesses)-1] == g.Word
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
			if currguesslen < 5 {
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
			if currguesslen < 5 {
				g.currentGuess = strings.ToLower(fmt.Sprintf("%s%c", g.currentGuess, char))
			}
		}
	}
	g.DisplayBoard()
	fmt.Printf("\x1b[?25h") // show cursor
}
