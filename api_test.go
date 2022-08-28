package rovers

import (
	"testing"
)

func TestRun(t *testing.T) {
	results := Run(Lines("input.txt"))
	expects := Lines("expects.txt")
	for {
		state, ok := <-results
		if !ok {
			break
		}

		expect := <-expects
		if state.Unmarshal() != expect {
			t.Errorf("result: %s; want %s", state, expect)
			t.Fail()
		} else {
			t.Logf("got: %s", state.Unmarshal())
		}
	}
}
