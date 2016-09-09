package action

import (
	"fmt"
	"testing"
)

func TestActionChainAdd(t *testing.T) {
	chain := new(Chain)
	functionA := func() {
		fmt.Println("A")
	}
	functionB := func() {
		fmt.Println("B")
	}

	if len(chain.actions) != 0 {
		t.Fail()
	}

	chain.Add(functionA)
	if len(chain.actions) != 1 {
		t.Fail()
	}

	chain.Add(functionB)
	if len(chain.actions) != 2 {
		t.Fail()
	}
}

func TestActionChainExecute(t *testing.T) {
	chain := new(Chain)

	counter := 0
	functionA := func() {
		counter++
	}

	chain.
		Add(functionA).
		Add(functionA).
		Execute()

	if counter != 2 {
		t.Fail()
	}
}
