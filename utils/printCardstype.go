package utils

import (
	"fmt"
)

func printCardsType() {
	for i := 2; i <= 14; i++ {
		fmt.Printf("\"%d%d%d%d%d%d%d%d%d%d%d%d\":{TwelveBomb,%d},\n",
			i, i, i, i, i, i, i, i, i, i, i, i, i)
	}
}

func print510K() {
	for i := 0; i <= 1; i++ {
		fmt.Printf("\"%d%d%d%d%d%d%d%d%d%d%d%d\":{TwelveBomb,%d},\n",
			i, i, i, i, i, i, i, i, i, i, i, i, i)
	}
}
