package main

import (
	"testing"
)

func Test_stack(t *testing.T) {
	stack := Stack{}
	stack.push([2]int{1, 0})
	stack.push([2]int{2, 1})

	if stack.peek() != [2]int{2, 1} {
		t.Fatal()
	}

	if len(stack) != 2 {
		t.Fatal()
	}

	stored := stack.pop()
	if stored != [2]int{2, 1} {
		t.Fatal()
	}

	if stack.peek() != [2]int{1, 0} {
		t.Fatal()
	}

	if len(stack) != 1 {
		t.Fatal()
	}
}
