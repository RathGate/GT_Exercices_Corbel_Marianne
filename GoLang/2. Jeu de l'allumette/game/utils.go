package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func AskUserInt(min, max int, displayMessage bool, errMessage string) (answer int) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("   → ")

		if _, err := fmt.Fscanln(stdin, &answer); err != nil {
			if displayMessage {
				fmt.Printf("Veuillez entrer un nombre entier.\n\n")
				time.Sleep(time.Second)
				fmt.Print(errMessage)
			}

			stdin.Discard(stdin.Buffered())
			continue
		}
		if answer < min || answer > max {
			if displayMessage {
				fmt.Printf("Nombre invalide ! [%d:%d]\n\n", min, max)
				time.Sleep(time.Second)
				fmt.Print(errMessage)
			}

			stdin.Discard(stdin.Buffered())
			continue
		}

		break
	}
	return answer
}

func DrawTitleBox(width int, content string, title string) {

	// Ajoute un espace de padding horizontal au texte `title` s'il existe:
	if len(title) != 0 {
		title = " " + title + " "
	}

	// Positionne le texte `content au centre de la boîte:
	missing := width - len(content)
	if missing <= 0 {
		missing = 0
	}
	content = strings.Repeat(" ", missing/2) + content + strings.Repeat(" ", (missing/2+missing%2))

	// Print les trois lignes successives de la boîte et son contenu:
	fmt.Println("┌" + title + strings.Repeat("─", width-len(title)) + "┐")
	fmt.Println("|" + content + "|")
	fmt.Println("└" + strings.Repeat("─", width) + "┘" + "\n")
}
