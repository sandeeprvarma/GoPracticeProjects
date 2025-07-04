package userNotes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (note Note) Display() {
	fmt.Printf("Title: %v \n, body: %v \n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
	// os.WriteFile(fileName, []byte(note.content))
}
