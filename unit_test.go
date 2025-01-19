package main

import (
	"testing"
)

func Test_stack(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(2)

	if stack.peek() != 2 {
		t.Fatal()
	}

	if len(stack) != 2 {
		t.Fatal()
	}

	stored := stack.pop()
	if stored != 2 {
		t.Fatal()
	}

	if stack.peek() != 1 {
		t.Fatal()
	}

	if len(stack) != 1 {
		t.Fatal()
	}
}
