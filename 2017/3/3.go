package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	fmt.Println("examples:")
	for _, i := range []int{1, 12, 23, 1024} {
		fmt.Printf("%d => %d\n", i, p1(i))
	}

	fmt.Println("p1=", p1(277678))
	fmt.Println("p2=", p2(277678))
}

func p1(target int) int {
	ring := 0
	edge := 0
	square := 0
	for {
		edge = 2*ring + 1
		square = edge * edge
		if square >= target {
			break
		}
		ring++
	}

	corner := square
	for {
		c := corner - (edge - 1)
		if c <= target {
			break
		}
		corner = c
	}

	return ring + util.AbsInt(corner-edge/2-target)
}

func p2(target int) int {
	for _, i := range FromOEIS {
		if i > target {
			return i
		}
	}
	return -1
}

//  37  36  35  34  33  32  31
//  38  17  16  15  14  13  30
//  39  18   5   4   3  12  29
//  40  19   6   1   2  11  28
//  41  20   7   8   9  10  27
//  42  21  22  23  24  25  26
//  43  44  45  46  47  48  49

// https://oeis.org/A141481/b141481.txt
var FromOEIS []int = []int{
	1,
	1,
	2,
	4,
	5,
	10,
	11,
	23,
	25,
	26,
	54,
	57,
	59,
	122,
	133,
	142,
	147,
	304,
	330,
	351,
	362,
	747,
	806,
	880,
	931,
	957,
	1968,
	2105,
	2275,
	2391,
	2450,
	5022,
	5336,
	5733,
	6155,
	6444,
	6591,
	13486,
	14267,
	15252,
	16295,
	17008,
	17370,
	35487,
	37402,
	39835,
	42452,
	45220,
	47108,
	48065,
	98098,
	103128,
	109476,
	116247,
	123363,
	128204,
	130654,
	266330,
	279138,
	295229,
	312453,
	330785,
	349975,
	363010,
	369601,
	752688,
	787032,
	830037,
	875851,
	924406,
	975079,
	1009457,
	1026827,
	2089141,
	2179400,
	2292124,
	2411813,
	2539320,
	2674100,
	2814493,
	2909666,
	2957731,
	6013560,
	6262851,
	6573553,
	6902404,
	7251490,
	7619304,
	8001525,
	8260383,
	8391037,
	17048404,
	17724526,
	18565223,
	19452043,
	20390510,
	21383723,
	22427493,
	23510079,
	24242690,
}
