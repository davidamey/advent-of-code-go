package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

func main() {
	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	c := &cart{v: vector.New(0, 1)}

	g := grid.New[rune]()
	for y, l := range input {
		for x, r := range l {
			g.SetAt(x, y, r)

			if y == 0 && r == '|' {
				c.p = vector.New(x, y)
			}
		}
	}

	for {
		c.steps++
		c.p = c.p.Add(c.v)
		ch := g.Get(c.p)

		if ch == ' ' {
			break
		}

		if ch != '+' {
			if ch >= 'A' && ch <= 'Z' {
				c.path = append(c.path, ch)
			}
			continue
		}

		switch {
		case c.v.X == 0 && g.Get(c.p.Left()) != ' ':
			c.v = vector.New(-1, 0)
		case c.v.X == 0 && g.Get(c.p.Right()) != ' ':
			c.v = vector.New(1, 0)
		case c.v.Y == 0 && g.Get(c.p.Up()) != ' ':
			c.v = vector.New(0, -1)
		case c.v.Y == 0 && g.Get(c.p.Down()) != ' ':
			c.v = vector.New(0, 1)
		}
	}

	fmt.Println("p1=", string(c.path))
	fmt.Println("p2=", c.steps)
}

type cart struct {
	p, v  vector.Vec
	path  []rune
	steps int
}
