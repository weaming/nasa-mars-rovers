package rovers

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction string

const (
	UD Direction = "UD" // unknown direction
	E  Direction = "E"
	N  Direction = "N"
	W  Direction = "W"
	S  Direction = "S"
)

type Location struct {
	X, Y int
}

type Boundary = Location

type State struct {
	Location  *Location
	Direction Direction
}

func (s *State) String() string {
	return fmt.Sprintf("<State %v, %v>", s.Location, s.Direction)
}
func (s *State) Unmarshal() string {
	return fmt.Sprintf("%v %v %s", s.Location.X, s.Location.Y, s.Direction)
}
func (s *State) Clone() *State {
	return &State{&Location{s.Location.X, s.Location.Y}, s.Direction}
}

type Action string

const (
	UA Action = "UA" // unknown action
	L  Action = "L"
	R  Action = "R"
	M  Action = "M"
)

// input: 5 5
func ParseLocation(s string) *Location {
	xs := strings.SplitN(s, " ", 2)
	if len(xs) != 2 {
		return nil
	}

	x := int(Int64(xs[0]))
	y := int(Int64(xs[1]))

	return &Location{X: x, Y: y}
}

// input: 1 2 N
func ParseState(s string) *State {
	xs := strings.SplitN(s, " ", 3)
	if len(xs) != 3 {
		return nil
	}

	x := int(Int64(xs[0]))
	y := int(Int64(xs[1]))
	d := ParseDirection(xs[2])

	if d == UD {
		return nil
	}

	return &State{&Location{x, y}, d}
}

func ParseDirection(s string) Direction {
	switch s {
	case "E":
		return E
	case "N":
		return N
	case "W":
		return W
	case "S":
		return S
	}
	return UD
}

func ParseAction(s byte) Action {
	switch s {
	case 'L':
		return L
	case 'R':
		return R
	case 'M':
		return M
	}
	return UA
}

func ParseActions(s string) (xs []Action) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		xs = append(xs, ParseAction(c))
	}
	return
}

func Int64(v string) int64 {
	i, e := strconv.ParseInt(v, 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}
