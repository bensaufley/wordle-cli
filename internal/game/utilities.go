package game

import "crypto/rand"

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
	if guesslen == maxGuesses || g.Won() {
		return false
	}
	return true
}

func (g *Game) Won() bool {
	return g.guesses[len(g.guesses)-1] == g.Word
}
