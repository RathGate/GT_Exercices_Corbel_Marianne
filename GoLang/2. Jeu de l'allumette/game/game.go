package game

import (
	"fmt"
	"time"
)

// Initializes the game (player count, player names, matches count)
func PlayMatchGame() {
	DrawTitleBox(48, " J E U  D E  L ' A L L U M E T T E", "")

	// players := [2]string{"Joueur 1", "Joueur 2"}

	fmt.Print("Règles: Vous décidez d'un nombre total d'allumettes.\nChacun leur tour, les joueurs décident de prendre entre 1 et 3 allumettes.\nLe joueur qui tire la dernière allumette a perdu!\n\n")
	fmt.Println("Choisir un nombre total d'allumettes: ")
	matchCount := AskUserInt(3, 100, true, "Choisir un nombre total d'allumettes:\n")
	fmt.Println()
	for round := 1; matchCount > 0; round++ {
		DrawTitleBox(48, fmt.Sprintf("R O U N D %d", round), "Jeu de l'allumette")

		for currPlayer := 1; currPlayer <= 2; currPlayer++ {
			fmt.Printf("Joueur %d, choisis un nombre d'allumettes:\n", currPlayer)
			var count = AskUserInt(1, 3, true, fmt.Sprintf("Joueur %d, choisis un nombre d'allumettes:\n", currPlayer))

			if (matchCount - count) <= 0 {
				fmt.Printf("\nLe Joueur %d a pris la dernière allumette ! \n\n", currPlayer)
				time.Sleep(1 * time.Second)
				DrawTitleBox(48, fmt.Sprintf("P L A Y E R  %d  W I N S !", winnerIndex(currPlayer)), "Jeu de l'allumette")
				time.Sleep(1 * time.Second)
				return
			}
			matchCount -= count
			fmt.Println()
		}

	}
}
func winnerIndex(losingPlayer int) int {
	if losingPlayer == 1 {
		return 2
	}
	return 1
}
