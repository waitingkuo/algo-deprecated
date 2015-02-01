package algoutil

import (
	"bufio"
	"os"
)

type Input struct {
	scanner *bufio.Scanner
}

func NewInput(f *os.File) *Input {

	scanner := bufio.NewScanner(f)

	input := &Input{
		scanner: scanner,
	}

	return input

}

func NewInputFromFile(filename string) (*Input, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return NewInput(f), nil
}

func (input *Input) Iter() <-chan string {

	ch := make(chan string)

	go func() {
		for input.scanner.Scan() {
			ch <- input.scanner.Text()
		}
		close(ch)
	}()

	return ch
}
