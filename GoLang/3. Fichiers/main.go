package main

import (
	"fichiers/utils"
	"fmt"
	"os"
	"regexp"
	"time"
)

func OpenFileMenu(filename string) {
	for {
		// Intro lines:
		utils.DrawTitleBox(48, "F I L E  M E N U", "File Editor")

		fmt.Printf("\nCurrently editing file `%v`.\n\n", filename)
		fmt.Println("1 // Display content\n2 // Append text to the end of file\n3 // Empty file content\n4 // Replace file content\n5 // Delete file\n\n0 // Back")

		// Actual menu:
		answerInt := utils.AskUserInt(0, 5)
		fmt.Println()

		switch answerInt {
		case 1:
			utils.ReadFile(filename)
		case 2:
			utils.AppendToFile(filename)
		case 3:
			utils.EmptyFile(filename)
		case 4:
			utils.OverwriteFile(filename)
		case 5:
			if utils.DeleteFile(filename) {
				// Leaves the file menu if the deletion was successful.
				return
			}
		default:
			return
		}
	}
}

func MainMenu() {
	for {
		// Intro lines:
		utils.DrawTitleBox(48, "M A I N  M E N U", "File Editor")
		fmt.Println("\nWhich file do you wanna edit ?\n\n0 // Quit")

		// User must enter the file they want to create or edit:
		answerStr := utils.AskUserString()

		// If the user enters `0`, leaves the program after a short pause:
		re := regexp.MustCompile("^ *0 *$")
		if re.MatchString(answerStr) {
			fmt.Println("\nClosing the editor... Have a nice day !")
			time.Sleep(time.Second)
			return
		}

		// If the file doesn't exist yet, the user has the opportunity to create it:
		if _, err := os.Stat(answerStr); err != nil {
			fmt.Println("\nThat file doesn't exist yet.\nDo you wanna create it ?\n\n1 // Yes\n0 // No")
			if answerInt := utils.AskUserInt(0, 1); answerInt == 1 {
				// Checks if the creation was successful:
				if err = utils.CreateFile(answerStr); err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Print("\nSuccessfully created file ! ")

				// When the file is created, asks the user if they wanna jump into the editing menu:
				fmt.Println("Do you wanna open the editing menu ?\n\n1 // Yes\n0 // No")
				if answerInt = utils.AskUserInt(0, 1); answerInt == 1 {
					fmt.Println()
					OpenFileMenu(answerStr)
					continue
				}
			}
			fmt.Println()

			// If the file has not been created or has been created but its menu was not chosen,
			// asks again the user for another file, and so on.
			continue
		}
		// The file exists and can be edited:
		OpenFileMenu(answerStr)
	}
}

func main() {
	// I like when main() isn't cluttered... <3
	MainMenu()
}
