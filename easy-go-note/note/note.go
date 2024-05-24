package note

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

func New(title, content string) *Note {
	return &Note{Title: title, Content: content, CreatedAt: time.Now()}
}

func (n *Note) Display() {
	fmt.Println(n.Title)
	fmt.Println(n.Content)
}

func (n *Note) Save() error {
	fileName := strings.Replace(n.Title, " ", "_", -1)
	fileName = strings.ToLower(fileName) + ".json"
	noteJson, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, noteJson, 0644)
}
