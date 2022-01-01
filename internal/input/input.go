package input

import (
	"errors"
	"fmt"
	"os"

	"github.com/pkg/term"
)

func GetChar() (rune, error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bts := make([]byte, 3)
	defer func() {
		t.Restore()
		t.Close()
	}()

	numRead, err := t.Read(bts)
	if err != nil {
		return 0, err
	}

	if numRead != 1 {
		return 0, errors.New("invalid number of bytes in character")
	}

	return rune(bts[0]), nil
}

func Confirm(dflt bool) bool {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	defer func() {
		t.Restore()
		t.Close()
	}()

	for {
		bts := make([]byte, 1)
		t.Read(bts)
		switch bts[0] {
		case 3:
			os.Exit(130)
		case 13:
			output := "n"
			if dflt {
				output = "y"
			}
			fmt.Println(output)
			return dflt
		case 'y', 'Y':
			fmt.Println("y")
			return true
		case 'n', 'N':
			fmt.Println("n")
			return false
		}
	}
}
