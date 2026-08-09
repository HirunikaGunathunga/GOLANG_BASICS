package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	NoteTitle   string    `json:"title"`
	NoteContain string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note title %v has the following content: \n\n%v", note.NoteTitle, note.NoteContain)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.NoteTitle, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid entry")
	}

	return Note{
		NoteTitle:   title,
		NoteContain: content,
		CreatedAt:   time.Now(),
	}, nil
}
