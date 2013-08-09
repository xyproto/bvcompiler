package main

import (
	"fmt"
	"strings"
)

type CodeTree struct {
	code string
	parent *CodeTree
}

// TODO Output the CodeTree as an indented tree of code
func (ast *CodeTree) String() string {
	return "to be implemented"
}

// TODO Output the CodeTree as a line of code
func (ast *CodeTree) Repr() string {
	return "to be implemented"
}

// TODO Return a CodeTree
func parser(program string) *CodeTree {
	level := 0
	collected := ""
	parent := ""
	parentlevel := 0
	for _, letter := range program {
		if letter == '(' {
			if strings.TrimSpace(collected) != "" {
				fmt.Println(level, collected, "(parent:", parentlevel, parent, ")")
				parent = collected
				parentlevel = level
			}
			level++
			collected = ""
		} else if letter == ')' {
			if strings.TrimSpace(collected) != "" {
				fmt.Println(level, collected, "(parent:", parentlevel, parent, ")")
			}
			level--
			collected = ""
		} else {
			collected += string(letter)
		}
	}
	return new(CodeTree)
}

func main() {
	ast := parser("(lambda (x) (fold x 0 (lambda (y z) (or y z))))")
	fmt.Println(ast.String())
	fmt.Println(ast.Repr())
}
