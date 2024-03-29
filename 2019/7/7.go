package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	// data := util.MustReadCSInts("example")
	// data := util.MustReadCSInts("example2") // p2
	// data := util.MustReadCSInts("example3") // p2
	data := util.MustReadCSInts("input")

	fmt.Println("p1=", p1(data))
	fmt.Println("p2=", p2(data))
}

func p1(data []int) int {
	prog := intcode.Program(data)
	maxE := math.MinInt16
	// var maxEperm []int
	for p := range util.NewPermuter([]int{0, 1, 2, 3, 4}).Permutations() {
		a := prog.Run(p[0], 0)
		b := prog.Run(p[1], a[0])
		c := prog.Run(p[2], b[0])
		d := prog.Run(p[3], c[0])
		e := prog.Run(p[4], d[0])
		if e[0] > maxE {
			maxE = e[0]
			// maxEperm = p
		}
	}

	// fmt.Println(maxE, maxEperm)
	return maxE
}

func p2(data []int) (maxV int) {
	count := 5
	chanSize := 2

	prog := intcode.Program(data)
	for p := range util.NewPermuter([]int{5, 6, 7, 8, 9}).Permutations() {
		in := make(chan int, chanSize)
		in <- p[0]
		in <- 0

		pipes := make([]chan int, count)
		for i := range pipes {
			pipe := make(chan int, chanSize)
			pipes[i] = pipe

			if i+1 < len(pipes) {
				pipe <- p[i+1]
			}
		}

		for i := range pipes {
			go func(i int) {
				if i == 0 {
					prog.RunBuf(in, pipes[i])
				} else {
					prog.RunBuf(pipes[i-1], pipes[i])
				}
			}(i)
		}

		lastV := 0
		for v := range pipes[count-1] {
			lastV = v
			in <- v
		}

		if lastV > maxV {
			maxV = lastV
		}
	}

	return maxV
}
