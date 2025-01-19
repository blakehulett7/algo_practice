package main

import "fmt"

func main() {
	scratch("")
}

func fmtin() {
	fmt.Println()
}

type Stack []any

func (s *Stack) push(element any) {
	*s = append(*s, element)
}

func (s *Stack) pop() any {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() any {
	return s[len(s)-1]
}

func scratch(str string) bool {

}
