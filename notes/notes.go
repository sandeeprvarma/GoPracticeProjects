package main

import (
	"bufio"
	"errors"
	"fmt"
	"notes/userNotes"
	"os"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
	// var note Note

	userNote := userNotes.New(title, content)

	userNote.Display()
	userNote.Save()
	fmt.Println("Saving the note successfully!")
	// os.WriteFile(data)
}

func getNoteData() (string, string, error) {
	title, err1 := getUserInput("Note title: ")
	if err1 != nil {
		return "", "", err1
	}
	content, err2 := getUserInput("Note content: ")
	if err2 != nil {
		return "", "", err2
	}

	return title, content, nil
}

// func getUserInput(prompt string) (string, error) {
// 	fmt.Print(prompt)
// 	var value string
// 	fmt.Scanln(&value)

// 	if value == "" {
// 		return value, errors.New("value can not be blank")
// 	}
// 	return value, nil
// }

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("Invalid inpput")
	}
	return text, nil
}
