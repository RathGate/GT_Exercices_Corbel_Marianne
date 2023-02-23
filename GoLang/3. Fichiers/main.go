package main

import (
	"bufio"
	"fichiers/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	OpenFileMenu(filename)
}

func ReadFile(filename string) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FILE CONTENT:")
	fmt.Println(string(dat))
}

func AppendToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Content to append:")
	file.Write([]byte(AskUserString()))
}

func AskUserString() (answer string) {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscanln(stdin, &answer)
	stdin.Discard(stdin.Buffered())
	return answer
}

func OverwriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Overwrite file content with:")
	file.Write([]byte(AskUserString()))
}

func EmptyFile(filename string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func OpenFileMenu(filename string) {
	utils.DrawTitleBox(48, "F I L E  M E N U", "")

	fmt.Println("Currently open file:\n\t" + filename + "\n")
	ReadFile(filename)
	AppendToFile(filename)
	ReadFile(filename)
	// EmptyFile(filename)
	// ReadFile(filename)
	os.Create("moncul.txt")
	err := os.Remove("moncul.txt")
	if err != nil {
		fmt.Println(err)
	}
}
