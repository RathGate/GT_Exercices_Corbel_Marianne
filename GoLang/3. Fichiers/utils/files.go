// !NOTE: This file contains all functions making changes to the file chosen.
// !They are all under the form of a menu.

package utils

import (
	"fmt"
	"os"
)

func ReadFile(filename string) {
	// Cosmetic purpose only:
	DrawTitleBox(48, " D I S P L A Y  C O N T E N T", "File Editor")

	// Opens the file, checks for errors:
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return
	}

	// Displays content and close menu:
	fmt.Printf("\n[%v]\n\n", string(content))
	ConfirmMenu("Back")
}

func AppendToFile(filename string) {
	// Cosmetic purpose only:
	DrawTitleBox(48, "A P P E N D  C O N T E N T", "File Editor")

	// Opens the file, checks for errors:
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return
	}
	defer file.Close()

	// Asks the user for new content to append to file:
	fmt.Println("\nNOTE: Only \\t and \\n are converted into escape chars.\n     `\\n` is automatically converted to `\\r\\n`.\n\nContent to append:")
	answer := ConvertEscapeChars(AskUserString())

	// Checks for errors while writing in the file:
	if _, err = file.Write([]byte(answer)); err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return
	}

	// Prints confirmation and close menu:
	fmt.Print("\nSuccessfully appended content to the file !\n\n")
	ConfirmMenu("Back")
}

func OverwriteFile(filename string) {
	// Cosmetic purpose only:
	DrawTitleBox(48, "O V E R W R I T E  C O N T E N T", "File Editor")

	// Opens the file, checks for errors:
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return
	}
	defer file.Close()

	// Asks the user for new content to write into file:
	fmt.Println("\nNOTE: Only \\t and \\n are converted into escape chars.\n     `\\n` is automatically converted to `\\r\\n`.\n\nOverwrite file content with:")

	// Checks for errors while writing in the file:
	if _, err = file.Write([]byte([]byte(AskUserString()))); err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return
	}

	// Prints confirmation and close menu:
	fmt.Print("\nSuccessfully overwrote file content !\n\n")
	ConfirmMenu("Back")
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	file.Close()
	return err
}

func EmptyFile(filename string) {
	// Cosmetic purpose only:
	DrawTitleBox(48, "E M P T Y  F I L E  C O N T E N T", "File Editor")
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		return
	}
	defer file.Close()

	// Asks for user confirmation as it is a destruction action:
	fmt.Printf("\nAre you sure you wanna remove all content from `%v` ?\n\n1 // Yes\n0 // No\n", filename)
	if answer := AskUserInt(0, 1); answer == 0 {
		fmt.Println()
		return
	}

	// Prints confirmation and close menu:
	fmt.Println("\nSuccessfully emptied file content !\n\n0 // Back")
	ConfirmMenu("Back")
}

func DeleteFile(filename string) bool {
	// Cosmetic purpose only:
	DrawTitleBox(48, "D E L E T E  F I L E", "File Editor")

	// Asks for user confirmation as it is a destruction action:
	fmt.Printf("\nAre you sure you wanna delete %v ?\nThis action is irreversible and cannot be undone.\n\n1 // Yes\n0 // No\n", filename)
	if answer := AskUserInt(0, 1); answer == 0 {
		fmt.Println()
		return false
	}

	// Checks for errors when trying to delete the file:
	if err := os.Remove(filename); err != nil {
		fmt.Printf("\n%v\n", err)
		fmt.Println()
		ConfirmMenu("Back")
		return false
	}

	fmt.Println("\nSuccessfully deleted file ! Back to main menu...\n\n0 // Back")
	ConfirmMenu("Back")
	return true
}
