package rovers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func Run(lines <-chan string) chan *State {
	boundary := ParseLocation(<-lines)
	log.Println("boundary", boundary)

	results := make(chan *State)
	go func() {
		defer close(results)
		for {
			stateStr, ok := <-lines
			if !ok {
				break
			}

			state := ParseState(stateStr)
			actions := ParseActions(<-lines)

			result := DoActions(state, actions, boundary)
			log.Println(state, actions, "=>", result)
			results <- result
		}
	}()
	return results
}

func DoActions(state *State, actions []Action, boundary *Location) *State {
	if state.Direction == UD {
		return nil
	}

	state = state.Clone()

	for _, action := range actions {
		switch action {
		case L:
			switch state.Direction {
			case E:
				state.Direction = N
			case S:
				state.Direction = E
			case W:
				state.Direction = S
			case N:
				state.Direction = W
			}
		case R:
			switch state.Direction {
			case E:
				state.Direction = S
			case S:
				state.Direction = W
			case W:
				state.Direction = N
			case N:
				state.Direction = E
			}
		case M:
			switch state.Direction {
			case E:
				state.Location.X += 1
			case S:
				state.Location.Y -= 1
			case W:
				state.Location.X -= 1
			case N:
				state.Location.Y += 1
			}
		}
	}
	return state
}

func Lines(file string) <-chan string {
	path, e := filepath.Abs(file)
	if e != nil {
		panic(e)
	}

	f, e := os.Open(path)
	if e != nil {
		panic(e)
	}

	ch := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(ch)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()
	return ch
}
