package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mobamoh/go-playground/easy-go-note/note"
	"github.com/mobamoh/go-playground/easy-go-note/todo"
)

func main() {

	println("------- Note -------")
	title, err := getUserData("Note title: ")
	if err != nil {
		panic(err)
	}
	content, err := getUserData("Note Content: ")
	if err != nil {
		panic(err)
	}

	note := note.New(title, content)
	if err = outputData(note); err != nil {
		panic(err)
	}
	// note.Display()
	// if err = note.Save(); err != nil {
	// 	panic(err)
	// }

	println("------- Todo -------")
	todoTitle, err := getUserData("Todo: ")
	if err != nil {
		panic(err)
	}
	todo := todo.New(todoTitle)
	if err = outputData(todo); err != nil {
		panic(err)
	}
	// todo.Display()
	// if err = todo.Save(); err != nil {
	// 	panic(err)
	// }

}

func getUserData(text string) (string, error) {

	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\r\n")
	return text, nil
}

func outputData(data Ouputter) error {
	data.Display()
	if err := data.Save(); err != nil {
		return err
	}
	return nil
}

type Displayer interface {
	Display()
}

type Saver interface {
	Save() error
}

type Ouputter interface {
	Displayer
	Saver
}
