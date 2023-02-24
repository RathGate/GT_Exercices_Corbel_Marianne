// !This file contains all functions not necessarily asked by the instructions,
// !but kinda necessary for the efficiency of the program.

// !These functions are mostly related to input taking and processing.

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Asks the user for an int.
// Over `max` and under `min` values, the input is discarded and the user is
// asked to enter another value.
func AskUserInt(min, max int) (answer int) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("   → ")

		if _, err := fmt.Fscanln(stdin, &answer); err != nil {
			stdin.Discard(stdin.Buffered())
			continue
		}
		if answer < min || answer > max {
			stdin.Discard(stdin.Buffered())
			continue
		}
		break
	}
	return answer
}

// Asks user for a string.
// This function uses bufio method ReadString in order to keep spaces in the string input,
// but this method adds mandatory newline at the end of the string and automatically
// cancels escape sequences (`\n` is changed to `\\n`).
// The function calls ConvertEscapeChars and RemoveLastNewline to revert these changes.
func AskUserString() (answer string) {
	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("   → ")
	answer, _ = stdin.ReadString('\n')

	stdin.Discard(stdin.Buffered())
	return ConvertEscapeChars(RemoveLastNewLine(answer))
}

// Removes the mandatory \n added at the end of the string by bufio Scanner.
func RemoveLastNewLine(baseStr string) string {
	byteArr := []byte(baseStr)
	return string(byteArr[:len(byteArr)-2])
}

// Converts \\n and \\t back to their original escape sequence.
// ? NOTE: For now, this function only converts \\t into \t and
// ? \n into to \r\n (proper carriage return used in Windows, CRLF).
func ConvertEscapeChars(baseStr string) string {
	byteArr := []byte(baseStr)
	for i := 0; i < len(byteArr); i++ {
		// Looking for backslash:
		if byteArr[i] == byte(92) {
			// Double backslash: not an escape character.
			if i != 0 && byteArr[i-1] == byte(92) {
				continue
			}
			// Horizontal tab:
			if i != len(byteArr)-1 && byteArr[i+1] == byte(116) {
				byteArr = append(append(byteArr[:i], byte(9)), byteArr[i+2:]...)
				continue
			}

			// Newline feed:
			if i != len(byteArr)-1 && byteArr[i+1] == byte(110) {
				byteArr = append(append(byteArr[:i], []byte{13, 10}...), byteArr[i+2:]...)
				continue
			}
		}
	}
	return string(byteArr)
}

// !MENU

// Print une petite boîte de dialogue.
func DrawTitleBox(width int, content string, title string) {
	// Ajoute un espace de padding horizontal au texte `title` s'il existe:
	if len(title) != 0 {
		title = " " + title + " "
	}

	// Positionne le texte `content`` au centre de la boîte:
	missing := width - len(content)
	if missing <= 0 {
		missing = 0
	}
	content = strings.Repeat(" ", missing/2) + content + strings.Repeat(" ", (missing/2+missing%2))

	// Print les trois lignes successives de la boîte et son contenu:
	fmt.Println("┌" + title + strings.Repeat("─", width-len(title)) + "┐")
	fmt.Println("|" + content + "|")
	fmt.Println("└" + strings.Repeat("─", width) + "┘")
}

// !SMALL MENUING

func ConfirmMenu(optionName string) {
	fmt.Printf("0 // %v\n", optionName)
	_ = AskUserInt(0, 0)
	fmt.Println()
}
