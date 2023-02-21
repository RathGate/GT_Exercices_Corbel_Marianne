package main

import (
	"allumettes/game"
	"fmt"
	"time"
)

func main() {
	retry := 1
	for retry != 0 {
		game.PlayMatchGame()

		fmt.Println("Do you wanna play again ?\n\n1// Yes\n0// No")
		retry = game.AskUserInt(0, 1, false, "")
		fmt.Println()
	}
	fmt.Println("Bonne journée ♪")
	time.Sleep(2 * time.Second)

}
