package todo

import (
	"encoding/json"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) *Todo {
	return &Todo{
		content,
	}
}

func (t *Todo) Save() error {
	jsonTodo, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile("todo.json", jsonTodo, 0644)
}

func (t *Todo) Display() {
	println(t.Content)
}
