package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NoteFunc() {
	title, content := getNoteData()
	userNote, err := New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Cannot save the file")
	}
	fmt.Println("\nSAVE SUCCESSFUL")
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note contain: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	//var value string
	//fmt.Scanln(&value)//scan cannot get long text
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	//return value
	return text
}
