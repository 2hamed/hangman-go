package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	game := NewGame("bizzare")

	scanner := bufio.NewScanner(os.Stdin)

	game.PrintStats()
	fmt.Print("\n What's your first guess? ")
	for scanner.Scan() {
		input := scanner.Bytes()
		if len(input) == 0 {
			return
		}

		state := game.Try(rune(input[0]))

		game.PrintStats()

		switch state {
		case STATE_WON:
			fmt.Println("Congrats, you just won the game!")
			return
		case STATE_LOST:
			fmt.Println("Sorry, but you lost!")
			return
		case STATE_REPETETIVE:
			fmt.Println("You've already tried that!")
		}
		fmt.Print("\n Next guess: ")

	}
}
