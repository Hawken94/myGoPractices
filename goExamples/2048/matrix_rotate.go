package main

import (
	"fmt"
)

type g2048 [4][4]int

func (g *g2048) mirrorV() {
	gn := new(g2048)
	for i, line := range g {
		for j, num := range line {
			gn[len(g)-i-1][j] = num
		}
	}
	*g = *gn
}

func (g *g2048) right90() {
	gn := new(g2048)
	for i, line := range g {
		for j, num := range line {
			gn[j][len(g)-i-1] = num
		}
	}
	*g = *gn
}

func (g *g2048) left90() {
	gn := new(g2048)
	for i, line := range g {
		for j, num := range line {
			gn[len(g)-j-1][i] = num
		}
	}
	*g = *gn
}
func (g *g2048) right180() {
	gn := new(g2048)
	for i, line := range g {
		for j, num := range line {
			gn[len(g)-i-1][len(g)-j-1] = num
		}
	}
	*g = *gn
}

func (g *g2048) print() {
	for _, line := range g {
		for _, num := range line {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
	fmt.Println()

	gn := g2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
	*g = gn
}

func main() {
	fmt.Println("Origin")
	g := g2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
	g.print()

	fmt.Println("MirrorV")
	g.mirrorV()
	g.print()

	fmt.Println("Left90")
	g.left90()
	g.print()

	fmt.Println("Right90")
	g.right90()
	g.print()

	fmt.Println("Right180")
	g.right180()
	g.print()

}
